package message

import (
	"encoding/binary"
	"reflect"
	"testing"
)

func TestPublishMessageSetTopicName(t *testing.T) {
	p := &PublishMessage{}
	tn := []byte("a/b")
	p.SetTopicName(tn)
	if p.TopicNameLen() != 3 {
		t.Error("Topic Name Length shuold be 3")
	}
	if !reflect.DeepEqual(p.TopicName(), tn) {
		t.Error("Topic Name shuold be same as input")
	}

}

func TestPublishMessageSetPacketID(t *testing.T) {
	p := &PublishMessage{}
	pid := make([]byte, 2)
	binary.BigEndian.PutUint16(pid, 12567)
	p.SetPacketID(pid)
	if !reflect.DeepEqual(p.PacketID(), pid) {
		t.Error("PacketID shuold be same as input")
	}
}

func TestPublishMessageSetPayload(t *testing.T) {
	p := &PublishMessage{}
	pl := []byte("Hi MQTT")
	p.SetPayload(pl)
	if !reflect.DeepEqual(p.Payload(), pl) {
		t.Error("Payload shuold be same as input")
	}
}
