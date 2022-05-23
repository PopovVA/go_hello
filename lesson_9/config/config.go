package config

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
	"strconv"
)

type Configer interface {
	Read() error
	validate() error
	Print()
}

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
