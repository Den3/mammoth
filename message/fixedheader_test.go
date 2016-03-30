package message

import "testing"

func TestFixedHeaderSetControlPacketType(t *testing.T) {
	fh := &fixedHeader{}
	cpt := byte(0x0A)
	fh.SetControlPacketType(cpt)
	if fh.ControlPacketType() != cpt {
		t.Error("Control Packet Type should be same as input")
	}
}

func TestFixedHeaderSetControlPacketTypeFlag(t *testing.T) {
	fh := &fixedHeader{}
	cptf := byte(0x0A)
	fh.SetControlPacketTypeFlag(cptf)
	if fh.ControlPacketTypeFlag() != cptf {
		t.Error("Control Packet Type Flag should be same as input")
	}
}

func TestFixedHeaderEncodeLength(t *testing.T) {
	fh := fixedHeader{}
	v := 127
	ev := fh.encodeLength(v)
	if ev[0] != 0x7F {
		t.Error("Encode value of 127 should be 0x7F")
	}

	v = 128
	ev = fh.encodeLength(v)
	if ev[0] != 0x80 || ev[1] != 0x01 {
		t.Error("Encode value of 128 should be 0x80 0x01")
	}

	v = 16383
	ev = fh.encodeLength(v)
	if ev[0] != 0xFF || ev[1] != 0x7F {
		t.Error("Encode value of 16383 should be 0xFF 0x7F")
	}

	v = 16384
	ev = fh.encodeLength(v)
	if ev[0] != 0x80 || ev[1] != 0x80 || ev[2] != 0x01 {
		t.Error("Encode value of 16383 should be 0x80 0x80 0x01")
	}

	v = 2097151
	ev = fh.encodeLength(v)
	if ev[0] != 0xFF || ev[1] != 0xFF || ev[2] != 0x7F {
		t.Error("Encode value of 16383 should be 0xFF 0xFF 0x7F")
	}

	v = 2097152
	ev = fh.encodeLength(v)
	if ev[0] != 0x80 || ev[1] != 0x80 || ev[2] != 0x80 || ev[3] != 0x01 {
		t.Error("Encode value of 16383 should be 0x80 0x80 0x80 0x01")
	}

	v = 268435455
	ev = fh.encodeLength(v)
	if ev[0] != 0xFF || ev[1] != 0xFF || ev[2] != 0xFF || ev[3] != 0x7F {
		t.Error("Encode value of 16383 should be 0xFF 0xFF 0xFF 0x7F")
	}

}

type mockConn struct {
	buf []byte
}

func (mc *mockConn) Read(b []byte) (int, error) {
	if len(mc.buf) > 1 {
		b[0], mc.buf = mc.buf[0], mc.buf[1:]
	} else if len(mc.buf) == 1 {
		b[0] = mc.buf[0]
	}
	return 1, nil
}

func TestFixedHeaderDecodeLength(t *testing.T) {
	fh := &fixedHeader{}
	mc := &mockConn{}

	mc.buf = make([]byte, 1)
	mc.buf[0] = 0x7F
	v, err := fh.decodeLength(mc)
	if err != nil {
		panic(err)
	}
	if v != 127 {
		t.Error("Decode value of 0x7F should be 127")
	}

	mc.buf = make([]byte, 2)
	mc.buf[0] = 0x80
	mc.buf[1] = 0x01
	v, err = fh.decodeLength(mc)
	if err != nil {
		panic(err)
	}
	if v != 128 {
		t.Error("Decode value of 0x80 0x01 should be 128")
	}

	mc.buf = make([]byte, 2)
	mc.buf[0] = 0xFF
	mc.buf[1] = 0x7F
	v, err = fh.decodeLength(mc)
	if err != nil {
		panic(err)
	}
	if v != 16383 {
		t.Error("Decode value of 0xFF 0x7F should be 16383")
	}

	mc.buf = make([]byte, 3)
	mc.buf[0] = 0x80
	mc.buf[1] = 0x80
	mc.buf[2] = 0x01
	v, err = fh.decodeLength(mc)
	if err != nil {
		panic(err)
	}
	if v != 16384 {
		t.Error("Decode value of 0x80 0x80 0x01 should be 16384")
	}

	mc.buf = make([]byte, 3)
	mc.buf[0] = 0xFF
	mc.buf[1] = 0xFF
	mc.buf[2] = 0x7F
	v, err = fh.decodeLength(mc)
	if err != nil {
		panic(err)
	}
	if v != 2097151 {
		t.Error("Decode value of 0xFF 0xFF 0x7F should be 2097151")
	}

	mc.buf = make([]byte, 4)
	mc.buf[0] = 0x80
	mc.buf[1] = 0x80
	mc.buf[2] = 0x80
	mc.buf[3] = 0x01
	v, err = fh.decodeLength(mc)
	if err != nil {
		panic(err)
	}
	if v != 2097152 {
		t.Error("Decode value of 0x80 0x80 0x80 0x01 should be 2097152")
	}

	mc.buf = make([]byte, 4)
	mc.buf[0] = 0xFF
	mc.buf[1] = 0xFF
	mc.buf[2] = 0xFF
	mc.buf[3] = 0x7F
	v, err = fh.decodeLength(mc)
	if err != nil {
		panic(err)
	}
	if v != 268435455 {
		t.Error("Decode value of 0xFF 0xFF 0xFF 0x7F should be 268435455")
	}
}
