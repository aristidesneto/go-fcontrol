package configs

import (
	"errors"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppName  string
	Database Database
}

type Database struct {
	Name string
	Uri  string
}

var EnvConfig *Config

func LoadConfig() error {
	log.Printf("Carregando configurações...")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	appName := os.Getenv("APP_NAME")
	dbName := os.Getenv("MONGODB_DATABASE")
	dbUri := os.Getenv("MONGODB_URI")

	if appName == "" || dbName == "" || dbUri == "" {
		return errors.New("missing required environment variables")
	}

	EnvConfig = &Config{
		AppName: os.Getenv("APP_NAME"),
		Database: Database{
			Name: os.Getenv("MONGODB_DATABASE"),
			Uri:  os.Getenv("MONGODB_URI"),
		},
	}

	return nil
}
