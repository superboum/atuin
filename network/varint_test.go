package network

import (
	"testing"
)

func TestDecodeVarInt(t *testing.T) {
	res, size := DecodeVarInt([]byte{0x03})
	if res != 3 || size != 1 {
		t.Error("Failed to decode single byte VarInt")
	}

	res, size = DecodeVarInt([]byte{0xAC, 0x02})
	if res != 300 || size != 2 {
		t.Error("Failed to decode 2 bytes VarInt")
	}
}
