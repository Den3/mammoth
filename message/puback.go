package message

// PubackMessage is that a PUBACK Packet is the response to a PUBLISH Packet with QoS
// level 1
type PubackMessage struct {
	fixedHeader

	// This contains the Packet Identifier from the PUBLISH Packet that is bening
	// acknowledged.
	packetID []byte
}

// SetPacketID sets Packet Identifier
func (p *PubackMessage) SetPacketID(v []byte) {
	p.packetID = v
}

// PacketID returns Packet Identifier
func (p *PubackMessage) PacketID() []byte {
	return p.packetID
}
