package message

import "encoding/binary"

// PublishMessage is a PUBLISH Control Packet is sent from a Client to a Server or from
// Server to a Client to transport an Application Message
type PublishMessage struct {
	// 3.3.1.1 DUP
	// Postion: byte 1, bit3.
	// If the DUP flag is set to 0, it indicates that this is the occasion that the
	// Client or Server has attempted to send this MQTT PUBLISH Packet. If the DUP
	// flag is set tot 1, it indicates that this might be re-delivery of an earlier
	// attempt to send the Packet
	//
	// The DUP flag MUST be set to 1 by the Client or Server when it attempts to
	// re-deliver a PUBLISH Packet[MQTT-3.3.1-1]. The DUP flag MUST be set to 0
	// for all QoS messages[MQTT-3.3.1-2].
	//
	// The value of the DUP flag from an incoming packet is not propagated when
	// the PUBLISH Packet is sent to subscribers by the Server. The DUP flag
	// in the outgoing PUBLISH packet is set independently to the incoming
	// PUBLISH packet, its value MUST be determined solely by whether the
	// outgoing PUBLISH packet is a retransmission[MQTT-3.3.1-3].
	//
	// 3.3.1.2 QoS
	// Postion: byte 1, bits 2-1.
	// This field indicates the level of assurance for delivery of an Application
	// Message. The QoS levels are listed in the Table 3.2 - QoS definitions, below
	// -----------------------------------------------------------------------------------
	// | QoS value |  Bit 2  |  bit 1  |                Description                      |
	// -----------------------------------------------------------------------------------
	// |     0     |    0    |    0    | At most once delivery                           |
	// -----------------------------------------------------------------------------------
	// |     1     |    0    |    1    | At least once delivery                          |
	// -----------------------------------------------------------------------------------
	// |     2     |    1    |    0    | Exactly once delivery                           |
	// -----------------------------------------------------------------------------------
	// |     -     |    1    |    1    | Reserved - must not be used                     |
	// -----------------------------------------------------------------------------------
	// A PUBLISH Packet MUST NOT have both QoS bits set to 1. If a Server or Client
	// receives a PUBLISH Packet which both QoS bits set to 1 it MUST close the Network
	// Connection[MQTT-3.3.1-4].
	//
	// 3.3.1.3 RETAIN
	// Position: byte 1, bit 0.
	// If the RETAIN flag is set to 1, in a PUBLISH Packet sent by a Client to a Server,
	// the Server MUST store the Application Message and its QoS, so that it can be
	// delivered to future subscribers whose subscriptions match its topic name
	// [MQTT-3.3.1-5]. When a new subscription is established, the last retained message,
	// if any, on each matching topic name MUST be sent to the subscriber[MQTT-3.3.1-6].
	// If the Server receives a QoS 0 message with the RETAIN flag set to 1 it MUST
	// discard any message previously retained for that topic. It SHOULD store the new
	// QoS message as the new retained message for that topic, but MAY choose to discard
	// it at any time - if this happens there will be no retained message for that topic
	// [MQTT-3.3.1-7]. See Section 4.1 for more infromation on storing state.
	//
	// When sending a PUBLISH Packet to a Client the Server MUST set the RETAIN flag to
	// 1 if a message is sent as a result of a new subscription being made by a Client
	// [MQTT-3.3.1-8]. It MUST set the RETAIN flag to 0 when a PUBLISH Packet is sent
	// to a Client because it matches an established subscription regardless of how the
	// flag was set in the message it received[MQTT-3.3.1-9].
	//
	// A PUBLISH Packet with a RETAIN flag set to 1 and payload containing zero bytes
	// will be processed as normal by the Server and sent to Clients with a subscription
	// matching the topic name. Additionally any existing retained message with the same
	// topic name MUST be removed and any future subscribers for the topic will not
	// receive a retained message[MQTT-3.3.1-10]. "As normal" means that the RETAIN flag
	// is not set in the message received by existing Clients. A zero byte retained
	// message MUST NOT be stored as a retained message on the Server[MQTT-3.3.1-11].
	//
	// If the RETAIN flag is 0, in a PUBLISH Packet sent by a Client to a Server, the
	// Server MUST NOT store the message and MUST NOT remove or replace any existing
	// retained message[MQTT-3.3.1-12]
	fixedHeader

	// The Topic Name identifies the information channel to which payload data is
	// published.
	//
	// The Topic Name MUST be present as the first field in the PUBLISH Packet
	// Variable header. It MUST be a UTF-8 encoded string[MQTT-3.3.2-1] as defined
	// in section 1.5.3
	// The Topic Name in the PUBLISH Packet sent by a Server to a subscribing
	// Client MUST match the Subscription's Topic Filter according to the matching
	// process defined in Section 4.7[MQTT-3.3.2-3]. However, since the Server
	// is permitted to override the Topic Name, it might not be the same as the
	// Topic Name in the original PUBLISH Packet.
	topicName []byte

	// The Packet identifier field is only present in PUBLISH Packets where the
	// QoS level is 1 or 2. Section 2.3.1 provides more infromation about Packet
	// Identifiers.
	packetID []byte

	// The Payload contains the Application Message that is bening published. The
	// content and format of the data is application specific. The length of the
	// payload can be calculated by subtracting the length of the variable header
	// from the Remaining Length field that is in the Fixed Header. It is valid for
	// a PUBLISH Packet to contain a zero length payload.
	payload []byte
}

// SetTopicName sets Topic Name and its length
func (p *PublishMessage) SetTopicName(v []byte) {
	length := make([]byte, 2)
	binary.BigEndian.PutUint16(length, uint16(len(v)))
	p.topicName = append(p.topicName, length...)
	p.topicName = append(p.topicName, v...)
}

// TopicNameLen returns Topic Name first two bytes representing length
func (p *PublishMessage) TopicNameLen() uint16 {
	return binary.BigEndian.Uint16(p.topicName[:2])
}

// TopicName returns real Topic Name content
func (p *PublishMessage) TopicName() []byte {
	return p.topicName[2:]
}

// SetPacketID sets Packet Identifiers
func (p *PublishMessage) SetPacketID(v []byte) {
	p.packetID = v
}

// PacketID returns Packet Identifiers
func (p *PublishMessage) PacketID() []byte {
	return p.packetID
}

// SetPayload sets payload
func (p *PublishMessage) SetPayload(v []byte) {
	p.payload = v
}

// Payload returns payload
func (p *PublishMessage) Payload() []byte {
	return p.payload
}
