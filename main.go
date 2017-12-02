package main

import "sword/service"

func main() {
	server := service.NewService(":8888")

	server.Start()
}
