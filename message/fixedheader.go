package message

import (
	"errors"
)

const (
	// CONNECT type is Client request to connect to Server
	CONNECT = iota + 1

	// CONNACK is Connect acknowledgment
	CONNACK

	// PUBLISH is to publish message
	PUBLISH

	// PUBACK is PUBLISH acknowledgment
	PUBACK

	// PUBREC is PUBLISH received (assured delivery part 1)
	PUBREC

	// PUBREL is PUBLISH release (assured delivery part 2)
	PUBREL

	// PUBCOMP is PUBLISH complete (assured delivery part 3)
	PUBCOMP

	// SUBSCRIBE is Client subscribe request
	SUBSCRIBE

	// SUBACK is SUBSCRIBE acknowledgment
	SUBACK

	// UNSUBSCRIBE is UNSUBSCRIBE request
	UNSUBSCRIBE

	// UNSUBACK is UNSUBSCRIBE acknowledgment
	UNSUBACK

	// PINGREQ is PING request
	PINGREQ

	// PINGREP is PING response
	PINGREP

	// DISCONNECT is Client is disconnecting
	DISCONNECT
)

var (
	// ErrRemainingLengthInvalid indicates Remaining Length is less than 1 or larger than 268435455
	ErrRemainingLengthInvalid = errors.New("invalid Remaining Length")
)

type fixedHeader struct {
	// Control Packet type (7-4) 4 bits only
	// Flags for Control Packet type (3-0) 4 bits only
	controlPacket byte

	// The Remaining Length is the number of bytes remaining within the current packet, including data in the
	// variable header and the payload. The Remaining Length does not include the bytes used to encode the
	// Remaining Length.
	//
	// The Remaining Length is encoded using a variable length encoding scheme which uses a single byte for
	// values up to 127. Larger values are handled as follows. The least significant seven bits of each byte
	// encode the data, and the most significant bit is used to indicate that there are following bytes in the
	// representation. Thus each byte encodes 128 values and a "continuation bit". The maximum number of
	// bytes in the Remaining Length field is 4
	// _ _ _ _ _ _ _ _ ( 8 bits)
	// ↑ indicate if there are following bytes
	remainingLength uint32
}

// SetControlPacketType sets Control Packet Type
func (fh *fixedHeader) SetControlPacketType(cpt byte) {
	// 11110000
	fh.controlPacket = 0xF0 & (cpt << 4)
}

// ControlPacketType returns Control Packet Type
func (fh *fixedHeader) ControlPacketType() byte {
	return fh.controlPacket >> 4
}

// SetControlPacketTypeFlag sets Control Packet Type Flag
func (fh *fixedHeader) SetControlPacketTypeFlag(cptf byte) {
	// 00001111
	fh.controlPacket = 0x0F & cptf
}

// ControlPacketTypeFlag returns Conrol Packet Type Flag
func (fh *fixedHeader) ControlPacketTypeFlag() byte {
	// 00001111
	return fh.controlPacket & 0x0F
}

// SetRemainingLength sets Remaining Length including Variable Header and Payload
func (fh *fixedHeader) SetRemainingLength(l uint) error {
	return errors.New("Should implement encode to Remaining Length")
}
