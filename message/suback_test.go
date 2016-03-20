package message

import (
	"encoding/binary"
	"reflect"
	"testing"
)

func TestSubackSetPacketID(t *testing.T) {
	s := &SubackMessage{}
	pid := make([]byte, 2)
	binary.BigEndian.PutUint16(pid, 12567)
	s.SetPacketID(pid)
	if !reflect.DeepEqual(pid, s.PacketID()) {
		t.Error("PacketID should be same as input")
	}
}

func TestSubackSetQoS(t *testing.T) {
	s := &SubackMessage{}

	q := byte(1)
	s.SetQoS(q)
	qos := s.QoS()
	if qos != q {
		t.Error("QoS should be 1")
	}
	q = byte(2)
	s.SetQoS(q)
	qos = s.QoS()
	if qos != q {
		t.Error("QoS should be 2")
	}
	q = byte(3)
	s.SetQoS(q)
	qos = s.QoS()
	if qos != q {
		t.Error("QoS should be 3")
	}

	err := s.SetQoS(4)
	if err == nil {
		t.Error("QoS sholud't add QoS larger than 3")
	}
	q = byte(128)
	s.SetQoS(q)
	qos = s.QoS()
	if qos != q {
		t.Error("QoS should be 128")
	}
}
