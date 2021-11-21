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
	"path/filepath"
	"strconv"
	"strings"
)

type SyncDbSchemaConfigFromFilesCommand struct {
	command *cobra.Command
}

func (c *SyncDbSchemaConfigFromFilesCommand) Command() *cobra.Command {
	return c.command
}

func NewSyncDbSchemaConfigFromFilesCommand() *SyncDbSchemaConfigFromFilesCommand {
	i := &SyncDbSchemaConfigFromFilesCommand{
		command: &cobra.Command{
			Use:   "doc:sync-db-schema-config-from-files",
			Short: "Synchronizes database-schema-reference.yml from markdown files (inverse)",
		},
	}

	i.command.Args = i.Validate
	i.command.Run = i.Handle

	return i
}

func (c *SyncDbSchemaConfigFromFilesCommand) Handle(_ *cobra.Command, _ []string) {
	// schema
	db := database.NewInstance()
	schema, _ := db.GetSchema()

	// read config
	cfg, err := os.ReadFile(config.DbSchemaReferenceConfigFile)

	// yaml
	var configYaml map[interface{}]map[interface{}]map[interface{}]interface{}
	err = yaml.Unmarshal(cfg, &configYaml)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	// walk existing files
	err = filepath.Walk("docs/schema/categories",
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			dir := filepath.Dir(path)
			parent := filepath.Base(dir)
			fileName := filepath.Base(path)
			fileNameNoExt := strings.ReplaceAll(fileName, ".md", "")
			categoryNameHuman := strings.ReplaceAll(parent, "-", " ")
			categoryNameHuman = strings.Title(categoryNameHuman)

			//fmt.Printf("dir [%v] file [%v] cat [%v]\n", parent, fileNameNoExt, categoryNameHuman)

			// open file
			if strings.Contains(path, ".md") {
				mdFile, err := os.ReadFile(path)
				if err != nil {
					log.Fatal(err)
				}

				markdown := string(mdFile)

				// loop through lines
				for _, line := range strings.Split(markdown, "\n") {
					if strings.Contains(line, "|") {
						cols := strings.Split(line, "|")
						if len(cols) > 1 {
							col := strings.TrimSpace(cols[1])
							dataType := strings.TrimSpace(cols[2])
							desc := strings.TrimSpace(cols[3])

							//fmt.Println(line)
							//fmt.Printf("[%v] [%v] [%v]\n", col, dataType, desc)

							// write new values
							if _, ok := configYaml[fileNameNoExt][col]; ok {
								configYaml[fileNameNoExt][col]["dataType"] = dataType
								configYaml[fileNameNoExt][col]["description"] = desc

								for _, row := range schema {
									for _, entry := range row {
										if entry.Table == fileNameNoExt && entry.Column == col {
											pos, _ := strconv.Atoi(entry.OrdinalPosition)
											configYaml[fileNameNoExt][col]["ordinal_position"] = pos
										}
									}
								}
							}
						}
					}
				}
			}

			return nil
		})
	if err != nil {
		log.Println(err)
	}

	// yaml
	var newConfig bytes.Buffer
	yamlEncoder := yaml.NewEncoder(&newConfig)
	yamlEncoder.SetIndent(2) // this is what you're looking for
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

func (c *SyncDbSchemaConfigFromFilesCommand) Validate(_ *cobra.Command, _ []string) error {
	return nil
}
