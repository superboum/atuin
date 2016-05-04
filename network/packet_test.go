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

	p = NewPacket([]byte{0x05, 0x01, 0xAC, 0x02})
	if p.ReadVarInt() != 300 {
		t.Error("Failed to decode 2 bytes VarInt")
	}
}
