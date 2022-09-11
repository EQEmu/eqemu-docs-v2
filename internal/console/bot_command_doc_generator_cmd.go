package console

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

type BotCommandsDocsGenerateCommand struct {
	command *cobra.Command
}

func (c *BotCommandsDocsGenerateCommand) Command() *cobra.Command {
	return c.command
}

func NewBotCommandsDocsGenerateCommand() *BotCommandsDocsGenerateCommand {
	i := &BotCommandsDocsGenerateCommand{
		command: &cobra.Command{
			Use:   "doc:bot-commands",
			Short: "Command used for generated Bot Commands Documentation",
		},
	}

	i.command.Args = i.Validate
	i.command.Run = i.Handle

	return i
}

func (c *BotCommandsDocsGenerateCommand) Validate(_ *cobra.Command, _ []string) error {
	return nil
}

func (c *BotCommandsDocsGenerateCommand) Handle(_ *cobra.Command, _ []string) {
	resp, err := http.Get("https://raw.githubusercontent.com/EQEmu/Server/master/zone/bot_command.cpp")
	if err != nil {
		fmt.Println(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	b, err := ioutil.ReadFile("./templates/bot-commands-page.template")
	if err != nil {
		fmt.Println(err)
	}

	template := string(b)

	currentTime := time.Now()
	generatedTime := currentTime.Format("2006.01.02")
	template = strings.ReplaceAll(template, "{{generated_time}}", generatedTime)

	commandData := fmt.Sprintf("%v\n", template)

	commandData += "| Command | Description | Status Level |\n"

	commandData += "| :--- | :--- | :--- |\n"

	for _, line := range strings.Split(string(body), "\n") {
		if strings.Contains(line, "bot_command_add(\"") {
			line = strings.ReplaceAll(line, "bot_command_add(", "")
			line = strings.ReplaceAll(line, "\t", "")

			lineData := strings.Split(line, "\", ")

			command := strings.TrimSpace(lineData[0])
			command = strings.ReplaceAll(command, "\"", "")

			help := strings.TrimSpace(lineData[1])
			help = strings.ReplaceAll(help, "\"", "")
			help = strings.ReplaceAll(help, "|", "&#124;")

			accountStatus := getStringInBetween(line, "AccountStatus::", ",")

			statusValue := GetStatusValue(accountStatus)

			commandData += fmt.Sprintf(
				"| ^%v | %v | %v (%v) |\n",
				command,
				help,
				accountStatus,
				statusValue,
			)
		}
	}

	fmt.Println(commandData)

	err = os.WriteFile("./docs/server/bots/bot-commands.md", []byte(commandData), os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}
}
