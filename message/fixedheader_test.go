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
	testCases := []struct {
		in       uint32
		expected uint32
	}{
		{in: 0, expected: 0x00},
		{in: 127, expected: 0x7F},
		{in: 128, expected: 0x8001},
		{in: 16383, expected: 0xFF7F},
		{in: 16384, expected: 0x808001},
		{in: 2097151, expected: 0xFFFF7F},
		{in: 2097152, expected: 0x80808001},
		{in: 268435455, expected: 0xFFFFFF7F},
	}

	fh := fixedHeader{}
	for _, tc := range testCases {
		if el := fh.encodeLength(tc.in); el != tc.expected {
			t.Errorf("expected %x, got %x", tc.expected, el)
		}
	}
}

type mockReader struct {
	buf []byte
}

func (mc *mockReader) Read(b []byte) (int, error) {
	if len(mc.buf) > 1 {
		b[0], mc.buf = mc.buf[0], mc.buf[1:]
	} else if len(mc.buf) == 1 {
		b[0] = mc.buf[0]
	}
	return 1, nil
}

func TestFixedHeaderDecodeLength(t *testing.T) {

	testCases := []struct {
		r        *mockReader
		expected uint32
	}{
		{r: &mockReader{buf: []byte{0}}, expected: 0},
		{r: &mockReader{buf: []byte{0x7F}}, expected: 127},
		{r: &mockReader{buf: []byte{0x80, 0x01}}, expected: 128},
		{r: &mockReader{buf: []byte{0xFF, 0x7F}}, expected: 16383},
		{r: &mockReader{buf: []byte{0x80, 0x80, 0x01}}, expected: 16384},
		{r: &mockReader{buf: []byte{0xFF, 0xFF, 0x7F}}, expected: 2097151},
		{r: &mockReader{buf: []byte{0x80, 0x80, 0x80, 0x01}}, expected: 2097152},
		{r: &mockReader{buf: []byte{0xFF, 0xFF, 0xFF, 0x7F}}, expected: 268435455},
	}

	fh := fixedHeader{}
	for _, tc := range testCases {
		if l, _ := fh.decodeLength(tc.r); l != tc.expected {
			t.Errorf("expected %d, got %d", tc.expected, l)
		}
	}
}
