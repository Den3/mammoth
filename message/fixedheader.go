package message

import (
	"encoding/binary"
	"errors"
	"io"
	"math"
)

var (
	// ErrMalformedReaminingLength indicates multiplier is larger than 128 * 128 * 128
	ErrMalformedReaminingLength = errors.New("malformed Remaining Length")

	// ErrRemainingLengthInvalid indicates Remaining Length is less than 1 or larger than 268435455
	ErrRemainingLengthInvalid = errors.New("invalid Remaining Length")
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
func (fh *fixedHeader) SetRemainingLength(l uint32) {
	fh.remainingLength = fh.encodeLength(l)
}

// SetRemainingLength gets Remaining Length including Variable Header and Payload
func (fh *fixedHeader) GetRemainingLength() uint32 {
	return fh.remainingLength
}

func (fh *fixedHeader) Encode(dest []byte) (int, error) {
	dest[0] = fh.controlPacket
	p := 1

	n := binary.PutUvarint(dest, uint64(fh.remainingLength))
	p += n

	return p, nil
}

// EncodeLength implements non normative commented on line 280 - 294
func (fh *fixedHeader) encodeLength(length uint32) uint32 {
	var encodedLength, encodedByte uint32
	for length > 0 {
		if length > 0 {
			encodedLength <<= 8
		}
		encodedByte = length % 128
		length /= 128

		if length > 0 {
			encodedByte |= 128
		}

		encodedLength |= encodedByte
	}

	return encodedLength
}

// decodeLength implements non normative commented on line 296 - 309
func (fh *fixedHeader) decodeLength(r io.Reader) (uint32, error) {
	multiplier := uint32(0)
	value := uint32(0)
	encodedByte := make([]byte, 1)
	limit := uint32(22)
	for {
		_, err := r.Read(encodedByte)
		if err != nil {
			return 0, err
		}

		value |= uint32(encodedByte[0]&0x7F) << multiplier
		if (encodedByte[0] & 0x80) == 0 {
			break
		}
		multiplier += 7
		if multiplier > limit {
			return 0, ErrMalformedReaminingLength
		}
	}
	return value, nil
}

func (fh *fixedHeader) length() uint32 {
	// TODO: avoid calculation here
	return uint32(math.Log2(float64(fh.remainingLength)))/7 + 1 + 1
}
