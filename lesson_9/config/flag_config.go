package config

import (
	"errors"
	"flag"
	"fmt"
	"net/url"
)

type FlagConfig struct {
	Port   *int
	DbPath *string
	AppId  *string
	Token  *string
}

func (config *FlagConfig) Print() {
	fmt.Println("Config:")
	fmt.Printf("Port: %v\n", *config.Port)
	fmt.Printf("App Id: %v\n", *config.AppId)
	fmt.Printf("DB Path: %v\n", *config.DbPath)
	fmt.Printf("Token: %v\n", *config.Token)
	fmt.Println()
}

func (config *FlagConfig) validate() (err error) {
	url, err := url.Parse(*config.DbPath)
	if url.String() == "" {
		return errors.New("dbPath is empty")
	}
	return err
}

func (config *FlagConfig) Read() error {

	config.Port = flag.Int("port", 0, "port for connection to a database")
	config.AppId = flag.String("appId", "", "application id")
	config.Token = flag.String("token", "", "firebase token")
	config.DbPath = flag.String("DBPath", "", "database path")
	flag.Parse()

	if err := config.validate(); err != nil {
		return err
	}

	return nil
}
