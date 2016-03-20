package message

// SubackMessage is that A SUBACK Packet is sent by the Server to the Client to
// confirm receipt and processing of a SUBSCRIBE Packet
//
// A SUBACK Packet contains a list of return codes, this specify the maximum QoS
// level that was granted in each Subscription that was requested by the SUBSCRIBE
type SubackMessage struct {
	fixedHeader

	// The variable header contains the Packet Identifier from the SUBSCRIBE Packet
	// that is being acknowledged.
	packetID []byte

	// The payload contains a list of return codes. Each return code corresponds to
	// a Topic Filter in the SUBSCRIBE Packet being acknowledged. The order of return
	// codes in the SUBACK Packet MUST match the order of Topic Filters in the SUBSCRIBE
	// Packet[MQTT-3.9.3-1]
	//
	// Allowed return codes:
	// 0x00 - Success - Maximum QoS 0
	// 0x01 - Success - Maximum QoS 1
	// 0x02 - Success - Maximum QoS 2
	// 0x80 - Failure
	//
	// SUBACK return codes other than 0x00, 0x01, 0x02 and 0x80 are reserved and MUST NOT
	// be used[MQTT-3.9.3-2]
	qos byte
}

// SetPacketID sets Packet Identifier
func (s *SubackMessage) SetPacketID(v []byte) {
	s.packetID = v
}

// PacketID returns Packet Identifier
func (s *SubackMessage) PacketID() []byte {
	return s.packetID
}

// SetQoS sets QoS
func (s *SubackMessage) SetQoS(v byte) error {
	if (v > 0 && v < 4) || v == 128 {
		s.qos = v
		return nil
	}
	return ErrQoSInvalid
}

// QoS returns QoS
func (s *SubackMessage) QoS() byte {
	return s.qos
}
