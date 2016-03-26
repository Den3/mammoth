package message

import "io"

// Message is the interface defined for all MQTT Control types
type Message interface {
	Write(io.Writer) error
	Read(io.Reader) error
}

// NewMessageByFixedHeader returns Control Packet by Fixed Header
func NewMessageByFixedHeader(fh *fixedHeader) Message {
	return nil
}
