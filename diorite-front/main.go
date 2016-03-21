package main

import (
	"diorite/network"
	"fmt"
)

const (
	CONN_HOST = "0.0.0.0"
	CONN_PORT = "25565"
)

func main() {
	fmt.Println("DIORITE FRONT")

	networkManager := network.NewManager()
	networkManager.Listen(CONN_HOST, CONN_PORT)
}
