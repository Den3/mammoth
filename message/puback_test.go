package message

import (
	"reflect"
	"testing"
)

func TestPubackSetPackID(t *testing.T) {
	p := &PubackMessage{}
	pid := []byte("Hi MQTT")
	p.SetPacketID(pid)
	if !reflect.DeepEqual(p.PacketID(), pid) {
		t.Error("Packet Identifier should be same as input")
	}
}
