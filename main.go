package main

import (
	"github.com/joho/godotenv"
	"log"
	"msa-auth/api"
	"os"
)

func main() {
	errEnv := godotenv.Load()
	if errEnv != nil {
		log.Panicf("Error loading .env file\n%v", errEnv)
	}

	release := os.Getenv("Release")
	allowOrigin := os.Getenv("AllowOrigin")

	api.RunServer(release == "true", allowOrigin)
}
