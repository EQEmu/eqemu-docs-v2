package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

func GetDbRelationShipsConfig() (map[string][]string , error) {
	fmt.Printf("Reading [%v]\n", DbRelationshipsConfigFile)

	// read config
	cfg, err := os.ReadFile(DbRelationshipsConfigFile)

	// yaml
	var configYaml map[string][]string
	err = yaml.Unmarshal(cfg, &configYaml)
	if err != nil {
		return nil, err
	}

	return configYaml, nil
}
