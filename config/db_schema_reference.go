package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

func GetDbSchemaConfig() (map[interface{}]map[interface{}]map[interface{}]interface{}, error) {
	fmt.Printf("Reading [%v]\n", DbSchemaReferenceConfigFile)

	// read config
	cfg, err := os.ReadFile(DbSchemaReferenceConfigFile)

	// yaml
	var configYaml map[interface{}]map[interface{}]map[interface{}]interface{}
	err = yaml.Unmarshal(cfg, &configYaml)
	if err != nil {
		return nil, err
	}

	return configYaml, nil
}
