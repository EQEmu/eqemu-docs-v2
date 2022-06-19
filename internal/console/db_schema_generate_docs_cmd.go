package console

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/EQEmu/eqemu-docs-v2/config"
	"github.com/EQEmu/eqemu-docs-v2/internal/database"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

const erdDiagramChunkSize = 8

type DbGenerateDocsCommand struct {
	command *cobra.Command
	tables  []database.TableRelationships
	schema  map[string][]database.DbSchemaRowResult
}

func (c *DbGenerateDocsCommand) Command() *cobra.Command {
	return c.command
}

func NewDbGenerateDocsCommand() *DbGenerateDocsCommand {
	i := &DbGenerateDocsCommand{
		command: &cobra.Command{
			Use:   "doc:db-schema-generate",
			Short: "Generates database schema docs from configs",
		},
	}

	i.command.Args = i.Validate
	i.command.Run = i.Handle

	return i
}

func (c *DbGenerateDocsCommand) Handle(_ *cobra.Command, _ []string) {
	c.tables = database.NewRelationshipService().GetTables()
	// schema
	db := database.NewInstance()
	c.schema, _ = db.GetSchema()

	c.WriteDbSchema()
	c.WriteNav()
	c.WriteDbDocs()
}

func (c *DbGenerateDocsCommand) Validate(_ *cobra.Command, _ []string) error {
	return nil
}

func (c *DbGenerateDocsCommand) WriteNav() {
	fmt.Println("> Writing navigation")

	// ./config/db-schema-nav-ordering.yml
	cfg, err := config.GetSchemaNavOrderingCfg()
	if err != nil {
		log.Fatal(err)
	}

	// mkdocs.yml
	mkdocsCfg, err := config.GetMkdocsConfig()
	if err != nil {
		log.Fatal(err)
	}

	// write nav block
	for i, nav := range mkdocsCfg.Nav {
		if len(nav.DatabaseSchema) > 0 {
			mkdocsCfg.Nav[i].DatabaseSchema = cfg
		}
	}

	// yaml
	var newConfig bytes.Buffer
	yamlEncoder := yaml.NewEncoder(&newConfig)
	yamlEncoder.SetIndent(2)
	err = yamlEncoder.Encode(&mkdocsCfg)
	if err != nil {
		log.Println(err)
	}

	// write
	err = os.WriteFile(config.MkDocsConfigFile, newConfig.Bytes(), 755)
	if err != nil {
		log.Println(err)
	}

	fmt.Printf("Writing database navigation\n")
	fmt.Printf("Wrote [%v]\n", config.MkDocsConfigFile)
}

func (c *DbGenerateDocsCommand) WriteDbSchema() {
	fmt.Println("> Writing database schema")

	// yaml
	configYaml, err := config.GetDbSchemaConfig()
	if err != nil {
		log.Fatal(err)
	}

	for _, row := range c.schema {
		for _, entry := range row {
			if _, ok := configYaml[entry.Table]; !ok {
				configYaml[entry.Table] = make(map[interface{}]map[interface{}]interface{}, 0)
			}
			if _, ok := configYaml[entry.Table][entry.Column]; !ok {
				configYaml[entry.Table][entry.Column] = make(map[interface{}]interface{}, 0)
			}

			configYaml[entry.Table][entry.Column]["dataType"] = entry.DataType

			pos, _ := strconv.Atoi(entry.OrdinalPosition)
			configYaml[entry.Table][entry.Column]["ordinal_position"] = pos
			configYaml[entry.Table][entry.Column]["nullable"] = entry.IsNullable
			if _, ok := configYaml[entry.Table][entry.Column]["description"]; !ok {
				configYaml[entry.Table][entry.Column]["description"] = ""
			}
		}
	}

	// yaml
	var newConfig bytes.Buffer
	yamlEncoder := yaml.NewEncoder(&newConfig)
	yamlEncoder.SetIndent(2)
	err = yamlEncoder.Encode(&configYaml)
	if err != nil {
		log.Println(err)
	}

	// write
	err = os.WriteFile(config.DbSchemaReferenceConfigFile, newConfig.Bytes(), 755)
	if err != nil {
		log.Println(err)
	}

	fmt.Printf("Wrote [%v]\n", config.DbSchemaReferenceConfigFile)
}

