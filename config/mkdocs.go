package config

import (
	"gopkg.in/yaml.v3"
	"os"
)

type MkDocsCfg struct {
	SiteName string   `yaml:"site_name"`
	Plugins  []string `yaml:"plugins"`
	RepoURL  string   `yaml:"repo_url"`
	RepoName string   `yaml:"repo_name"`
	EditURI  string   `yaml:"edit_uri"`
	ExtraCSS []string `yaml:"extra_css"`
	Theme    struct {
		Font struct {
			Text string `yaml:"text"`
			Code string `yaml:"code"`
		} `yaml:"font"`
		Features []string `yaml:"features"`
		Name     string   `yaml:"name"`
		Palette  []struct {
			Media   string `yaml:"media"`
			Scheme  string `yaml:"scheme"`
			Primary string `yaml:"primary"`
			Accent  string `yaml:"accent,omitempty"`
			Toggle  struct {
				Icon string `yaml:"icon"`
				Name string `yaml:"name"`
			} `yaml:"toggle"`
		} `yaml:"palette"`
	} `yaml:"theme"`
	MarkdownExtensions []interface{} `yaml:"markdown_extensions"`
	Nav                []struct {
		PlayClient     []interface{} `yaml:"Play (Client),omitempty"`
		Server         []interface{} `yaml:"Server,omitempty"`
		DatabaseSchema []interface{} `yaml:"Database Schema,omitempty"`
		Changelog      []struct {
			Num2003 string `yaml:"2003,omitempty"`
			Num2004 string `yaml:"2004,omitempty"`
			Num2005 string `yaml:"2005,omitempty"`
			Num2006 string `yaml:"2006,omitempty"`
			Num2007 string `yaml:"2007,omitempty"`
			Num2008 string `yaml:"2008,omitempty"`
			Num2009 string `yaml:"2009,omitempty"`
			Num2010 string `yaml:"2010,omitempty"`
			Num2011 string `yaml:"2011,omitempty"`
			Num2012 string `yaml:"2012,omitempty"`
			Num2013 string `yaml:"2013,omitempty"`
			Num2014 string `yaml:"2014,omitempty"`
			Num2015 string `yaml:"2015,omitempty"`
			Num2016 string `yaml:"2016,omitempty"`
			Num2017 string `yaml:"2017,omitempty"`
			Num2018 string `yaml:"2018,omitempty"`
			Num2019 string `yaml:"2019,omitempty"`
			Num2020 string `yaml:"2020,omitempty"`
			Num2021 string `yaml:"2021,omitempty"`
		} `yaml:"Changelog,omitempty"`
	} `yaml:"nav"`
}

func GetMkdocsConfig() (MkDocsCfg, error) {
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
