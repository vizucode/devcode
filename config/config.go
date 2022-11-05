package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	MYSQL_USER     string
	MYSQL_HOST     string
	MYSQL_PORT     int
	MYSQL_DBNAME   string
	MYSQL_PASSWORD string
}

func GetConfig() *AppConfig {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	MYSQL_PORT, err := strconv.Atoi(os.Getenv("MYSQL_PORT"))
	if err != nil {
		log.Fatal(err.Error())
	}

	return &AppConfig{
		MYSQL_DBNAME:   os.Getenv("MYSQL_DBNAME"),
		MYSQL_HOST:     os.Getenv("MYSQL_HOST"),
		MYSQL_PORT:     MYSQL_PORT,
		MYSQL_USER:     os.Getenv("MYSQL_USER"),
		MYSQL_PASSWORD: os.Getenv("MYSQL_PASSWORD"),
	}
}
