package message

import (
	"reflect"
	"testing"
)

func TestPubrelSetPackID(t *testing.T) {
	p := &PubrelMessage{}
	pid := []byte("Hi MQTT")
	p.SetPacketID(pid)
	if !reflect.DeepEqual(p.PacketID(), pid) {
		t.Error("Packet Identifier should be same as input")
	}
}
