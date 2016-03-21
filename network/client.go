package network

import (
	//"bytes"
	"fmt"
	"net"
)

type Client struct {
	co net.Conn
}

func NewClient(co net.Conn) *Client {
	c := new(Client)
	c.co = co
	return c
}

func (c *Client) HandleRequest() {
	buf := make([]byte, 4096)
	for {
		n, err := c.co.Read(buf)
		if err != nil {
			fmt.Println("Unable to read message.", err.Error())
			return
		}
		fmt.Printf("msg: %#x\n", buf[:n])

		//Hardcoded test
		p := NewPacket(buf)
		if p.GetLength() == 15 {
			fmt.Printf("Packet\nsize: %d bytes\n", p.GetLength())
			fmt.Printf("command: %#x\n", p.GetCommand())
			fmt.Printf("protocol version: %d\n", p.ReadVarInt())
			fmt.Printf("address: %s\n", p.ReadString())
			fmt.Printf("port: %d\n", p.ReadUnsignedShort())
			fmt.Printf("next state: %#x\n", p.ReadVarInt())
		}

	}

}
