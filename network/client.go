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
	}

}
