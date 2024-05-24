package config

import (
	"github.com/joho/godotenv"
	"log"
)

func Load() {
	err := godotenv.Load()
	if err != nil {
		// log.Fatal method call os.Exit(1)
		log.Fatal("Error loading .env file")
	}
	log.Println("Successfully loaded .env file")
}
