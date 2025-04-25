package main

import (
	"gateway/routers"
	"gateway/config"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		panic("Do not load env")
	}
	// Setup the router
	r := routers.SetupRouters()
	config.GetDb()

	// Start the server
	r.Run(":8080")
}
