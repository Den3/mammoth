package message

import "testing"

func TestSetControlPacketType(t *testing.T) {
	fh := &fixedHeader{}
	cpt := byte(0x0A)
	fh.SetControlPacketType(cpt)
	if fh.ControlPacketType() != cpt {
		t.Error("Control Packet Type should be same as input")
	}
}

func TestSetControlPacketTypeFlag(t *testing.T) {
	fh := &fixedHeader{}
	cptf := byte(0x0A)
	fh.SetControlPacketTypeFlag(cptf)
	if fh.ControlPacketTypeFlag() != cptf {
		t.Error("Control Packet Type Flag should be same as input")
	}
}

func TestSetRemainingLength(t *testing.T) {
	fh := &fixedHeader{}
	length := uint(12387)
	err := fh.SetRemainingLength(length)
	if err != nil {
		t.Error(err)
	}
}
