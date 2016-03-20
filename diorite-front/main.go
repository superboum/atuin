package main

import (
	"diorite/network"
	"fmt"
	"net"
	"os"
)

const (
	CONN_HOST = "0.0.0.0"
	CONN_PORT = "25565"
)

func main() {
	fmt.Println("DIORITE FRONT")
	ln, err := net.Listen("tcp", CONN_HOST+":"+CONN_PORT)

	if err != nil {
		fmt.Println("Unable to listen on port 25565.", err.Error())
		os.Exit(1)
	}

	defer ln.Close()
	fmt.Println("Listening on " + CONN_HOST + ":" + CONN_PORT)

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Unable to accept connection.", err.Error())
		}
		defer conn.Close()

		client := network.NewClient(conn)
		go client.HandleRequest()
	}
}