func (c *DbGenerateDocsCommand) WriteDbDocs() {
	fmt.Println("> Writing database documentation")

	navTables := config.GetMkDocsDbSchemaNavTables()

	// schema
	db := database.NewInstance()
	tables, _ := db.GetSchemaTableNames()

	// print tables that have not been defined
	for _, table := range tables {
		if _, ok := navTables[table]; !ok {
			fmt.Printf(
				"New table [%v] needs to be placed in [%v] before it will be generated\n",
				table,
				config.DbSchemaNavOrderConfigFile,
			)
		}
	}

	// db schema config
	schemaConfig, err := config.GetDbSchemaConfig()
	if err != nil {
		log.Fatal(err)
	}

	// sections
	wroteFiles := 0
	sections := config.GetMkDocsDbSchemaSections()
	for _, pages := range sections {
		for _, page := range pages {
			table := path.Base(page)
			table = strings.ReplaceAll(table, ".md", "")

			if _, ok := navTables[table]; ok {
				mdPath := fmt.Sprintf("./docs/%v", page)

				// write markdown
				markdown := c.BuildMarkdownForTable(table, schemaConfig)
				if len(markdown) > 0 {

					err := os.MkdirAll(path.Dir(mdPath), os.ModePerm)
					if err != nil {
						log.Println(err)
					}

					err = os.WriteFile(mdPath, []byte(markdown), os.ModePerm)
					if err != nil {
						log.Println(err)
					}
					wroteFiles++
				}

				if len(markdown) == 0 {
					//fmt.Printf("Table [%v] doesn't exist, deleting\n", table)
					//err = os.Remove(mdPath)
					//if err != nil {
					//	log.Println(err)
					//}
				}

			}
		}
	}

	fmt.Printf("Wrote [%v] files\n", wroteFiles)
}

type TableFieldEntry struct {
	column      string
	dataType    string
	nullable    string
	description string
}

func (c *DbGenerateDocsCommand) BuildMarkdownForTable(table string, schemaConfig map[interface{}]map[interface{}]map[interface{}]interface{}) string {
	tableFields := make(map[int]TableFieldEntry, 0)
	for cTable, entries := range schemaConfig {
		if cTable == table {
			for columnI, fields := range entries {
				// type cast interface
				// skip cast validation
				column, _ := columnI.(string)
				position, ok := fields["ordinal_position"].(int)
				dataType, _ := fields["dataType"].(string)
				nullable, _ := fields["nullable"].(string)
				description, _ := fields["description"].(string)

				if ok {

					if table == "aa_actions" {
						fmt.Println(table)
						fmt.Println(position)
						fmt.Println(ok)
					}

					tableFields[position] = TableFieldEntry{
						column:      column,
						dataType:    dataType,
						nullable:    nullable,
						description: description,
					}
				}
			}
		}
	}

	diagrams := c.buildMermaidDiagrams(table)
	imageLink := ""
	diagramLink := ""
	diagramsMarkdown := ""
	if len(diagrams) > 0 {
		for _, diagram := range strings.Split(diagrams, "erDiagram") {
			if len(diagram) > 0 {
				mermaid := MermaidLiveJsDiagram{
					Code: fmt.Sprintf("erDiagram%v", diagram),
					Mermaid: struct {
						Theme string `json:"theme"`
					}{Theme: "default"},
					UpdateEditor:  true,
					AutoSync:      true,
					UpdateDiagram: true,
				}

				b, err := json.Marshal(mermaid)
				if err != nil {
					fmt.Println("error:", err)
				}

				encoded := base64.StdEncoding.EncodeToString(b)
				imageLink = fmt.Sprintf("[![](https://mermaid.ink/img/%v)](https://mermaid.ink/img/%v){target=diagrams}", encoded, encoded)
				diagramLink = fmt.Sprintf("[Diagram Edit](https://mermaid.live/edit#%v){target=diagrams-edit}", encoded)
				diagramsMarkdown += fmt.Sprintf("%v\n\n%v\n\n", diagramLink, imageLink)
			}
		}
	}

	// build table
	tableHeader := `| Column | Data Type | Description |
| :--- | :--- | :--- |`

	markdown := fmt.Sprintf("# %v\n", table)

	currentTime := time.Now()
	generatedTime := currentTime.Format("2006.01.02")

	markdown += fmt.Sprintf("\n!!! info\n\tThis page was last generated %v\n", generatedTime)

	// if we have a diagrams, we have relationships
	if len(imageLink) > 0 {
		markdown += fmt.Sprintf("\n## Relationship Diagram(s)\n\n%v", diagramsMarkdown)
		markdown += fmt.Sprintf("\n## Relationships\n\n")
		markdown += `| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
`
		for _, t := range c.tables {
			if t.Name == table {
				for _, relationship := range t.Relationships {
					relationshipTypeString := "relates"
					if relationship.RelationshipType == "1-*" {
						relationshipTypeString = "Has-Many"
					}
					if relationship.RelationshipType == "1-1" {
						relationshipTypeString = "One-to-One"
					}

					markdown += fmt.Sprintf(
						"| %v | %v | %v | %v |\n",
						relationshipTypeString,
						relationship.LocalKey,
						c.buildLinkToTablePage(relationship.RemoteTable),
						relationship.RemoteKey,
					)
				}
			}
		}
		markdown += "\n"

	}

	markdown += fmt.Sprintf("\n## Schema\n\n%v", tableHeader)

	entries := 0
	for i := 0; i <= 1000; i++ {
		if val, ok := tableFields[i]; ok {
			markdown = fmt.Sprintf(
				"%v\n| %v | %v | %v |",
				markdown,
				val.column,
				val.dataType,
				val.description,
			)
			entries++
		}
	}

	if entries == 0 {
		return ""
	}

	markdown = fmt.Sprintf("%v\n\n", markdown)

	return markdown
}

