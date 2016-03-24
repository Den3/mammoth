package server

import "testing"

func TestServerGetControlPacketType(t *testing.T) {
	s := Server{}

	if s.getControlPacketType(1<<4) != "CONNECT" {
		t.Error("result should be CONNECT")
	}
	if s.getControlPacketType(2<<4) != "CONNACK" {
		t.Error("result should be CONNACK")
	}
	if s.getControlPacketType(3<<4) != "PUBLISH" {
		t.Error("result should be PUBLISH")
	}
	if s.getControlPacketType(4<<4) != "PUBACK" {
		t.Error("result should be PUBACK")
	}
	if s.getControlPacketType(5<<4) != "PUBREC" {
		t.Error("result should be PUBREC")
	}
	if s.getControlPacketType(6<<4) != "PUBREL" {
		t.Error("result should be PUBREL")
	}
	if s.getControlPacketType(7<<4) != "PUBCOMP" {
		t.Error("result should be PUBCOMP")
	}
	if s.getControlPacketType(8<<4) != "SUBSCRIBE" {
		t.Error("result should be SUBSCRIBE")
	}
	if s.getControlPacketType(9<<4) != "SUBACK" {
		t.Error("result should be SUBACK")
	}
	if s.getControlPacketType(10<<4) != "UNSUBSCRIBE" {
		t.Error("result should be UNSUBSCRIBE")
	}
	if s.getControlPacketType(11<<4) != "UNSUBACK" {
		t.Error("result should be UNSUBACK")
	}
	if s.getControlPacketType(12<<4) != "PINGREQ" {
		t.Error("result should be PINGREQ")
	}
	if s.getControlPacketType(13<<4) != "PINGRESP" {
		t.Error("result should be PINGRESP")
	}
	if s.getControlPacketType(14<<4) != "DISCONNECT" {
		t.Error("result should be DISCONNECT")
	}

}
