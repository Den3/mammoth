package message

// PubrelMessage is that a PUBREL Packet is the response to a PUBREC Packet. It is the
// third packet of the QoS 2 protocol exchange
type PubrelMessage struct {
	// Bits 3, 2, 1 and 0 of the fixed header in the PUBREL Control Packet are reserved
	// and MUST be set to 0, 0, 1 and 0 respectively. The Server MUST treat any other
	// value as malformed and close the Network Connection
	//
	// Remaining Length field
	// This is the length of the variable header. For the PUBREL Packet this has value 2.
	fixedHeader

	// The variable header contains the same Packet Identifier as the PUBREC Packet that is being
	// acknowledged.
	packetID []byte
}

// SetPacketID sets Packet Identifier
func (p *PubrelMessage) SetPacketID(v []byte) {
	p.packetID = v
}

// PacketID returns Packet Identifier
func (p *PubrelMessage) PacketID() []byte {
	return p.packetID
}
