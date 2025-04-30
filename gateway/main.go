package main

import (
	"fmt"
	"gateway/config"
	"gateway/middleware"
	"gateway/routers"
	"gateway/services"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		panic("Do not load env")
	}
	// Setup the router
	r := routers.SetupRouters()

	r.Use(middleware.LoggerMiddleware())
	go services.KafkaConsumer()
	config.GetDb()
	fmt.Println("Server running on port ", 8080)
	// Start the server
	r.Run(":8080")
}
