package main

import "gateway/routers"

func main() {
	// Setup the router
	r := routers.SetupRouters()

	// Start the server
	r.Run(":8080")
}