type MermaidLiveJsDiagram struct {
	Code    string `json:"code"`
	Mermaid struct {
		Theme string `json:"theme"`
	} `json:"mermaid"`
	UpdateEditor  bool `json:"updateEditor"`
	AutoSync      bool `json:"autoSync"`
	UpdateDiagram bool `json:"updateDiagram"`
}

func (c *DbGenerateDocsCommand) chunkRelationships(slice []string, chunkSize int) [][]string {
	var chunks [][]string
	for i := 0; i < len(slice); i += chunkSize {
		end := i + chunkSize

		// necessary check to avoid slicing beyond
		// slice capacity
		if end > len(slice) {
			end = len(slice)
		}

		chunks = append(chunks, slice[i:end])
	}

	return chunks
}

func (c *DbGenerateDocsCommand) buildMermaidDiagrams(table string) string {
	//    npc_types ||--o{ PERSON : has-many
	//    npc_types {
	//        string registrationNumber
	//        string make
	//        string model
	//    }
	//    PERSON
	//    PERSON {
	//        string firstName
	//        string lastName
	//        int age
	//    }

	for _, t := range c.tables {

		if t.Name == table {

			// build remote tables
			remoteTables := []string{}
			for _, relationship := range t.Relationships {
				if !containsStringSlice(remoteTables, relationship.RemoteTable) {
					remoteTables = append(remoteTables, relationship.RemoteTable)
				}
			}

			//pp.Printf("Remote Table Count for [%v] count [%v]\n", table, len(remoteTables))
			//pp.Println(remoteTables)

			// add self table

			keys := c.getRelationshipKeysForTable(table)
			keysString := ""
			if len(keys) > 0 {
				for _, key := range keys {
					keysString += fmt.Sprintf(
						"        %v %v\n",
						c.getDataTypeForTableField(table, key),
						key,
					)
				}
			}

			diagrams := ""
			for _, r := range c.chunkRelationships(remoteTables, erdDiagramChunkSize) {
				tableDefinitions := ""
				tableDefinitions += fmt.Sprintf("    %v {\n%v\n    }\n", table, strings.TrimSuffix(keysString, "\n"))

				//pp.Printf("# chunk size [%v]\n", len(r))
				//pp.Println(r)

				//pp.Println("Table Definitions")
				//fmt.Println(tableDefinitions)

				// add remote tables
				for _, remoteTable := range r {
					keys := c.getRelationshipKeysForTable(remoteTable)
					keysString := ""
					if len(keys) > 0 {
						for _, key := range keys {
							keysString += fmt.Sprintf(
								"        %v %v\n",
								c.getDataTypeForTableField(remoteTable, key),
								key,
							)
						}
					}

					tableDefinitions += fmt.Sprintf("    %v {\n%v\n    }\n", remoteTable, strings.TrimSuffix(keysString, "\n"))
				}

				// add relationships
				relationshipsString := ""
				for _, relationship := range t.Relationships {
					if !containsStringSlice(r, relationship.RemoteTable) {
						continue
					}

					relationshipTypeString := "relates"
					if relationship.RelationshipType == "1-*" {
						relationshipTypeString = "Has-Many"
					}
					if relationship.RelationshipType == "1-1" {
						relationshipTypeString = "One-to-One"
					}

					relationshipsString += fmt.Sprintf(
						"    %v ||--o{ %v : %v\n",
						table,
						relationship.RemoteTable,
						relationshipTypeString,
					)
				}

				diagram := fmt.Sprintf("erDiagram\n%v%v\n", tableDefinitions, relationshipsString)
				diagrams += diagram

				//pp.Println("Diagram")
				//fmt.Println(diagram)
			}

			//pp.Println("Diagrams")
			//fmt.Println(diagrams)

			return diagrams
		}
	}
	return ""
}

func (c *DbGenerateDocsCommand) getRelationshipKeysForTable(table string) []string {
	keys := []string{}

	for _, t := range c.tables {
		if t.Name == table {
			for _, key := range t.Keys {
				if !containsStringSlice(keys, key) {
					keys = append(keys, key)
				}
			}
		}
		for _, relationship := range t.Relationships {
			if relationship.RemoteTable == table {
				if !containsStringSlice(keys, relationship.RemoteKey) {
					keys = append(keys, relationship.RemoteKey)
				}
			}
		}
	}

	return keys
}

func (c *DbGenerateDocsCommand) getDataTypeForTableField(table string, key string) string {
	for _, row := range c.schema {
		for _, entry := range row {
			if entry.Table == table && entry.Column == key {
				reg, err := regexp.Compile("[^a-zA-Z]+")
				if err != nil {
					log.Fatal(err)
				}
				return reg.ReplaceAllString(entry.ColumnType, "")
			}
		}
	}

	return "varchar"
}

func (c *DbGenerateDocsCommand) buildLinkToTablePage(table string) string {
	link := table

	err := filepath.Walk(
		"./docs/schema/",
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if strings.Contains(path, fmt.Sprintf("%v.md", table)) {
				newPath := strings.ReplaceAll(path, "docs", "")
				newPath = strings.ReplaceAll(newPath, "\\", "/")
				link = fmt.Sprintf("[%v](../..%v)", table, newPath)
			}

			return nil
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	return link
}
