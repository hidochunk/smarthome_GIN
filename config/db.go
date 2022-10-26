package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func EnvMongoURI() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	if value, ok := os.LookupEnv("mongodb_uri"); ok {
		return value
	}

	log.Fatal("Must set mongodb_uri in .env file")
	return ""
}
