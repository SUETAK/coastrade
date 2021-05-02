package config

import (
	"log"
	"os"

	"gopkg.in/ini.v1"
)

type ConfigList struct {
	ApiKey    string
	ApiSecret string
	User      string
	Host      string
	Password  string
}

var Config ConfigList

func init() {
	config, err := ini.Load("config.ini")
	if err != nil {
		log.Printf("Faild to read file: %v", err)
		os.Exit(1)
	}
	Config = ConfigList{
		config.Section("bitflyer").Key("api_key").String(),
		config.Section("bitflyer").Key("api_secret").String(),
		config.Section("mysql").Key("user").String(),
		config.Section("mysql").Key("host").String(),
		config.Section("mysql").Key("password").String(),
	}
}
