package console

import (
	"bytes"
	"fmt"
	"github.com/EQEmu/eqemu-docs-v2/config"
	"github.com/EQEmu/eqemu-docs-v2/internal/database"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"path"
	"strconv"
	"strings"
)

type DbGenerateDocsCommand struct {
	command *cobra.Command
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

	// schema
	db := database.NewInstance()
	schema, _ := db.GetSchema()

	// yaml
	configYaml, err := config.GetDbSchemaConfig()
	if err != nil {
		log.Fatal(err)
	}

	for _, row := range schema {
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

				// write markdown
				markdown := c.BuildMarkdownForTable(table, schemaConfig)
				if len(markdown) > 0 {
					mdPath := fmt.Sprintf("./docs/%v", page)

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
				position, _ := fields["ordinal_position"].(int)
				dataType, _ := fields["dataType"].(string)
				nullable, _ := fields["nullable"].(string)
				description, _ := fields["description"].(string)

				tableFields[position] = TableFieldEntry{
					column:      column,
					dataType:    dataType,
					nullable:    nullable,
					description: description,
				}
			}
		}
	}

	// build table
	tableHeader := `| Column | Data Type | Description |
| :--- | :--- | :--- |`

	markdown := fmt.Sprintf("# %v\n\n%v", table, tableHeader)

	for i := 0; i <= 1000; i++ {
		if val, ok := tableFields[i]; ok {
			markdown = fmt.Sprintf(
				"%v\n| %v | %v | %v |",
				markdown,
				val.column,
				val.dataType,
				val.description,
			)
		}
	}

	markdown = fmt.Sprintf("%v\n\n", markdown)

	return markdown
}
