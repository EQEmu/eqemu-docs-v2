package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	copy "github.com/otiai10/copy"
	"log"
	"net/http"
	"os"
	"os/exec"
	"time"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	SendDeployMessage()

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

type DiscordRequestBody struct {
	Content string `json:"content"`
}

func SendDeployMessage() {
	siteUrl := os.Getenv("SITE_URL")
	err := SendDiscordWebhook(fmt.Sprintf("[Docs Deploy] Changes are now live on [%v]", siteUrl))
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
