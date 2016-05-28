package network

import (
	//"bytes"
	"fmt"
	"net"
)

type Client struct {
	co      net.Conn
	manager *Manager
}

func NewClient(co net.Conn, m *Manager) *Client {
	c := new(Client)
	c.co = co
	c.manager = m
	return c
}

func (c *Client) HandleRequest() {
	buf := make([]byte, 4096)
	for {
		bytesRead, err := c.co.Read(buf)
		if err != nil {
			fmt.Println("Unable to read message.", err.Error())
			return
		}
		packetCursor := 0

		// @FIXME handle if packet is split
		for packetCursor < bytesRead {
			fmt.Printf("msg: %#x\n", buf[packetCursor:bytesRead])

			//Hardcoded test
			p := NewPacket(buf)
			cmd := c.manager.Dispatch(p)
			fmt.Println(cmd)

			packetCursor += int(p.GetFullLength())

			if packetCursor > bytesRead {
				fmt.Printf("Reading error. Read too many bytes. Read: %d Limit: %d...\n", packetCursor, bytesRead)
			}
		}
	}

}
