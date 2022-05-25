package config

import (
	"errors"
	"flag"
	"fmt"
	"net/url"
	"os"
	"strconv"
)

type EnvConfig struct {
	Port   int
	DbPath string
	AppId  string
	Token  string
}

func (config *EnvConfig) Print() {
	fmt.Println("Config:")
	fmt.Printf("Port: %v\n", config.Port)
	fmt.Printf("App Id: %v\n", config.AppId)
	fmt.Printf("DB Path: %v\n", config.DbPath)
	fmt.Printf("Token: %v\n", config.Token)
	fmt.Println()
}

func (config *EnvConfig) validate() error {
	url, err := url.Parse(config.DbPath)
	if url.String() == "" {
		return errors.New("dbPath is empty")
	}
	return err
}

func (config *EnvConfig) Read() error {

	p, err := strconv.Atoi(os.Getenv("port"))
	if err != nil {
		fmt.Printf("Unable to read the port: %v\n", err.Error())
	} else {
		config.Port = p
	}
	config.DbPath = os.Getenv("dbPath")
	config.Token = os.Getenv("token")
	config.AppId = os.Getenv("appId")

	flag.Parse()

	if err := config.validate(); err != nil {
		return err
	}

	return nil
}
