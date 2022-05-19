package config

import (
	"errors"
	"flag"
	"fmt"
	"net/url"
	"os"
	"strconv"
)

type Configer interface {
	Read() error
	validate() error
	Print()
}

type FlagConfig struct {
	port    *int
	db_path *string
	app_id  *string
	token   *string
}

func (config *FlagConfig) Print() {
	fmt.Println("Config:")
	fmt.Printf("Port: %v\n", *config.port)
	fmt.Printf("App Id: %v\n", *config.app_id)
	fmt.Printf("DB Path: %v\n", *config.db_path)
	fmt.Printf("Token: %v\n", *config.token)
}

func (config *FlagConfig) validate() (err error) {
	url, err := url.Parse(*config.db_path)
	if url.String() == "" {
		return errors.New("db_path is empty")
	}
	return err
}

func (config *FlagConfig) Read() error {

	config.port = flag.Int("port", 0, "port for connection to a database")
	config.app_id = flag.String("app_id", "", "application id")
	config.token = flag.String("token", "", "firebase token")
	config.db_path = flag.String("db_path", "", "database path")
	flag.Parse()

	if err := config.validate(); err != nil {
		return err
	}

	return nil
}

type EnvConfig struct {
	port    int
	db_path string
	app_id  string
	token   string
}

func (config *EnvConfig) Print() {
	fmt.Println("Config:")
	fmt.Printf("Port: %v\n", config.port)
	fmt.Printf("App Id: %v\n", config.app_id)
	fmt.Printf("DB Path: %v\n", config.db_path)
	fmt.Printf("Token: %v\n", config.token)
}

func (config *EnvConfig) validate() error {
	url, err := url.Parse(config.db_path)
	if url.String() == "" {
		return errors.New("db_path is empty")
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
	config.token = os.Getenv("db_path")
	config.token = os.Getenv("token")
	config.app_id = os.Getenv("app_id")

	flag.Parse()

	if err := config.validate(); err != nil {
		return err
	}

	return nil
}
