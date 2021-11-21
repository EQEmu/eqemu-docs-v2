package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

func GetSchemaNavOrderingCfg() ([]map[string][]map[string]string , error) {
	fmt.Printf("Reading [%v]\n", DbSchemaNavOrderConfigFile)

	// read config
	cfg, err := os.ReadFile(DbSchemaNavOrderConfigFile)

	// yaml
	var configYaml []map[string][]map[string]string
	err = yaml.Unmarshal(cfg, &configYaml)
	if err != nil {
		return nil, err
	}

	return configYaml, nil
}
