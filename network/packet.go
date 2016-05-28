package network

import (
	"encoding/binary"
)

type Packet struct {
	buffer  []byte
	cursor  int
	length  int64
	command int64
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
	p.cursor = 0
	p.length = p.ReadVarInt()
	p.command = p.ReadVarInt()
	return p
}

//@FIXME Should be a VarInt
func (p *Packet) GetLength() int64 {
	return p.length
}

//@FIXME Should be a VarInt
func (p *Packet) GetCommand() int64 {
	return p.command
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
	value, shift := DecodeVarInt(p.buffer[p.cursor:])
	p.cursor += int(shift)
	return value
}

func (p *Packet) ReadString() string {
	//According to wiki.vg size is on 32bit
	size := int(p.ReadVarInt())
	read := string(p.buffer[p.cursor : p.cursor+size])
	p.cursor += size
	return read
}
