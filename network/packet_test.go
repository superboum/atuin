package network

import (
	"testing"
)

//@TODO Test empty packet

func TestReadVarInt(t *testing.T) {
	p := NewPacket([]byte{0x02, 0x01, 0x03})
	if p.ReadVarInt() != 3 {
		t.Error("Failed to decode single byte VarInt")
	}

	p = NewPacket([]byte{0x03, 0x01, 0xAC, 0x02})
	if p.ReadVarInt() != 300 {
		t.Error("Failed to decode 2 bytes VarInt")
	}
}

func TestReadString(t *testing.T) {
	p := NewPacket([]byte{0x07, 0x01, 0x05, 0x48, 0x45, 0x4C, 0x4C, 0x4F})
	if p.ReadString() != "HELLO" {
		t.Error("Failed to decode HELLO string")
	}
}
