package config

import (
	"errors"
	"flag"
	"fmt"
	"net/url"
)

type FlagConfig struct {
	port   *int
	dbPath *string
	appId  *string
	token  *string
}

func (config *FlagConfig) Print() {
	fmt.Println("Config:")
	fmt.Printf("Port: %v\n", *config.port)
	fmt.Printf("App Id: %v\n", *config.appId)
	fmt.Printf("DB Path: %v\n", *config.dbPath)
	fmt.Printf("Token: %v\n", *config.token)
}

func (config *FlagConfig) validate() (err error) {
	url, err := url.Parse(*config.dbPath)
	if url.String() == "" {
		return errors.New("dbPath is empty")
	}
	return err
}

func (config *FlagConfig) Read() error {

	config.port = flag.Int("port", 0, "port for connection to a database")
	config.appId = flag.String("appId", "", "application id")
	config.token = flag.String("token", "", "firebase token")
	config.dbPath = flag.String("DBPath", "", "database path")
	flag.Parse()

	if err := config.validate(); err != nil {
		return err
	}

	return nil
}
