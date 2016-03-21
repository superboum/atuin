package network

import (
	"fmt"
)

type Handshake struct {
	Version     int64
	Address     string
	Port        uint16
	NextState   int64
	isMalformed bool
}

func NewHandshake(p *Packet) Command {
	h := new(Handshake)
	h.parse(p)
	return h
}

func (h *Handshake) IsMalformed() bool {
	return h.isMalformed
}

func (h *Handshake) String() string {
	value := "HANDSHAKE PACKET\n"
	if h.isMalformed {
		return value + "Malformed packet"
	}

	value += fmt.Sprintf("protocol version: %d\n", h.Version)
	value += fmt.Sprintf("address: %s\n", h.Address)
	value += fmt.Sprintf("port: %d\n", h.Port)
	value += fmt.Sprintf("next state: %#x\n", h.NextState)
	return value
}

func (h *Handshake) parse(p *Packet) {
	//@TODO not the correct way to check if packet is malformed
	if p.GetLength() < 10 {
		h.isMalformed = true
		return
	}

	h.Version = p.ReadVarInt()
	h.Address = p.ReadString()
	h.Port = p.ReadUnsignedShort()
	h.NextState = p.ReadVarInt()
}
