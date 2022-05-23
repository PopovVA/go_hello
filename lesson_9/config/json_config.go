package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/url"
	"os"
)

type JsonConfig struct {
	Port   int    `json:"port"`
	DbPath string `json:"dbPath"`
	AppId  string `json:"appId"`
	Token  string `json:"token"`
}

func (config *JsonConfig) Print() {
	fmt.Println("Config:")
	fmt.Printf("Port: %v\n", config.Port)
	fmt.Printf("App Id: %v\n", config.AppId)
	fmt.Printf("DB Path: %v\n", config.DbPath)
	fmt.Printf("Token: %v\n", config.Token)
}

func (config *JsonConfig) validate() (err error) {
	url, err := url.Parse(config.DbPath)
	if url.String() == "" {
		return errors.New(err.Error())
	}
	return err
}

func (config *JsonConfig) Read() error {

	file, err := os.Open("./config.json")
	if err != nil {
		log.Fatalf("Can't open the file: %v", err)
	}
	defer func() {
		err := file.Close()
		if err != nil {
			fmt.Printf("Can't close the file: %v", err)
		}
	}()
	err = json.NewDecoder(file).Decode(&config)
	if err != nil {
		fmt.Printf("Can't decode json to the structure: %v", err)
	}

	if err := config.validate(); err != nil {
		return err
	}

	return nil
}
