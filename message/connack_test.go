package message

import (
	"testing"
)

func TestConnackSetSessionPresent(t *testing.T) {
	c := &ConnackMessage{}
	c.SetSessionPresent(true)
	if c.SessionPresent() != 0x01 {
		t.Error("Connect Ack Flags should be set")
	}
}

func TestConnackSetConnectReturnCode(t *testing.T) {
	c := &ConnackMessage{}
	rc := byte(0x01)
	c.SetConnectReturnCode(rc)
	if c.ConnectReturnCode() != rc {
		t.Error("Connect Return Code should be same as input")
	}

	rc = byte(0x6)
	err := c.SetConnectReturnCode(rc)
	if err == nil {
		t.Error("Connect Return Code should be less than 6")
	}
}
