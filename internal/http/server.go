package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/EQEmu/eqemu-docs-v2/internal/http/static"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	copyo "github.com/otiai10/copy"
	"log"
	"net/http"
	"os"
	"os/exec"
	"time"
)

func Run() {
	e := echo.New()
	e.Use(
		middleware.GzipWithConfig(
			middleware.GzipConfig{
				Level: 5,
			},
		),
	)
	e.Use(middleware.Logger())
	e.POST(
		"/deploy", func(c echo.Context) error {
			if c.Request().Header.Get("X-Github-Event") != "" &&
				c.Request().Header.Get("X-Github-Delivery") != "" {
				err := SendDiscordWebhook(
					fmt.Sprintf("[Docs Deploy] Starting build..."),
				)
				if err != nil {
					log.Println(err)
				}

				start := time.Now()

				// run git pull
				cmd := exec.Command("git", "pull")
				_, err = cmd.Output()
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
				err = copyo.Copy("site", "public")
				if err != nil {
					log.Println(err)
				}

				SendDeployMessage(start)

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

	e.Use(
		static.StaticWithConfig(
			static.StaticConfig{
				Root:   "public",
				Index:  "index.html",
				Browse: false,
				HTML5:  false,
			},
		),
	)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%v", os.Getenv("HTTP_PORT"))))
}

type DiscordRequestBody struct {
	Content string `json:"content"`
}

func SendDeployMessage(start time.Time) {
	siteUrl := os.Getenv("SITE_URL")
	err := SendDiscordWebhook(
		fmt.Sprintf("[Docs Deploy] Changes are now live on [%v] **took %s second(s)**",
			siteUrl,
			time.Since(start).Round(time.Second),
		),
	)
	if err != nil {
		log.Println(err)
	}
}

func SendDiscordWebhook(msg string) error {
	webhookUrl := os.Getenv("DISCORD_WEBHOOK_URL")

	body, _ := json.Marshal(DiscordRequestBody{Content: msg})
	req, err := http.NewRequest(http.MethodPost, webhookUrl, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(resp.Body)
	if err != nil {
		return err
	}
	return nil
}
