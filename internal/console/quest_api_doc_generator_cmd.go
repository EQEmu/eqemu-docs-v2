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

const (
	NavPositionMethods   = 1
	NavPositionEvents    = 2
	NavPositionConstants = 3
)

func (c *QuestApiDocGeneratorCommand) Handle(_ *cobra.Command, _ []string) {
	response := c.GetSpireDefinitions()
	section := config.GetMkDocsQuestApiConfig()

	c.WriteMethodDocs(response, section)
	c.WriteEventDocs(response, section)
	c.WriteConstantsDocs(response, section)
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
	b, err := os.ReadFile("./templates/method-page.template")
	if err != nil {
		log.Fatal(err)
	}

	template := string(b)

	luaTemplate := template
	perlTemplate := template

	var perlMethods []string
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

	if len(perlMethods) > 0 {
		perlTemplate = strings.ReplaceAll(perlTemplate, "{{method_type}}", methodType)
		perlTemplate = strings.ReplaceAll(perlTemplate, "{{language_label}}", "Perl")
		perlTemplate = strings.ReplaceAll(perlTemplate, "{{language_syntax}}", "perl")
		perlTemplate = strings.ReplaceAll(perlTemplate, "{{language_methods}}", strings.Join(perlMethods, "\n"))

		count := strconv.Itoa(len(perlMethods))
		perlTemplate = strings.ReplaceAll(perlTemplate, "{{method_count}}", count)
		markdown += perlTemplate
	}
	if len(luaMethods) > 0 {
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
	fmt.Println("> Writing methods docs")

	// zero out first
	section[NavPositionMethods]["Methods"] = []map[string]string{}

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

	docsWritten := 0

	// loop through sorted
	for _, methodType := range keys {
		if containsStringSlice(ignoreMethods, methodType) {
			continue
		}

		markdown := c.BuildMethodPage(methodType, response)
		pagePath := fmt.Sprintf("./docs/quest-api/methods/%v.md", strings.ToLower(methodType))

		// write
		err := os.WriteFile(pagePath, []byte(markdown), os.ModePerm)
		if err != nil {
			log.Println(err)
		}

		docsWritten++
		fmt.Printf(" Wrote [%v]\n", pagePath)

		// pop docs off for mkdocs link
		pagePath = strings.ReplaceAll(pagePath, "./docs/", "")
		m := map[string]string{methodType: pagePath}
		section[NavPositionMethods]["Methods"] = append(section[NavPositionMethods]["Methods"], m)
	}

	mkdocsCfg, err := config.GetMkdocsConfig()
	if err != nil {
		log.Fatal(err)
	}

	// write nav block
	for i, nav := range mkdocsCfg.Nav {
		if len(nav.QuestApi) > 0 {
			mkdocsCfg.Nav[i].QuestApi[NavPositionMethods]["Methods"] = section[NavPositionMethods]["Methods"]
		}
	}

	config.WriteMkdocsConfig(mkdocsCfg)

	fmt.Printf("> Wrote method docs [%v]\n", docsWritten)
}

func (c *QuestApiDocGeneratorCommand) FormatMethodType(methodType string) string {
	methodType = strings.ReplaceAll(methodType, "Doors", "Door")
	return methodType
}

func (c *QuestApiDocGeneratorCommand) WriteEventDocs(response QuestApiResponse, section []map[string][]map[string]string) {
	fmt.Println("> Writing event docs")

	docsWritten := 0

	// zero out first
	section[NavPositionEvents]["Events"] = []map[string]string{}

	// perl
	eventTypes := map[string]bool{}
	for _, event := range response.Data.PerlApi.PerlEvents {
		eventTypes[event.EntityType] = true
	}

	// sort alpha
	keys := make([]string, 0, len(eventTypes))
	for k := range eventTypes {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// perl
	for _, eventType := range keys {
		pageLabel := fmt.Sprintf("Perl [%v]", eventType)
		markdown := c.BuildEventPagePerl(eventType, response.Data.PerlApi.PerlEvents)
		pagePath := fmt.Sprintf("./docs/quest-api/events/perl-%v.md", strings.ToLower(eventType))

		// write
		err := os.WriteFile(pagePath, []byte(markdown), os.ModePerm)
		if err != nil {
			log.Println(err)
		}

		docsWritten++
		fmt.Printf(" Wrote [%v]\n", pagePath)

		// pop docs off for mkdocs link
		pagePath = strings.ReplaceAll(pagePath, "./docs/", "")
		m := map[string]string{pageLabel: pagePath}
		section[NavPositionEvents]["Events"] = append(section[NavPositionEvents]["Events"], m)
	}

	// lua
	eventTypes = map[string]bool{}
	for _, event := range response.Data.LuaApi.LuaEvents {
		eventTypes[event.EntityType] = true
	}

	// sort alpha
	keys = make([]string, 0, len(eventTypes))
	for k := range eventTypes {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// lua
	for _, eventType := range keys {
		pageLabel := fmt.Sprintf("Lua [%v]", eventType)
		markdown := c.BuildEventPageLua(eventType, response.Data.LuaApi.LuaEvents)
		pagePath := fmt.Sprintf("./docs/quest-api/events/lua-%v.md", strings.ToLower(eventType))

		// write
		err := os.WriteFile(pagePath, []byte(markdown), os.ModePerm)
		if err != nil {
			log.Println(err)
		}

		docsWritten++
		fmt.Printf(" Wrote [%v]\n", pagePath)

		// pop docs off for mkdocs link
		pagePath = strings.ReplaceAll(pagePath, "./docs/", "")
		m := map[string]string{pageLabel: pagePath}
		section[NavPositionEvents]["Events"] = append(section[NavPositionEvents]["Events"], m)
	}

	mkdocsCfg, err := config.GetMkdocsConfig()
	if err != nil {
		log.Fatal(err)
	}

	// write nav block
	for i, nav := range mkdocsCfg.Nav {
		if len(nav.QuestApi) > 0 {
			mkdocsCfg.Nav[i].QuestApi[NavPositionEvents]["Events"] = section[NavPositionEvents]["Events"]
		}
	}

	config.WriteMkdocsConfig(mkdocsCfg)

	fmt.Printf("> Wrote event docs [%v]\n", docsWritten)
}

func (c *QuestApiDocGeneratorCommand) BuildEventPagePerl(eventType string, events []Event) string {
	markdown := `!!! info end

    Also see [Spire Quest API Explorer](http://spire.akkadius.com/quest-api-explorer?lang=perl){:target="perl_event"} for latest definitions and Quest examples

`

	b, err := os.ReadFile("./templates/event-block-perl.template")
	if err != nil {
		log.Fatal(err)
	}

	for _, event := range events {
		if eventType == event.EntityType {

			eventBlock := string(b)
			eventMarkdown := `
sub {{event_name}} {
{{exported_vars}}}`

			eventMarkdown = strings.ReplaceAll(eventMarkdown, "{{event_name}}", event.EventName)

			exportedVars := ""
			for _, eventVar := range event.EventVars {
				exportedVars += fmt.Sprintf("\tquest::debug(\"%v \" . $%v);\n", eventVar, eventVar)
			}

			// pop vars onto event string
			eventMarkdown = strings.ReplaceAll(eventMarkdown, "{{exported_vars}}", exportedVars)

			// write event block into code block
			eventBlock = strings.ReplaceAll(eventBlock, "{{event_name}}", event.EventName)
			eventBlock = strings.ReplaceAll(eventBlock, "{{event}}", eventMarkdown)
			markdown += eventBlock
		}
	}

	return markdown
}

func (c *QuestApiDocGeneratorCommand) BuildEventPageLua(eventType string, events []Event) string {
	markdown := `!!! info end

    Also see [Spire Quest API Explorer](http://spire.akkadius.com/quest-api-explorer?lang=lua){:target="perl_event"} for latest definitions and Quest examples

`

	b, err := os.ReadFile("./templates/event-block-lua.template")
	if err != nil {
		log.Fatal(err)
	}

	for _, event := range events {
		if eventType == event.EntityType {
			eventBlock := string(b)
			eventMarkdown := `
function {{event_name}}(e) {
{{exported_vars}}}`

			eventMarkdown = strings.ReplaceAll(eventMarkdown, "{{event_name}}", event.EventName)

			exportedVars := ""
			for _, eventVar := range event.EventVars {
				exportedVars += fmt.Sprintf("\teq.debug(\"%v \" .. e.%v);\n", eventVar, eventVar)
			}

			// pop vars onto event string
			eventMarkdown = strings.ReplaceAll(eventMarkdown, "{{exported_vars}}", exportedVars)

			// write event block into code block
			eventBlock = strings.ReplaceAll(eventBlock, "{{event_name}}", event.EventName)
			eventBlock = strings.ReplaceAll(eventBlock, "{{event}}", eventMarkdown)
			markdown += eventBlock
		}
	}

	return markdown
}

func (c *QuestApiDocGeneratorCommand) WriteConstantsDocs(response QuestApiResponse, section []map[string][]map[string]string) {
	fmt.Println("> Writing constants docs")

	docsWritten := 0

	// zero out first
	section[NavPositionConstants]["Constants"] = []map[string]string{}

	// perl
	constantTypes := map[string]bool{}
	for category, _ := range response.Data.PerlApi.PerlConstants {
		constantTypes[category] = true
	}

	// sort alpha
	keys := make([]string, 0, len(constantTypes))
	for k := range constantTypes {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// perl
	for _, constantType := range keys {
		pageLabel := fmt.Sprintf("Perl [%v]", constantType)
		markdown := c.BuildConstantPage("perl", constantType, response.Data.PerlApi.PerlConstants)
		pagePath := fmt.Sprintf("./docs/quest-api/constants/perl-%v.md", strings.ToLower(constantType))

		// write
		err := os.WriteFile(pagePath, []byte(markdown), os.ModePerm)
		if err != nil {
			log.Println(err)
		}

		docsWritten++

		fmt.Printf(" Wrote [%v]\n", pagePath)

		// pop docs off for mkdocs link
		pagePath = strings.ReplaceAll(pagePath, "./docs/", "")
		m := map[string]string{pageLabel: pagePath}
		section[NavPositionConstants]["Constants"] = append(section[NavPositionConstants]["Constants"], m)
	}

	// lua
	constantTypes = map[string]bool{}
	for category, _ := range response.Data.LuaApi.LuaConstants {
		constantTypes[category] = true
	}

	// sort alpha
	keys = make([]string, 0, len(constantTypes))
	for k := range constantTypes {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// lua
	for _, constantType := range keys {
		pageLabel := fmt.Sprintf("Lua [%v]", constantType)
		markdown := c.BuildConstantPage("lua", constantType, response.Data.LuaApi.LuaConstants)
		pagePath := fmt.Sprintf("./docs/quest-api/constants/lua-%v.md", strings.ToLower(constantType))

		// write
		err := os.WriteFile(pagePath, []byte(markdown), os.ModePerm)
		if err != nil {
			log.Println(err)
		}

		fmt.Printf(" Wrote [%v]\n", pagePath)

		docsWritten++

		// pop docs off for mkdocs link
		pagePath = strings.ReplaceAll(pagePath, "./docs/", "")
		m := map[string]string{pageLabel: pagePath}
		section[NavPositionConstants]["Constants"] = append(section[NavPositionConstants]["Constants"], m)
	}

	mkdocsCfg, err := config.GetMkdocsConfig()
	if err != nil {
		log.Fatal(err)
	}

	// write nav block
	for i, nav := range mkdocsCfg.Nav {
		if len(nav.QuestApi) > 0 {
			mkdocsCfg.Nav[i].QuestApi[NavPositionConstants]["Constants"] = section[NavPositionConstants]["Constants"]
		}
	}

	config.WriteMkdocsConfig(mkdocsCfg)

	fmt.Printf("> Wrote constants docs [%v]\n", docsWritten)
}

func (c *QuestApiDocGeneratorCommand) BuildConstantPage(language string, constantType string, constants map[string][]Constant) string {
	b, err := os.ReadFile("./templates/constants-page.template")
	if err != nil {
		log.Fatal(err)
	}

	template := string(b)
	constantsString := ""
	for category, constants := range constants {
		if category == constantType {
			for _, constant := range constants {
				if language == "perl" {
					constantsString += fmt.Sprintf("%v\n", constant.Constant)
				}
				if language == "lua" {
					constantsString += fmt.Sprintf("%v.%v\n", constantType, constant.Constant)
				}
			}
		}
	}

	template = strings.ReplaceAll(template, "{{constant_type}}", constantType)
	template = strings.ReplaceAll(template, "{{language_syntax}}", language)
	template = strings.ReplaceAll(template, "{{constants}}", constantsString)

	return template
}
