package main

import (
	"io/ioutil"
	"log"

	yaml "gopkg.in/yaml.v2"
)

// Yaml struct of yaml
type Yaml struct {
	Mysql struct {
		User     string `yaml:"user"`
		Host     string `yaml:"host"`
		Password string `yaml:"password"`
		Port     string `yaml:"port"`
		Name     string `yaml:"name"`
	}
	Redis struct {
		Host       string `yaml:"host"`
		Auth       string `yaml:"auth"`
		TWMsgMQKey string `yaml:"TWMsgMQKey"` //"tw:msg" # the MQ key of tweet record in MQ
	}
	Smtp struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Nickname string `yaml:"nickname"`
	}
	App struct {
		Port string `yaml:"port"`
	}
}

var yamlConfig Yaml
var yamlLoaded bool

func LoadConfig() {
	if yamlLoaded {
		return
	}

	// Read yaml file
	yamlFile, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Fatalf("yamlFile.Get err #%v ", err)
	}

	// Unmarshal yaml file
	//var yamlData Yaml
	err = yaml.Unmarshal(yamlFile, &yamlConfig)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	yamlLoaded = true
	// Print yaml data
	//log.Printf("yamlData: %+v\n", yamlConfig)
	//log.Println("yamlConfig loaded")
}

func init() {
	LoadConfig()
}
