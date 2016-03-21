package network

import (
	"encoding/binary"
	"fmt"
)

type Packet struct {
	buffer []byte
	cursor int
}

func NewPacket(buf []byte) *Packet {
	p := new(Packet)
	//@FIXME Should we recopy the content of the buffer
	//or keep it like it ?
	p.buffer = buf
	p.cursor = 2
	return p
}

func (p *Packet) GetLength() byte {
	return p.GetByte(0)
}

func (p *Packet) GetCommand() byte {
	return p.GetByte(1)
}

func (p *Packet) GetByte(pos int) byte {
	return p.buffer[pos]
}

func (p *Packet) ReadUnsignedShort() uint16 {
	value := binary.BigEndian.Uint16(p.buffer[p.cursor:])
	p.cursor += 2
	return value
}

func (p *Packet) ReadVarInt() int64 {
	value, size := binary.Varint(p.buffer[p.cursor:])
	if size == 0 {
		fmt.Println("Error, unable to decoded varint")
	}
	fmt.Printf("decoded size: %d\n", size)
	p.cursor += size
	return value
}

func (p *Packet) ReadString() string {
	//According to wiki.vg size is on 32bit
	size := int(p.ReadVarInt())
	read := string(p.buffer[p.cursor : p.cursor+size])
	p.cursor += size
	return read
}
