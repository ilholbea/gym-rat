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
	DbPort := os.Getenv("DB_PORT")
	DbUser := os.Getenv("DB_USER")
	DbPassword := os.Getenv("DB_PASSWORD")
	DbDatabase := os.Getenv("DB_DATABASE")
	DbSchema := os.Getenv("DB_SCHEMA")

	if DbHost == "" ||
		DbPort == "" ||
		DbUser == "" ||
		DbPassword == "" ||
		DbDatabase == "" ||
		DbSchema == "" {
		return &Config{}, fmt.Errorf("unable to load environment variable")
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
