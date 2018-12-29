package message

import (
	"errors"
	"regexp"
)

var (
	ErrQoSInvalid            = errors.New("invalid QoS value")
	ErrClientIdLengthInvalid = errors.New("invalid ClientId length")
	ErrClientIdInvalid       = errors.New("invalid ClientId")
)

// After a Network Connection is established by a Client to a Server, the first Packet
// sent from the Client to the Server MUST be a CONNECT Packet [MQTT-3.1.0-1].
//
// A Client can only send the CONNECT Packet once over a Network Connection. The Server
// MUST process a second CONNECT Packet sent from a Client as a protocol violation and
// disconnect the Client [MQTT-3.1.0-2].  See section 4.8 for information about
// handling errors.
type ConnectMessage struct {
	fixedHeader

	// The Protocol Name is a UTF-8 encoded string that represents the protocol name
	// “MQTT”, capitalized as shown. The string, its offset and length will not be
	// changed by future versions of the MQTT specification.
	protocolName []byte

	// The 8 bit unsigned value that represents the revision level of the protocol
	// used by the Client. The value of the Protocol Level field for the version 3.1.1
	// of the protocol is 4 (0x04).
	protocolLevel byte

	// bit 7 for User Name Flag
	// bit 6 for Password Flag
	// bit 5 for Will Retain
	// bit 4 - 3 for Will QoS (Quality of Service)
	// bit 2 for Will Flag
	// bit 1 for Clean Session
	// bit 0 for Reserved
	connectFlags byte

	// The Keep Alive is a time interval measured in seconds. Expressed as a 16-bit
	// word, it is the maximum time interval that is permitted to elapse between the
	// point at which the Client finishes transmitting one Control Packet and the point
	// it starts sending the next
	keepAlive uint16

	// The payload of the CONNECT Packet contains one or more length-prefixed
	// fields, whose presence is determined by the flags in the variable header.
	// These fields, if present, MUST appear in the order Client Identifier, Will
	// Topic, Will Message, User Name, Password [MQTT-3.1.3-1].

	// Each Client connecting to the Server has a unique ClientId. The ClientId Must
	// be used by Clients and by Servers to identify state that they hold relating to
	// this MQTT Session between the Client and the Server [MQTT-3.1.3-2]
	//
	// The ClientId MUST be present and MUST be the first field in the CONNECT packet payload
	//
	// The Server Must allow ClientIds which are between 1 and 23 UTF-8 encoded bytes in
	// length, and that contain only the characters "[0-9a-zA-z]" [MQTT-3.1.3-5]
	clientId    []byte
	willTopic   []byte
	willMessage []byte
	userName    []byte
	password    []byte
}

// NewConnectMessage returns pointer of ConnectMessage
func NewConnectMessage() *ConnectMessage {
	c := &ConnectMessage{}
	c.SetControlPacketType(CONNECT)
	return c
}

// SetProtocolName sets Protocol Name to "MQTT" by default
func (c *ConnectMessage) SetProtocolName() {
	content := "MQTT"
	c.protocolName = make([]byte, len(content)+2)
	c.protocolName[1] = byte(len(content))
	for i := 2; i < 6; i++ {
		c.protocolName[i] = content[i-2]
	}
}

// SetProtocolLevel sets Protocol Level to 4 by default
func (c *ConnectMessage) SetProtocolLevel() {
	c.protocolLevel = 4
}

// SetUserNameFlag sets User Name Flag
func (c *ConnectMessage) SetUserNameFlag(active bool) {
	if active {
		// 10000000
		c.connectFlags |= 0x80
	} else {
		// 01111111
		c.connectFlags &= 0x7F
	}
}

// UserNameFlag returns User Name Flag
func (c *ConnectMessage) UserNameFlag() byte {
	return (c.connectFlags >> 7) & 0x1
}

// SetPasswordFlag sets Password Flag
func (c *ConnectMessage) SetPasswordFlag(active bool) {
	if active {
		// 01000000
		c.connectFlags |= 0x40
	} else {
		// 10111111
		c.connectFlags &= 0xBF
	}
}

// PasswordFlag returns Password Flag
func (c *ConnectMessage) PasswordFlag() byte {
	return (c.connectFlags >> 6) & 0x1
}

// SetWillRetain sets Will Retain
func (c *ConnectMessage) SetWillRetain(active bool) {
	if active {
		// 00100000
		c.connectFlags |= 0x20
	} else {
		// 11011111
		c.connectFlags &= 0xCF
	}
}

