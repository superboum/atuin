package network

import (
	"encoding/binary"
)

type Packet struct {
	buffer []byte
	cursor int
}

const (
	SRV_UNDEFINED = 0x00
	SRV_HANDSHAKE = 0x01
)

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
	value := int64(0)
	size := uint32(0)

	for p.buffer[p.cursor]>>7 == 1 {
		value |= int64(p.buffer[p.cursor]&0x7F) << (size * 7)
		p.cursor++
		size++
	}
	value |= int64(p.buffer[p.cursor]&0x7f) << (size * 7)
	p.cursor += 1

	return value
}

func (p *Packet) ReadString() string {
	//According to wiki.vg size is on 32bit
	size := int(p.ReadVarInt())
	read := string(p.buffer[p.cursor : p.cursor+size])
	p.cursor += size
	return read
}
