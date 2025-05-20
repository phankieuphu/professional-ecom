package main

import (
	"log"
	"net"

	config "github.com/phankieuphu/ecom-user/configs"
	user "github.com/phankieuphu/ecom-user/gen/user/v1"
	"github.com/phankieuphu/ecom-user/internal/services"
	"google.golang.org/grpc"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error load env %v", err)
	}

	config.GetDb()

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpServer := grpc.NewServer()
	user.RegisterUserServer(grpServer, &services.UserService{})

	log.Println("gRPC server listening on port 50051...")
	if err := grpServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
