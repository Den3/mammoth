package message

import (
	"encoding/binary"
	"reflect"
	"testing"
)

func TestUnsubackSetPacketID(t *testing.T) {
	s := &UnsubackMessage{}
	pid := make([]byte, 2)
	binary.BigEndian.PutUint16(pid, 12567)
	s.SetPacketID(pid)
	if !reflect.DeepEqual(pid, s.PacketID()) {
		t.Error("PacketID should be same as input")
	}
}
