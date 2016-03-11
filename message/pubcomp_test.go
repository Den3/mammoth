package message

import (
	"reflect"
	"testing"
)

func TestPubcompSetPackID(t *testing.T) {
	p := &PubcompMessage{}
	pid := []byte("Hi MQTT")
	p.SetPacketID(pid)
	if !reflect.DeepEqual(p.PacketID(), pid) {
		t.Error("Packet Identifier should be same as input")
	}
}
