package config

import (
	"fmt"
	"os"
)

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
	Schema   string
}

func GetDatabaseConfig() (*DatabaseConfig, error) {
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
		return &DatabaseConfig{}, fmt.Errorf("unable to load environment variable")
	}

	return &DatabaseConfig{
		Host:     DbHost,
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Database: os.Getenv("DB_DATABASE"),
		Schema:   os.Getenv("DB_SCHEMA"),
	}, nil
}
