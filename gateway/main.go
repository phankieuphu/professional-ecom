package main

import (
	"gateway/database"
	"gateway/routers"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		panic("Do not load env")
	}
	// Setup the router
	r := routers.SetupRouters()
	database.GetDb()

	// Start the server
	r.Run(":8080")
}
