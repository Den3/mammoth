package message

// The CONNACK Packet is the packet sent by the Server in response to a CONNECT Packet
// received from a Client. The first packet sent from the Server to the Client MUST be
// a CONNACK Packet [MQTT-3.2.0-1]
//
// If the Client does not receive a CONNACK Packet from the Server within a reasonable
// amount of time, the Client SHOULD close the Network Connection. A "reasonable"
// amount of time depends on the type of application and the communications infrastructure

type ConnackMessage struct {
	fixedHeader

	// Connect Acknowledge Flags must set
	// bits 7-1 are reserved and MUST be set to 0
	// bit 0 is the Session Present Flag
	//
	// If the Server accepts a connection with CleanSession set to 1, The Server MUST
	// set Session Present to 0 in the CONNACK packet in addition to setting a zero
	// return code in the CONNACK packet [MQTT-3.2.2-1]
	//
	// If the Server accepts a connection with CleanSession set to 0, the value set in
	// Session Present depends on whether the Server already has stored Session state
	// for the supplied client ID. If the Server has stored Session state, it MUST set
	// Session Present to 1 in the CONNACK packet [MQTT-3.2.2-2]. If the Server does
	// not have stored Session state, it MUST set Session Present to 0 in the CONNACK
	// packet. This is in addition to setting a zero return code in the CONNACK packet
	// [MQTT-3.2.2-3]
	connectAckFlags byte
}
