package config

import (
	"os"
)

type Configuration struct {
	DbPort string
	DbUser string
	DbPass string
	DbHost string
	DbName string
	PORT   string
}

func NewConfiguration() *Configuration {
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	port := os.Getenv("PORT")
	return &Configuration{
		DbPort: dbPort,
		DbUser: dbUser,
		DbPass: dbPass,
		DbHost: dbHost,
		DbName: dbName,
		PORT:   port,
	}
}
