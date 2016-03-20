package message

// UnsubscribeMessage is that an UNSUBSCRIBE Packet is sent by the Client to the Server,
// to unsubscribe from topics.
type UnsubscribeMessage struct {
	// Bit 3, 2, 1 and 0 of the fixed header of the UNSUBSCRIBE Control Packet are
	// reserved and MUST be set to 0, 0, 1 and 0 respectively. Teh Server MUST treat
	// any other value as malformed and close the Network Connection[MQTT-3.10.1-1].
	fixedHeader

	// The variable header contains a Packet Identifier. Section 2.3.1 provides more
	// information about Packet Identifier.
	packetID []byte

	// The payload for the UNSUBSCRIBE Packet contains the list of Topic Filters that
	// the Client wishes to unsubscribe from. The Topic Filters in a UNSUBSCRIBE packet
	// MUST be UTF-8 encoded strings as defined in Section 1.5.3, packed contiguously
	// [MQTT-3.10.3-1]
	//
	// The Payload of an UNSUBSCRIBE packet MUST contain at least one Topic Filter.
	// An UNSUBSCRIBE packet with no payload is a protocol violation[MQTT-3.10.3-2].
	// See section 4.8 for information about handling errors.
	topics [][]byte
}

// SetPacketID sets Packet Identifier
func (s *UnsubscribeMessage) SetPacketID(v []byte) {
	s.packetID = v
}

// PacketID returns Packet Identifier
func (s *UnsubscribeMessage) PacketID() []byte {
	return s.packetID
}

// AddTopic adds topic
func (s *UnsubscribeMessage) AddTopic(t []byte) {
	s.topics = append(s.topics, t)
}

// Topics returns all topics
func (s *UnsubscribeMessage) Topics() [][]byte {
	return s.topics
}
