package console

import (
	"encoding/json"
	"fmt"
	"github.com/EQEmu/eqemu-docs-v2/config"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type QuestApiDocGeneratorCommand struct {
	command *cobra.Command
}

func (c *QuestApiDocGeneratorCommand) Command() *cobra.Command {
	return c.command
}

func NewQuestApiDocGeneratorCommand() *QuestApiDocGeneratorCommand {
	i := &QuestApiDocGeneratorCommand{
		command: &cobra.Command{
			Use:   "doc:quest-api-doc-generate",
			Short: "Generates Quest API Docs from Spire API",
		},
	}

	i.command.Args = i.Validate
	i.command.Run = i.Handle

	return i
}

type QuestApiResponse struct {
	Data struct {
		LastRefreshed time.Time `json:"last_refreshed"`
		PerlApi       `json:"perl"`
		LuaApi        `json:"lua"`
	} `json:"data"`
}

type PerlApi struct {
	PerlMethods   map[string][]Method   `json:"methods"`
	PerlEvents    []Event               `json:"events"`
	PerlConstants map[string][]Constant `json:"constants"`
}

type LuaApi struct {
	LuaMethods   map[string][]Method   `json:"methods"`
	LuaEvents    []Event               `json:"events"`
	LuaConstants map[string][]Constant `json:"constants"`
}

type Event struct {
	EntityType      string   `json:"entity_type"`      // @eg Player
	EventName       string   `json:"event_name"`       // @eg EVENT_SAY
	EventIdentifier string   `json:"event_identifier"` // @eg EVENT_SAY
	EventVars       []string `json:"event_vars"`       // @eg []{"text", "langid"}
}
type Constant struct {
	Constant string `json:"constant"` // @eg var
}

type Method struct {
	Method     string   `json:"method"`
	Params     []string `json:"params"`
	MethodType string   `json:"method_type"`
	ReturnType string   `json:"return_type"`
	Categories []string `json:"categories"`
}

func (c *QuestApiDocGeneratorCommand) Handle(_ *cobra.Command, _ []string) {
	response := c.GetSpireDefinitions()

	section := config.GetMkDocsQuestApiConfig()

	c.WriteMethodDocs(response, section)
}

func (c *QuestApiDocGeneratorCommand) Validate(_ *cobra.Command, _ []string) error {
	return nil
}

func (c *QuestApiDocGeneratorCommand) GetSpireDefinitions() QuestApiResponse {
	resp, err := http.Get("http://spire.akkadius.com/api/v1/quest-api/definitions")
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var response QuestApiResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println(err)
		return QuestApiResponse{}
	}

	return response
}

