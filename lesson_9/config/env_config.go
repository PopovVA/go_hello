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
	port   int
	dbPath string
	appId  string
	token  string
}

func (config *EnvConfig) Print() {
	fmt.Println("Config:")
	fmt.Printf("Port: %v\n", config.port)
	fmt.Printf("App Id: %v\n", config.appId)
	fmt.Printf("DB Path: %v\n", config.dbPath)
	fmt.Printf("Token: %v\n", config.token)
}

func (config *EnvConfig) validate() error {
	url, err := url.Parse(config.dbPath)
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
		config.port = p
	}
	config.token = os.Getenv("DBPath")
	config.token = os.Getenv("token")
	config.appId = os.Getenv("appId")

	flag.Parse()

	if err := config.validate(); err != nil {
		return err
	}

	return nil
}
