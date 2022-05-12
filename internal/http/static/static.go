package static

import (
	"fmt"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/labstack/echo/v4"
)

type (
	// StaticConfig defines the config for Static middleware.
	StaticConfig struct {
		// Skipper defines a function to skip middleware.
		Skipper middleware.Skipper

		// Root directory from where the static content is served.
		// Required.
		Root string `yaml:"root"`

		// Index file for serving a directory.
		// Optional. Default value "index.html".
		Index string `yaml:"index"`

		// Enable HTML5 mode by forwarding all not-found requests to root so that
		// SPA (single-page application) can handle the routing.
		// Optional. Default value false.
		HTML5 bool `yaml:"html5"`

		// Enable directory browsing.
		// Optional. Default value false.
		Browse bool `yaml:"browse"`

		// Enable ignoring of the base of the URL path.
		// Example: when assigning a static middleware to a non root path group,
		// the filesystem path is not doubled
		// Optional. Default value false.
		IgnoreBase bool `yaml:"ignoreBase"`

		// Filesystem provides access to the static content.
		// Optional. Defaults to http.Dir(config.Root)
		Filesystem http.FileSystem `yaml:"-"`
	}
)

var (
	// DefaultStaticConfig is the default Static middleware config.
	DefaultStaticConfig = StaticConfig{
		Skipper: middleware.DefaultSkipper,
		Index:   "index.html",
	}
)

// Static returns a Static middleware to serves static content from the provided
// root directory.
func Static(root string) echo.MiddlewareFunc {
	c := DefaultStaticConfig
	c.Root = root
	return StaticWithConfig(c)
}

// StaticWithConfig returns a Static middleware with config.
// See `Static()`.
func StaticWithConfig(config StaticConfig) echo.MiddlewareFunc {
	// Defaults
	if config.Root == "" {
		config.Root = "." // For security we want to restrict to CWD.
	}
	if config.Skipper == nil {
		config.Skipper = DefaultStaticConfig.Skipper
	}
	if config.Index == "" {
		config.Index = DefaultStaticConfig.Index
	}
	if config.Filesystem == nil {
		config.Filesystem = http.Dir(config.Root)
		config.Root = "."
	}

	// Index template
	//t, err := template.New("index").Parse(html)
	//if err != nil {
	//	panic(fmt.Sprintf("echo: %v", err))
	//}

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			if config.Skipper(c) {
				return next(c)
			}

			p := c.Request().URL.Path
			if strings.HasSuffix(c.Path(), "*") { // When serving from a group, e.g. `/static*`.
				p = c.Param("*")
			}
			p, err = url.PathUnescape(p)
			if err != nil {
				return
			}
			name := filepath.Join(config.Root, filepath.Clean("/"+p)) // "/"+ for security

			if config.IgnoreBase {
				routePath := path.Base(strings.TrimRight(c.Path(), "/*"))
				baseURLPath := path.Base(p)
				if baseURLPath == routePath {
					i := strings.LastIndex(name, routePath)
					name = name[:i] + strings.Replace(name[i:], routePath, "", 1)
				}
			}

			file, err := openFile(config.Filesystem, name)
			if err != nil {
				if !os.IsNotExist(err) {
					return err
				}

				if err = next(c); err == nil {
					return err
				}

				he, ok := err.(*echo.HTTPError)
				if !(ok && config.HTML5 && he.Code == http.StatusNotFound) {
					file, err = openFile(config.Filesystem, filepath.Join(config.Root, "404.html"))
					if err != nil {
						fmt.Println(err)
					}
					defer file.Close()
					info, err := file.Stat()
					if err != nil {
						return err
					}

					c.Response().WriteHeader(http.StatusGone)

					return serveFile(c, file, info)
				}

				file, err = openFile(config.Filesystem, filepath.Join(config.Root, config.Index))
				if err != nil {
					return err
				}
			}

			defer file.Close()

			info, err := file.Stat()
			if err != nil {
				return err
			}

			if info.IsDir() {
				index, err := openFile(config.Filesystem, filepath.Join(name, config.Index))
				if err != nil {

					if os.IsNotExist(err) {
						return next(c)
					}
				}

				defer index.Close()

				info, err = index.Stat()
				if err != nil {
					return err
				}

				return serveFile(c, index, info)
			}

			return serveFile(c, file, info)
		}
	}
}

func openFile(fs http.FileSystem, name string) (http.File, error) {
	pathWithSlashes := filepath.ToSlash(name)
	return fs.Open(pathWithSlashes)
}

func serveFile(c echo.Context, file http.File, info os.FileInfo) error {
	http.ServeContent(c.Response(), c.Request(), info.Name(), info.ModTime(), file)
	return nil
}