func (c *QuestApiDocGeneratorCommand) BuildMethodPage(methodType string, response QuestApiResponse) string {

	b, err := ioutil.ReadFile("./templates/method-page.template")
	if err != nil {
		log.Fatal(err)
	}

	template := string(b)

	luaTemplate := template
	perlTemplate := template

	perlMethods := []string{}
	for _, methods := range response.Data.PerlApi.PerlMethods {
		for _, method := range methods {
			if method.MethodType == methodType {
				delim := "->"
				prefix := "$"
				if method.MethodType == "quest" {
					delim = "::"
					prefix = ""
				}

				perlMethods = append(perlMethods,
					fmt.Sprintf(
						"    %v%v%v%v(%v);",
						prefix,
						strings.ToLower(method.MethodType),
						delim,
						method.Method,
						strings.Join(method.Params, ", "),
					),
				)
			}
		}
	}

	luaMethods := []string{}
	for _, methods := range response.Data.LuaApi.LuaMethods {
		for _, method := range methods {
			if method.MethodType == methodType {
				delim := ":"
				if method.MethodType == "eq" {
					delim = "."
				}

				luaMethods = append(luaMethods,
					fmt.Sprintf(
						"    %v%v%v(%v);",
						strings.ToLower(method.MethodType),
						delim,
						method.Method,
						strings.Join(method.Params, ", "),
					),
				)
			}
		}
	}

	markdown := ``

	currentTime := time.Now()
	generatedTime := currentTime.Format("2006.01.02 15:04:05")

	if len(perlMethods) > 0 {
		perlTemplate = strings.ReplaceAll(perlTemplate, "{{generated_time}}", generatedTime)
		perlTemplate = strings.ReplaceAll(perlTemplate, "{{method_type}}", methodType)
		perlTemplate = strings.ReplaceAll(perlTemplate, "{{language_label}}", "Perl")
		perlTemplate = strings.ReplaceAll(perlTemplate, "{{language_syntax}}", "perl")
		perlTemplate = strings.ReplaceAll(perlTemplate, "{{language_methods}}", strings.Join(perlMethods, "\n"))

		count := strconv.Itoa(len(perlMethods))
		perlTemplate = strings.ReplaceAll(perlTemplate, "{{method_count}}", count)
		markdown += perlTemplate
	}
	if len(luaMethods) > 0 {
		luaTemplate = strings.ReplaceAll(luaTemplate, "{{generated_time}}", generatedTime)
		luaTemplate = strings.ReplaceAll(luaTemplate, "{{method_type}}", methodType)
		luaTemplate = strings.ReplaceAll(luaTemplate, "{{language_label}}", "Lua")
		luaTemplate = strings.ReplaceAll(luaTemplate, "{{language_syntax}}", "lua")
		luaTemplate = strings.ReplaceAll(luaTemplate, "{{language_methods}}", strings.Join(luaMethods, "\n"))

		count := strconv.Itoa(len(luaMethods))
		luaTemplate = strings.ReplaceAll(luaTemplate, "{{method_count}}", count)
		markdown += luaTemplate
	}

	return markdown
}

func (c *QuestApiDocGeneratorCommand) WriteMethodDocs(response QuestApiResponse, section []map[string][]map[string]string) {
	// zero out first
	section[0]["Methods"] = []map[string]string{}

	methodTypes := map[string]bool{}
	for _, methods := range response.Data.PerlApi.PerlMethods {
		for _, method := range methods {
			methodTypes[method.MethodType] = true
		}
	}

	for _, methods := range response.Data.LuaApi.LuaMethods {
		for _, method := range methods {
			methodTypes[method.MethodType] = true
		}
	}

	// sort alpha
	keys := make([]string, 0, len(methodTypes))
	for k := range methodTypes {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	ignoreMethods := []string{
		"EQDBRes",
		"EQDB",
		"EntityListDeprecated",
	}

	// loop through sorted
	for _, methodType := range keys {
		if contains(ignoreMethods, methodType) {
			continue
		}

		markdown := c.BuildMethodPage(methodType, response)
		pagePath := fmt.Sprintf("./docs/quest-api/methods/%v.md", strings.ToLower(methodType))

		// write
		err := os.WriteFile(pagePath, []byte(markdown), os.ModePerm)
		if err != nil {
			log.Println(err)
		}

		// pop docs off for mkdocs link
		pagePath = strings.ReplaceAll(pagePath, "./docs/", "")
		m := map[string]string{methodType: pagePath}
		section[0]["Methods"] = append(section[0]["Methods"], m)
	}

	mkdocsCfg, err := config.GetMkdocsConfig()
	if err != nil {
		log.Fatal(err)
	}

	// write nav block
	for i, nav := range mkdocsCfg.Nav {
		if len(nav.QuestApi) > 0 {
			mkdocsCfg.Nav[i].QuestApi[0]["Methods"] = section[0]["Methods"]
		}
	}

	config.WriteMkdocsConfig(mkdocsCfg)
}

func (c *QuestApiDocGeneratorCommand) FormatMethodType(methodType string) string {
	methodType = strings.ReplaceAll(methodType, "Doors", "Door")
	return methodType
}
