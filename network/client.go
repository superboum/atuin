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
		//@FIXME Read only first varint and after read a buffer of this size.
		n, err := c.co.Read(buf)
		if err != nil {
			fmt.Println("Unable to read message.", err.Error())
			return
		}
		fmt.Printf("msg: %#x\n", buf[:n])

		//Hardcoded test
		p := NewPacket(buf)
		cmd := c.manager.Dispatch(p)
		fmt.Println(cmd)

	}

}
