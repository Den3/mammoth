package message

import (
	"reflect"
	"testing"
)

func TestSetProtocolName(t *testing.T) {
	c := &ConnectMessage{}

	c.SetProtocolName()
	if c.protocolName[1] != 4 ||
		c.protocolName[2] != []byte("M")[0] ||
		c.protocolName[3] != []byte("Q")[0] ||
		c.protocolName[4] != []byte("T")[0] ||
		c.protocolName[5] != []byte("T")[0] {
		t.Error("Protocol Name should be MQTT")
	}
}

func TestSetProtocolLevel(t *testing.T) {
	c := &ConnectMessage{}

	c.SetProtocolLevel()
	if c.protocolLevel != 4 {
		t.Error("Protocol Level should be 4")
	}
}

func TestSetUserNameFlag(t *testing.T) {
	c := &ConnectMessage{}

	c.SetUserNameFlag(true)
	if c.UserNameFlag() != 0x1 {
		t.Error("User Name Flag should be true")
	}

	c.SetUserNameFlag(false)
	if c.UserNameFlag() != 0x0 {
		t.Error("User Name Flag should be false")
	}
}

func TestSetPasswordFlag(t *testing.T) {
	c := &ConnectMessage{}

	c.SetPasswordFlag(true)
	if c.PasswordFlag() != 0x1 {
		t.Error("Password Flag should be true")
	}

	c.SetPasswordFlag(false)
	if c.PasswordFlag() != 0x0 {
		t.Error("Password Flag should be false")
	}
}

func TestSetWillRetain(t *testing.T) {
	c := &ConnectMessage{}

	c.SetWillRetain(true)
	if c.WillRetain() != 0x1 {
		t.Error("Retain Flag should be true")
	}

	c.SetWillRetain(false)
	if c.WillRetain() != 0x0 {
		t.Error("Retain Flag should be false")
	}
}

func TestSetWillQoS(t *testing.T) {
	c := &ConnectMessage{}

	c.SetWillQoS(0)
	if c.WillQoS() != 0x0 {
		t.Error("QoS Flag should be 0")
	}

	c.SetWillQoS(1)
	if c.WillQoS() != 0x1 {
		t.Error("QoS Flag should be 1")
	}

	c.SetWillQoS(2)
	if c.WillQoS() != 0x2 {
		t.Error("QoS Flag should be 2")
	}

	c.SetWillQoS(3)
	if c.WillQoS() != 0x3 {
		t.Error("QoS Flag should be 3")
	}

	err := c.SetWillQoS(4)
	if err == nil {
		t.Error("setQosFlag error: 4 is invalid")
	}
}

func TestSetWillFlag(t *testing.T) {
	c := &ConnectMessage{}

	c.SetWillFlag(true)
	if c.WillFlag() != 0x1 {
		t.Error("Will Flag should be true")
	}

	c.SetWillFlag(false)
	if c.WillFlag() != 0x0 {
		t.Error("Will Flag should be false")
	}
}

func TestSetCleanSession(t *testing.T) {
	c := &ConnectMessage{}

	c.SetCleanSession(true)
	if c.CleanSession() != 0x1 {
		t.Error("Clean Session should be true")
	}

	c.SetCleanSession(false)
	if c.CleanSession() != 0x0 {
		t.Error("Clean Session should be false")
	}
}

func TestSetClientId(t *testing.T) {
	c := &ConnectMessage{}

	cid := []byte("")
	err := c.SetClientId(cid)
	if err == nil {
		t.Error("ClientId length should be biger than 1")
	}

	cid = []byte("123456789012345678901234567890")
	err = c.SetClientId(cid)
	if err == nil {
		t.Error("ClientId length should be less than or equal 23")
	}

	cid = []byte("1234567890abcDefGhijklm")
	err = c.SetClientId(cid)
	if err != nil {
		t.Error(err)
	}

	cid = []byte("1234567890abcDefGhijklm_(")
	err = c.SetClientId(cid)
	if err == nil {
		t.Error("ClientId should be [0-9a-zA-Z]")
	}
}

func TestSetWillTopic(t *testing.T) {
	c := &ConnectMessage{}
	willTopic := []byte("a/b")
	c.SetWillTopic(willTopic)
	if !reflect.DeepEqual(willTopic, c.willTopic) {
		t.Error("Will Topic should be same as input")
	}

	if c.WillFlag() != 0x1 {
		t.Error("Will Flag should be set")
	}
}

func TestSetWillMessage(t *testing.T) {
	c := &ConnectMessage{}
	willMessage := []byte("Hi MQTT")
	c.SetWillMessage(willMessage)
	if !reflect.DeepEqual(willMessage, c.willMessage) {
		t.Error("Will Message is not same as input")
	}

	if c.WillFlag() != 0x1 {
		t.Error("Will Flag should be set")
	}
}

func TestSetUserName(t *testing.T) {
	c := &ConnectMessage{}
	userName := []byte("mqtt")
	c.SetUserName(userName)
	if !reflect.DeepEqual(userName, c.userName) {
		t.Error("User Name should be same as input")
	}

	if c.UserNameFlag() != 0x1 {
		t.Error("User Name Flag should be set")
	}
}

func TestSetPassword(t *testing.T) {
	c := &ConnectMessage{}
	password := []byte("mqtt")
	c.SetPassword(password)
	if !reflect.DeepEqual(password, c.password) {
		t.Error("Password should be same as input")
	}

	if c.PasswordFlag() != 0x1 {
		t.Error("Password Flag should be set")
	}
}
