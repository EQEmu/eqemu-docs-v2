package console

import (
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type GMSubcommandsDocsGenerateCommand struct {
	command *cobra.Command
}

func (c *GMSubcommandsDocsGenerateCommand) Command() *cobra.Command {
	return c.command
}

func NewGMSubcommandsDocsGenerateCommand() *GMSubcommandsDocsGenerateCommand {
	i := &GMSubcommandsDocsGenerateCommand{
		command: &cobra.Command{
			Use:   "doc:gm-subcommands",
			Short: "Command used for generated GM Subcommands Documentation",
		},
	}

	i.command.Args = i.Validate
	i.command.Run = i.Handle

	return i
}

func (c *GMSubcommandsDocsGenerateCommand) Validate(_ *cobra.Command, _ []string) error {
	return nil
}

func (c *GMSubcommandsDocsGenerateCommand) Handle(_ *cobra.Command, _ []string) {
	b, err := ioutil.ReadFile("./templates/subcommands-page.template")
	if err != nil {
		fmt.Println(err)
	}

	template := string(b)

	commandData := fmt.Sprintf("%v\n", template)

	commandData += "| Command | Subcommand | Usage |\n"

	commandData += "| :--- | :--- | :--- |\n"

	commandFiles := [3]string{"find", "set", "show"}

	for _, command := range commandFiles {
		resp, err := http.Get(fmt.Sprintf("https://raw.githubusercontent.com/EQEmu/Server/master/zone/gm_commands/%v.cpp", command))
		if err != nil {
			fmt.Println(err)
		}
	
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
		}
	
		for _, line := range strings.Split(string(body), "\n") {
			if strings.Contains(line, "Cmd{.cmd = \"") {
				line = strings.ReplaceAll(line, "Cmd{.cmd = \"", "")
				line = strings.ReplaceAll(line, "\t", "")
				line = strings.ReplaceAll(line, ".u = ", "")
	
				lineData := strings.Split(line, "\", ")
	
				subcommand := strings.TrimSpace(lineData[0])
				subcommand = strings.ReplaceAll(subcommand, "\"", "")
	
				help := strings.TrimSpace(lineData[1])
				help = strings.ReplaceAll(help, "\"", "")
				help = strings.ReplaceAll(help, "|", "&#124;")
	
				commandData += fmt.Sprintf(
					"| #%v | %v | #%v %v |\n",
					command,
					subcommand,
					command,
					help,
				)
			}
		}
	}

	fmt.Println(commandData)

	err = os.WriteFile("./docs/server/operation/in-game-subcommand-reference.md", []byte(commandData), os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}
}
