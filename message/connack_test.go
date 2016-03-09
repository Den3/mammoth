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
