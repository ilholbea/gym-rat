package config

import (
	"os"
)

type Database struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
	Schema   string
}

type Config struct {
	Database
}

func NewConfig() *Config {
	return &Config{
		Database{
			Host: os.Getenv("MYCONSTANT"),
		},
	}
}
