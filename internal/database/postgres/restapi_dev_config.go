package db_restapi_dev

import (
	"log"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Host     string `toml:"host"`
	Port     int    `toml:"port"`
	User     string `toml:"user"`
	Password string `toml:"password"`
	DBname   string `toml:"dbname"`
}

func NewConfig() *Config {
	var config Config
    _, err := toml.DecodeFile("configs/dbAuth.toml", &config)
    if err != nil {
        log.Fatalf("Error: open db config file: \n%v ", err)
    }
    return &config
}