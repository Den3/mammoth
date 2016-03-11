package message

// PubcompMessage is that The PUBCOMP Packet is the response to a PUBREL Packet. It is
// the fourth and final packet of the QoS 2 protocol exchange.
type PubcompMessage struct {
	// Remaining Length field
	// This is the length of the variable header. For the PUBCOMP Packet this has the value 2.
	fixedHeader

	// This contains the Packet Identifier from the PUBLISH Packet that is bening
	// acknowledged.
	packetID []byte
}

// SetPacketID sets Packet Identifier
func (p *PubcompMessage) SetPacketID(v []byte) {
	p.packetID = v
}

// PacketID returns Packet Identifier
func (p *PubcompMessage) PacketID() []byte {
	return p.packetID
}
