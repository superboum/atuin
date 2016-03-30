package main

import (
	"atuin/network"
	"fmt"
)

const (
	CONN_HOST = "0.0.0.0"
	CONN_PORT = "25565"
)

func main() {
	fmt.Println("DIORITE FRONT")

	networkManager := network.NewManager()
	networkManager.RegisterCommand(0, network.NewHandshake)
	networkManager.Listen(CONN_HOST, CONN_PORT)
}
