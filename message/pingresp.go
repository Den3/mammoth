package message

// PingrespMessage is that a PINGRESP Packet is sent by the Server to the Client in
// response to a PINGREQ Packet. It indicates that the Server is alive.
type PingrespMessage struct {
	fixedHeader
}
