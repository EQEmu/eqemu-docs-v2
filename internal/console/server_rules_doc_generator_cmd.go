package console

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

type ServerRulesDocsGenerateCommand struct {
	command *cobra.Command
}

func (c *ServerRulesDocsGenerateCommand) Command() *cobra.Command {
	return c.command
}

func NewServerRulesDocsGenerateCommand() *ServerRulesDocsGenerateCommand {
	i := &ServerRulesDocsGenerateCommand{
		command: &cobra.Command{
			Use:   "doc:server-rules",
			Short: "Command used for generated Server Rules Documentation",
		},
	}

	i.command.Args = i.Validate
	i.command.Run = i.Handle

	return i
}

func (c *ServerRulesDocsGenerateCommand) Validate(_ *cobra.Command, _ []string) error {
	return nil
}

func (c *ServerRulesDocsGenerateCommand) Handle(_ *cobra.Command, _ []string) {
	resp, err := http.Get("https://raw.githubusercontent.com/EQEmu/Server/master/common/ruletypes.h")
	if err != nil {
		fmt.Println(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	b, err := ioutil.ReadFile("./templates/server-rules-page.template")
	if err != nil {
		fmt.Println(err)
	}

	template := string(b)

	currentTime := time.Now()
	generatedTime := currentTime.Format("2006.01.02")
	template = strings.ReplaceAll(template, "{{generated_time}}", generatedTime)

	serverRulesCategoryData := make(map[string]string)

	serverRulesData := fmt.Sprintf("%v\n", template)

	for _, line := range strings.Split(string(body), "\n") {
		if strings.HasPrefix(line, "RULE_BOOL(") || strings.HasPrefix(line, "RULE_INT(") || strings.HasPrefix(line, "RULE_REAL(") {
			currentLine := strings.ReplaceAll(line, "RULE_BOOL", "")
			currentLine = strings.ReplaceAll(currentLine, "RULE_INT", "")
			currentLine = strings.ReplaceAll(currentLine, "RULE_REAL", "")

			lineData := strings.Split(currentLine, ", \"")

			currentServerRuleData := strings.Split(lineData[0], ",")
			ruleType := strings.ReplaceAll(currentServerRuleData[0], "(", "")
			ruleName := strings.ReplaceAll(currentServerRuleData[1], " ", "")
			ruleName = fmt.Sprintf("%v:%v", ruleType, ruleName)
			ruleDefaultValue := strings.ReplaceAll(currentServerRuleData[2], " ", "")
			ruleDesc := strings.ReplaceAll(lineData[1], "\")", "")
			ruleDesc = strings.ReplaceAll(ruleDesc, "|", "&#124;")

			serverRulesCategoryData[ruleType] += fmt.Sprintf("| %v | %v | %v |\n", ruleName, ruleDefaultValue, ruleDesc)
		}
	}

	keys := make([]string, 0)

	for k, _ := range serverRulesCategoryData {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	for _, serverRuleType := range keys {
		serverRulesData += fmt.Sprintf("## %v Rules\n", serverRuleType)

		serverRulesData += "| Rule Name | Default Value | Description |\n"

		serverRulesData += "| :--- | :--- | :--- |\n"

		lineEnding := "\n"

		if serverRuleType == keys[len(keys)-1] {
			lineEnding = ""
		}

		serverRulesData += fmt.Sprintf("%v%v", serverRulesCategoryData[serverRuleType], lineEnding)
	}

	err = os.WriteFile("./docs/server/operation/server-rules.md", []byte(serverRulesData), os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}
}
