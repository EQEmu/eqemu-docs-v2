package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	copy "github.com/otiai10/copy"
	"log"
	"net/http"
	"os/exec"
	"time"
)

func main() {
	e := echo.New()
	e.POST(
		"/deploy", func(c echo.Context) error {
			if c.Request().Header.Get("X-Github-Event") != "" &&
				c.Request().Header.Get("X-Github-Delivery") != "" {
				start := time.Now()

				// run git pull
				cmd := exec.Command("git", "pull")
				_, err := cmd.Output()
				if err != nil {
					log.Println(err)
				}
				
				// run mkdocs build
				cmd = exec.Command("mkdocs", "build")
				_, err = cmd.Output()

				if err != nil {
					log.Println(err)
				}

				// copy from site to public
				err = copy.Copy("site", "public")
				if err != nil {
					log.Println(err)
				}

				return c.JSON(
					http.StatusOK,
					echo.Map{
						"data": "Deployed", "time": fmt.Sprintf("%s", time.Since(start)),
					},
				)
			}

			return c.JSON(http.StatusOK, echo.Map{"data": "Invalid request"})
		},
	)
	e.Static("/", "public")
	e.Logger.Fatal(e.Start(":8081"))
}
