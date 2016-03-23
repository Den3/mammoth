package message

// DisconnectMessage is that the DISCONNECT Packet is the final Control Packet sent
// from the Client to the Server. It indicates that the Client is disconnecting cleanly.
type DisconnectMessage struct {
	// The Server MUST validate that reserved bits are set to zero and disconnect the
	// Client if they are not zero[MQTT-3.14.1-1].
	fixedHeader
}
