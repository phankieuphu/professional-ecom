package main

import (
	routers "users/routes"
)

func main() {
	r := routers.InitRouters()

	r.Run(":8001")
}
