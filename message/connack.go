package message

// ConnackMessage is that the CONNACK Packet is the packet sent by the Server in
// response to a CONNECT Packet received from a Client. The first packet sent
// from the Server to the Client MUST be a CONNACK Packet [MQTT-3.2.0-1]
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
	// [MQTT-3.2.2-3].
	connectAckFlags byte

	// The values for the one byte unsigned Connect Return code field are listed in
	// Table 3.1 - Connect Return code values. If a well formed CONNECT Packet is
	// received by the Server, but the Server is unable to process it for some
	// reason, then the Server SHOULD attempt to send a CONNACK packet containing
	// the appropriate non-zero Connect return code from this table. If a server
	// sends a CONNACK packet cotaining a non-zero return code it MUST then close
	// the Network Connection [MQTT-3.2.2-5].
	// -----------------------------------------------------------------------------
	// | Value | Return Code Response      |           Description                 |
	// -----------------------------------------------------------------------------
	// |   0   | 0x00 Connection Accepted  | Connection accepted                   |
	// -----------------------------------------------------------------------------
	// |       | 0x01 Connection Refused,  | The Server does not support the level |
	// |   1   | unacceptable protocol     | of the MQTT protocol requested by the |
	// |       | version                   | Client                                |
	// -----------------------------------------------------------------------------
	// |   2   | 0x02 Connection Refused,  | The Client identifier is correct UTF-8|
	// |       | identifier rejected       | but not allowed by the Server         |
	// -----------------------------------------------------------------------------
	// |   3   | 0x03 Connection Refused,  | The Network Connection has been made  |
	// |       | Server unavailable        | but the MQTT service is unavailable   |
	// -----------------------------------------------------------------------------
	// |   4   | 0x04 Connection Refused,  | The data in the user name or password |
	// |       | bad user name or password | is malformed                          |
	// -----------------------------------------------------------------------------
	// |   5   | 0x05 Connection Refused,  | The Client is not authorized to       |
	// |       | not authorized            | connect                               |
	// -----------------------------------------------------------------------------
	// | 6-255 |                           | Reserved for future use               |
	// -----------------------------------------------------------------------------
	// If none of the return codes listed in Table 3.1 - Connect Return code values
	// are deemed applicable, then the Server MUST close the Network Connection
	// without sending a CONNACK [MQTT-3.2.2-6]
	connectReturnCode byte
}

// SetSessionPresent actives Session Present
func (c *ConnackMessage) SetSessionPresent(active bool) {
	if active {
		// 00000001
		c.connectAckFlags |= 0x01
	} else {
		// 11111110
		c.connectAckFlags &= 0xFE
	}
}

// SessionPresent returns Session Present
func (c *ConnackMessage) SessionPresent() byte {
	return c.connectAckFlags & 0x01
}
