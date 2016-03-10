package message

import (
	"reflect"
	"testing"
)

func TestPubrecSetPackID(t *testing.T) {
	p := &PubrecMessage{}
	pid := []byte("Hi MQTT")
	p.SetPacketID(pid)
	if !reflect.DeepEqual(p.PacketID(), pid) {
		t.Error("Packet Identifier should be same as input")
	}
}
