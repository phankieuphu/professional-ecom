package main

import (
	"log"
	"users/config"
	routers "users/routes"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error load env", err)
	}
	r := routers.InitRouters()
	config.GetDb()
	r.Run(":8080")
}
