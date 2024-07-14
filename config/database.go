package config

import (
	"fmt"
	"os"
	"strconv"
)

type DatabaseConfig struct {
	Host     string
	Port     int
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

	port, err := strconv.Atoi(DbPort)
	if err != nil {
		return &DatabaseConfig{}, fmt.Errorf("unable to parse DB_PORT")
	}
	return &DatabaseConfig{
		Host:     DbHost,
		Port:     port,
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Database: os.Getenv("DB_DATABASE"),
		Schema:   os.Getenv("DB_SCHEMA"),
	}, nil
}
