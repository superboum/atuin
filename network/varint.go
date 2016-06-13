package network

// Decode a VarInt
// VarInt is a way to encode numbers documented by Google
// By using it, you can save a lot of space with small numbers
func DecodeVarInt(buffer []byte) (int64, uint32) {
	value := int64(0)
	size := uint32(0)

	for buffer[size]>>7 == 1 {
		value |= int64(buffer[size]&0x7F) << (size * 7)
		size++
	}
	value |= int64(buffer[size]&0x7f) << (size * 7)
	size += 1

	return value, size
}
