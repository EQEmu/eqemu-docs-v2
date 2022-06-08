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

type GMCommandsDocsGenerateCommand struct {
	command *cobra.Command
}

func (c *GMCommandsDocsGenerateCommand) Command() *cobra.Command {
	return c.command
}

func NewGMCommandsDocsGenerateCommand() *GMCommandsDocsGenerateCommand {
	i := &GMCommandsDocsGenerateCommand{
		command: &cobra.Command{
			Use:   "doc:gm-commands",
			Short: "Command used for generated GM Commands Documentation",
		},
	}

	i.command.Args = i.Validate
	i.command.Run = i.Handle

	return i
}

func (c *GMCommandsDocsGenerateCommand) Validate(_ *cobra.Command, _ []string) error {
	return nil
}

func (c *GMCommandsDocsGenerateCommand) Handle(_ *cobra.Command, _ []string) {
	resp, err := http.Get("https://raw.githubusercontent.com/EQEmu/Server/master/zone/command.cpp")
	if err != nil {
		fmt.Println(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	b, err := ioutil.ReadFile("./templates/commands-page.template")
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
		if strings.Contains(line, "command_add(\"") {
			line = strings.ReplaceAll(line, "command_add(", "")
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
				"| #%v | %v | %v (%v) |\n",
				command,
				help,
				accountStatus,
				statusValue,
			)
		}
	}

	fmt.Println(commandData)

	err = os.WriteFile("./docs/server/operation/in-game-command-reference.md", []byte(commandData), os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}
}

func GetStatusValue(statusLevel string) int {
	m := map[string]int{
		"Player":          0,
		"Steward":         10,
		"ApprenticeGuide": 20,
		"Guide":           50,
		"QuestTroupe":     80,
		"SeniorGuide":     81,
		"GMTester":        85,
		"EQSupport":       90,
		"GMStaff":         95,
		"GMAdmin":         100,
		"GMLeadAdmin":     150,
		"QuestMaster":     160,
		"GMAreas":         170,
		"GMCoder":         180,
		"GMMgmt":          200,
		"GMImpossible":    250,
		"Max":             255,
	}

	if _, ok := m[statusLevel]; !ok {
		return 0
	}

	return m[statusLevel]
}
