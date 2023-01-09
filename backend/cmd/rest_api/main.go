package main

import (
	"airport_web_server/internal/rest_api/routes"
	dotenv "github.com/joho/godotenv"
	"log"
	"os"
)

func init() {
	InitEnvVariables()
}

// InitEnvVariables loads the environment variables from the .env file
func InitEnvVariables() {
	err := dotenv.Load()
	if err != nil {
		// we try to load the .env file with the relative path
		err = dotenv.Load("../../.env")
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}
}

func main() {
	router := routes.InitRouter()
	println(os.Getenv("API_PORT"))
	err := router.Run(os.Getenv("API_PORT"))
	if err != nil {
		panic(err)
	}
}
