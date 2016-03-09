package message

import (
	"testing"
)

func TestSetSessionPresent(t *testing.T) {
	c := &ConnackMessage{}
	c.SetSessionPresent(true)
	if c.SessionPresent() != 0x01 {
		t.Error("Connect Ack Flags should be set")
	}
}

func TestSetConnectReturnCode(t *testing.T) {
	c := &ConnackMessage{}
	rc := byte(0x01)
	c.SetConnectReturnCode(rc)
	if c.ConnectReturnCode() != rc {
		t.Error("Connect Return Code should be same as input")
	}
}
