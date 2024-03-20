package config

import (
	"fmt"
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
	DbConfig Database
}

func NewConfig() (*Config, error) {

	DbHost := os.Getenv("DB_HOST")
	if DbHost == "" {
		return &Config{}, fmt.Errorf("unable to load db_host")
	}

	return &Config{
		DbConfig: Database{
			Host:     DbHost,
			Port:     os.Getenv("DB_PORT"),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			Database: os.Getenv("DB_DATABASE"),
			Schema:   os.Getenv("DB_SCHEMA"),
		},
	}, nil
}