// WillRetain returns Will Retain
func (c *ConnectMessage) WillRetain() byte {
	return (c.connectFlags >> 5) & 0x1
}

// SetWillQoS sets Will QoS
func (c *ConnectMessage) SetWillQoS(quality byte) error {
	if quality > 3 {
		return ErrQoSInvalid
	}
	// 11100111
	c.connectFlags = (c.connectFlags & 0xE7) | (quality << 3)
	return nil
}

// WillQoS returns Will QoS
func (c *ConnectMessage) WillQoS() byte {
	return (c.connectFlags >> 3) & 0x3
}

// SetWillFlag sets Will Flag
func (c *ConnectMessage) SetWillFlag(active bool) {
	if active {
		// 00000100
		c.connectFlags |= 0x04
	} else {
		// 11111011
		c.connectFlags &= 0xFB
	}
}

// WillFlag returns Will Flag
func (c *ConnectMessage) WillFlag() byte {
	return (c.connectFlags >> 2) & 0x1
}

// SetCleanSession sets Clean Session
func (c *ConnectMessage) SetCleanSession(active bool) {
	if active {
		// 00000010
		c.connectFlags |= 0x02
	} else {
		// 11111101
		c.connectFlags &= 0xFC
	}
}

// CleanSession returns Clean Session value
func (c *ConnectMessage) CleanSession() byte {
	return (c.connectFlags >> 1) & 0x1
}

// SetKeepAlive sets Keep Alive
func (c *ConnectMessage) SetKeepAlive(v uint16) {
	c.keepAlive = v
}

// SetClientId sets ClientId and validates its correctness
func (c *ConnectMessage) SetClientId(cid []byte) error {
	// A Server MAY allow a Client to supply a ClientId that has a length of zero byte
	// however if it does so the Server MUST treat this as a special case and assign a unique
	// ClientId to the Client. It MUST then process the CONNECT packet as if the Client has
	// provided that unique ClientId [MQTT-3.1.3-6]
	//
	// If the Client supplies a zero-byte ClientId, the Client MUST also set CleanSession to 1
	// [MQTT-3.1.3-7]
	//
	// TODO: handle condition when cid length is zero
	if len(cid) == 0 || len(cid) > 23 {
		return ErrClientIdLengthInvalid
	}

	if !regexp.MustCompile("^[0-9a-zA-Z]*$").Match(cid) {
		return ErrClientIdInvalid
	}

	c.clientId = cid

	return nil
}

// SetWillTopic sets Will Topic and actives Will Flag
func (c *ConnectMessage) SetWillTopic(wt []byte) {
	// TODO: validate will topic format
	if len(wt) == 0 {
		c.SetWillFlag(false)
		return
	}

	c.SetWillFlag(true)
	c.willTopic = wt
}

// SetWillMessage sets Will Message and actives Will Flag
func (c *ConnectMessage) SetWillMessage(wm []byte) {
	// This field consists of a two byte length followed by the payload for the Will
	// Message expressed as a sequence of zero or more bytes. The length gives the
	// the number of bytes in the data that follows and does not include the 2 bytes
	// taken up by the length itself.
	// TODO: validate will message format
	if len(wm) == 0 {
		c.SetWillFlag(false)
		return
	}

	c.SetWillFlag(true)
	c.willMessage = wm
}

// SetUserName sets User Name and actives User Name Flag
func (c *ConnectMessage) SetUserName(un []byte) {
	if len(un) == 0 {
		c.SetUserNameFlag(false)
		return
	}

	c.SetUserNameFlag(true)
	c.userName = un
}

// SetPassword sets Password and actives Password Flag
func (c *ConnectMessage) SetPassword(pw []byte) {
	if len(pw) == 0 {
		c.SetPasswordFlag(false)
		return
	}

	c.SetPasswordFlag(true)
	c.password = pw
}

// Encode convert the struct to bytes
func (c *ConnectMessage) Encode(dest []byte) (int, error) {
	p := 0
	n, err := c.fixedHeader.Encode(dest)
	p += n
	if err != nil {
		return p, errors.New("failed to encode: " + err.Error())
	}

	n = copy(dest[p:], c.protocolName)
	p += n

	dest[p] = c.protocolLevel
	p++

	return p, nil
}
