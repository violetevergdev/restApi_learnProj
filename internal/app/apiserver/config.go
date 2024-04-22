package apiserver

import (
	"restAPI/internal/app/store"
)

type Config struct {
	BinAddr  string `toml: "bing_addr"`
	LogLevel string `toml: "log_level"`
	Store    *store.Config
}

func NewConfig() *Config {
	//set default value
	return &Config{
		BinAddr:  ":8080",
		LogLevel: "debug",
		Store:    store.NewConfig(),
	}
}
