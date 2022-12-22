package config

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"path"
	"strings"

	"gopkg.in/yaml.v3"
)

type MkDocsCfg struct {
	SiteName string   `yaml:"site_name"`
	Plugins  []string `yaml:"plugins"`
	RepoURL  string   `yaml:"repo_url"`
	SiteUrl  string   `yaml:"site_url"`
	RepoName string   `yaml:"repo_name"`
	EditURI  string   `yaml:"edit_uri"`
	Extra    struct {
		Analytics struct {
			Provider string `yaml:"provider"`
			Property string `yaml:"property"`
		} `yaml:"analytics"`
	} `yaml:"extra"`
	ExtraCSS        []string `yaml:"extra_css"`
	ExtraJavascript []string `yaml:"extra_javascript"`
	CopyRight       string   `yaml:"copyright"`
	Theme           struct {
		Favicon string `yaml:"favicon"`
		Logo    string `yaml:"logo"`
		Font    struct {
			Text string `yaml:"text"`
			Code string `yaml:"code"`
		} `yaml:"font"`
		Features []string `yaml:"features"`
		Name     string   `yaml:"name"`
		Palette  []struct {
			Scheme  string `yaml:"scheme"`
			Primary string `yaml:"primary"`
			Accent  string `yaml:"accent,omitempty"`
		} `yaml:"palette"`
	} `yaml:"theme"`
	MarkdownExtensions []interface{} `yaml:"markdown_extensions"`
	Nav                []struct {
		Play           []interface{}                    `yaml:"Play,omitempty"`
		Server         []interface{}                    `yaml:"Server,omitempty"`
		DatabaseSchema []map[string][]map[string]string `yaml:"Database Schema,omitempty"`
		QuestApi       []map[string][]map[string]string `yaml:"Quest API,omitempty"`
		Changelog      []struct {
			Num2022                string `yaml:"2022,omitempty"`
			Num2021                string `yaml:"2021,omitempty"`
			Num2020                string `yaml:"2020,omitempty"`
			Num2019                string `yaml:"2019,omitempty"`
			Num2018                string `yaml:"2018,omitempty"`
			Num2017                string `yaml:"2017,omitempty"`
			Num2016                string `yaml:"2016,omitempty"`
			Num2015                string `yaml:"2015,omitempty"`
			Num2014                string `yaml:"2014,omitempty"`
			Num2013                string `yaml:"2013,omitempty"`
			Num2012                string `yaml:"2012,omitempty"`
			Num2011                string `yaml:"2011,omitempty"`
			Num2010                string `yaml:"2010,omitempty"`
			Num2009                string `yaml:"2009,omitempty"`
			Num2008                string `yaml:"2008,omitempty"`
			Num2007                string `yaml:"2007,omitempty"`
			Num2006                string `yaml:"2006,omitempty"`
			Num2005                string `yaml:"2005,omitempty"`
			Num2004                string `yaml:"2004,omitempty"`
			Num2003                string `yaml:"2003,omitempty"`
			ChangelogContributions string `yaml:"Changelog Contributions,omitempty"`
		} `yaml:"Changelog,omitempty"`
		Contributing []interface{} `yaml:"Contributing,omitempty"`
	} `yaml:"nav"`
}

func GetMkdocsConfig() (MkDocsCfg, error) {
	fmt.Printf("Reading [%v]\n", MkDocsConfigFile)

	// read config
	cfg, err := os.ReadFile(MkDocsConfigFile)

	// yaml
	var configYaml MkDocsCfg
	err = yaml.Unmarshal(cfg, &configYaml)
	if err != nil {
		return MkDocsCfg{}, err
	}

	return configYaml, nil
}

func WriteMkdocsConfig(config MkDocsCfg) {
	fmt.Printf("Writing [%v]\n", MkDocsConfigFile)

	// yaml
	var newConfig bytes.Buffer
	yamlEncoder := yaml.NewEncoder(&newConfig)
	yamlEncoder.SetIndent(2)
	err := yamlEncoder.Encode(&config)
	if err != nil {
		log.Println(err)
	}

	// post processing for any special yaml objects since go can't marshal / unmarshal properly
	configString := newConfig.String()
	configString += ``

	// write
	err = os.WriteFile(MkDocsConfigFile, []byte(configString), 755)
	if err != nil {
		log.Println(err)
	}
}

func GetMkDocsDatabaseConfig() []map[string][]map[string]string {
	mkdocsCfg, err := GetMkdocsConfig()
	if err != nil {
		log.Fatal(err)
	}

	// write nav block
	for i, nav := range mkdocsCfg.Nav {
		if len(nav.DatabaseSchema) > 0 {
			return mkdocsCfg.Nav[i].DatabaseSchema
		}
	}

	return nil
}

func GetMkDocsQuestApiConfig() []map[string][]map[string]string {
	mkdocsCfg, err := GetMkdocsConfig()
	if err != nil {
		log.Fatal(err)
	}

	// write nav block
	for i, nav := range mkdocsCfg.Nav {
		if len(nav.QuestApi) > 0 {
			return mkdocsCfg.Nav[i].QuestApi
		}
	}

	return nil
}

func GetMkDocsDbSchemaSections() map[string][]string {
	files := make(map[string][]string, 0)
	for _, entry := range GetMkDocsDatabaseConfig() {
		//fmt.Println(entry)
		for section, i := range entry {
			for _, v := range i {
				for _, page := range v {
					files[section] = append(files[section], page)
				}
			}
		}
	}

	return files
}

// GetMkDocsDbSchemaNavTables returns simple list of defined tables in mkdocs nav
// map[string]bool{
// "object":                                 true,
// "login_api_tokens":                       true,
// "keyring":                                true,
// "altadv_vars":                            true,
// "data_buckets":                           true,
// "login_world_servers":                    true,
func GetMkDocsDbSchemaNavTables() map[string]bool {
	var tables = map[string]bool{}
	for _, pages := range GetMkDocsDbSchemaSections() {
		for _, page := range pages {
			table := path.Base(page)
			table = strings.ReplaceAll(table, ".md", "")
			tables[table] = true
		}
	}
	return tables
}
