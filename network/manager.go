package network

import (
	"fmt"
	"net"
	"os"
)

type Manager struct {
	constructors map[byte]func(*Packet) Command
}

func NewManager() *Manager {
	m := new(Manager)
	m.constructors = make(map[byte]func(*Packet) Command)
	return m
}

func (m *Manager) RegisterCommand(cmd byte, fn func(*Packet) Command) {
	m.constructors[cmd] = fn
}

func (m *Manager) Dispatch(p *Packet) Command {
	if fn, ok := m.constructors[p.GetCommand()]; ok {
		return fn(p)
	}
	return nil
}

func (m *Manager) Listen(host string, port string) {
	ln, err := net.Listen("tcp", host+":"+port)

	if err != nil {
		fmt.Println("Unable to listen on port "+port+".", err.Error())
		os.Exit(1)
	}

	defer ln.Close()
	fmt.Println("Listening on " + host + ":" + port)

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Unable to accept connection.", err.Error())
		}
		defer conn.Close()

		client := NewClient(conn, m)
		go client.HandleRequest()
	}
}
