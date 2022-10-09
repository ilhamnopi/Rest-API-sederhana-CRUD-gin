package main

import "gn/routers"

func main() {
	var PORT = ":8080"

	routers.StartServer().Run(PORT)

}