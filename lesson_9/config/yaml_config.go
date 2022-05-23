package config

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/url"

	"gopkg.in/yaml.v3"
)

type YamlConfig struct {
	Port   int    `yaml:"port"`
	DbPath string `yaml:"dbPath"`
	AppId  string `yaml:"appId"`
	Token  string `yaml:"token"`
}

func (config *YamlConfig) Print() {
	fmt.Println("Config:")
	fmt.Printf("Port: %v\n", config.Port)
	fmt.Printf("App Id: %v\n", config.AppId)
	fmt.Printf("DB Path: %v\n", config.DbPath)
	fmt.Printf("Token: %v\n", config.Token)
	fmt.Println()
}

func (config *YamlConfig) validate() (err error) {
	url, err := url.Parse(config.DbPath)
	if url.String() == "" {
		return errors.New(err.Error())
	}
	return err
}

func (config *YamlConfig) Read() error {
	file, err := ioutil.ReadFile("./config.json")
	if err != nil {
		log.Fatalf("Can't open the file: %v", err)
	}
	// defer func() {
	// err := file.Close()
	// if err != nil {
	// fmt.Printf("Can't close the file: %v", err)
	// }
	// }()
	err = yaml.Unmarshal(file, &config)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	if err := config.validate(); err != nil {
		return err
	}

	return nil
}
