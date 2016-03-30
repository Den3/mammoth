package message

import "io"

// Message is the interface defined for all MQTT Control types
type Message interface {
	Write(io.Writer) error
	Read(io.Reader) error
}

// NewMessage returns Control Packet by Fixed Header
func NewMessage(cpt byte) Message {
	switch cpt {
	case CONNECT:
	case CONNACK:
	case PUBLISH:
	case PUBACK:
	case PUBREC:
	case PUBREL:
	case PUBCOMP:
	case SUBSCRIBE:
	case SUBACK:
	case UNSUBSCRIBE:
	case UNSUBACK:
	case PINGREQ:
	case PINGREP:
	case DISCONNECT:
	default:
		return nil
	}
	return nil
}
