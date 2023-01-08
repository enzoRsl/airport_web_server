package main

import (
	"airport_web_server/internal/rest_api/routes"
)

func main() {
	router := routes.InitRouter()
	err := router.Run(":8086")
	if err != nil {
		panic(err)
	}
}
