package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Get() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("error when loading environment %s", err.Error())
	}

	return &Config{
		Server{
			Port: os.Getenv("PORT"),
		},
		Database{
			Name: os.Getenv("DB_NAME"),
		},
	}
}
