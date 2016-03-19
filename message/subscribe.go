package message

// SubscribeMessage is that The SUBSCRIBE Packet is sent from the Client to the Server
// to create one or more Subscriptions. Each Subscription registers a Client's interest
// in one or more Topics. The Server sends PUBLISH Packets to the Client in order to
// forward Application Messages that were published to Topics that match these
// Subscriptions. The SUBSCRIBE Packet also specifies (for each Subscription) the
// maximum QoS with which the Server can send Application Messages to the Client
type SubscribeMessage struct {
	// Bits 3, 2, 1 and 0 of the fixed header of the SUBSCRIBE Control Packet are
	// reserved and MUST be set to 0, 0, 1 and 0 respectively. The Server MUST treat
	// any other value as malformed close the Network Connection[MQTT-3.8.1-1].
	fixedHeader

	// The variable header contains a Packet Identifier. Section 2.3.1 provides more
	// information about Packet identifiers.
	packetID []byte

	// the payload of SUBSCRIBE Packet contains a list of Topic Filters indicating
	// the Topics to which the Client wants to subscribe. The Topic Filters in a
	// SUBSCRIBE packet payload MUST be UTF-8 encoded strings as defined Section
	// 1.5.3[MQTT-3.8.3-1]. A Server SHOULD support Topic filters that contain
	// the wildcard characters defined in Section 4.7.1. If it choose not to support
	// topic filters that contain wildcard characters it MUST reject any Subscription
	// request whose filter contains them[MQTT-3.8.3-2]. Each filter is followed by a
	// byte called the Requested QoS. This gives the maximum QoS level at which the
	// Server can send Application Message to the Client.
	//
	// The payload of a SUBSCRIBE packet MUST contain at least one Topic Filter / QoS
	// pair. A SUBSCRIBE packet with no payload is protocol violation[MQTT-3.8.3-3].
	// See section 4.8 for information about handling errors.
	//
	// The requested maximum QoS field is encoded in the byte following each UTF-8
	// encoded topic name, and these Topic Filter / QoS pairs are packed contiguously
	// The uppder 6 bit of the Requested QoS byte are not used in the current version
	// of the protocol. They are reserved for future use. The Server MUST treat a
	// SUBSCRIBE packet as malformed and close the Network Connection if any of
	// Reserved bits in the payload are non-zero, or QoS is not 0, 1 and 2[MQTT-3.8.3-4].
	topics [][]byte
	qos    []byte
}

// addTopic adds topic
func (s *SubscribeMessage) addTopic(t []byte) {
	s.topics = append(s.topics, t)
}

// Topics returns all topics
func (s *SubscribeMessage) Topics() [][]byte {
	return s.topics
}

// addQoS adds QoS
func (s *SubscribeMessage) addQoS(q byte) error {
	if q > 3 {
		return ErrQoSInvalid
	}
	s.qos = append(s.qos, q)
	return nil
}

// QoS returns all QoS
func (s *SubscribeMessage) QoS() []byte {
	return s.qos
}

// Add adds Topic with QoS
func (s *SubscribeMessage) Add(t []byte, q byte) error {
	s.addTopic(t)
	err := s.addQoS(q)
	if err != nil {
		return err
	}
	return nil
}
