package message

// PubrecMessage is that a PUBREC Packet is the response to a PUBLISH Packet with QoS 2.
// It is the second packet of the QoS 2 protoocl exchange.
type PubrecMessage struct {
	// Remaining Length field
	// This is the length of the variable header. For the PUBREC Packet this has the value 2.
	fixedHeader

	// This contains the Packet Identifier from the PUBLISH Packet that is bening
	// acknowledged.
	packetID []byte
}

// SetPacketID sets Packet Identifier
func (p *PubrecMessage) SetPacketID(v []byte) {
	p.packetID = v
}

// PacketID returns Packet Identifier
func (p *PubrecMessage) PacketID() []byte {
	return p.packetID
}
