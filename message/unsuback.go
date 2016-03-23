package message

// UnsubackMessage is that the UNSUBACK Packet is sent by the Server to the Client
// to confirm receipt of an UNSUBSCRIBE Packet
type UnsubackMessage struct {
	// Remaining Length field
	// this is the length of the variable header. For the UNSUBACK Packet this has
	// value 2.
	fixedHeader

	// The variable header contains the Packet Identifier of the UNSUBSCRIBE Packet
	// is being acknowledged.
	packetID []byte
}

// SetPacketID sets Packet Identifier
func (s *UnsubackMessage) SetPacketID(v []byte) {
	s.packetID = v
}

// PacketID returns Packet Identifier
func (s *UnsubackMessage) PacketID() []byte {
	return s.packetID
}
