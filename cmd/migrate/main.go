package main

import (
	"github.com/joho/godotenv"
	"log"
	"msa-auth/database"
)

func main() {
	errEnv := godotenv.Load()
	if errEnv != nil {
		log.Panicf("Error loading .env file\n%v", errEnv)
	}

	database.AutoMigrate()
}
