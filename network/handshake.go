package network

type Handshake struct {
	packet *Packet
}

func NewHandshake(p *Packet) Command {
	h := new(Handshake)
	h.packet = p
	return h
}

func (h *Handshake) IsMalformed() bool {
	//@TODO Check if packet is malformed
	return true
}
