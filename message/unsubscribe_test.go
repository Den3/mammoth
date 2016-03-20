package message

import (
	"encoding/binary"
	"reflect"
	"testing"
)

func TestUnsubscribeSetPacketID(t *testing.T) {
	s := &UnsubscribeMessage{}
	pid := make([]byte, 2)
	binary.BigEndian.PutUint16(pid, 12567)
	s.SetPacketID(pid)
	if !reflect.DeepEqual(pid, s.PacketID()) {
		t.Error("PacketID should be same as input")
	}
}

func TestSubscribeaAddTopic(t *testing.T) {
	s := &UnsubscribeMessage{}
	tn1 := []byte(`a\b`)
	tn2 := []byte(`c\d`)
	tn3 := []byte(`g\h`)
	s.AddTopic(tn1)
	s.AddTopic(tn2)
	s.AddTopic(tn3)
	topics := s.Topics()
	count := 0
	for _, tp := range topics {
		if reflect.DeepEqual(tp, tn1) {
			count++
		}
		if reflect.DeepEqual(tp, tn2) {
			count++
		}
		if reflect.DeepEqual(tp, tn3) {
			count++
		}
	}
	if count < 3 {
		t.Error("Topic should be added to topics")
	}
}
