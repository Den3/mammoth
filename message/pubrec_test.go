package message

import (
	"encoding/binary"
	"reflect"
	"testing"
)

func TestPubrecSetPackID(t *testing.T) {
	p := &PubrecMessage{}
	pid := make([]byte, 2)
	binary.BigEndian.PutUint16(pid, 12567)
	p.SetPacketID(pid)
	if !reflect.DeepEqual(p.PacketID(), pid) {
		t.Error("Packet Identifier should be same as input")
	}
}
