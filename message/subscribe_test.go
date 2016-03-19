package message

import (
	"reflect"
	"testing"
)

func TestSubscribeaddTopic(t *testing.T) {
	s := &SubscribeMessage{}
	tn1 := []byte(`a\b`)
	tn2 := []byte(`c\d`)
	tn3 := []byte(`g\h`)
	s.addTopic(tn1)
	s.addTopic(tn2)
	s.addTopic(tn3)
	topics := s.Topics()
	count := 0
	for _, tp := range topics {
		if reflect.DeepEqual(tp, tn1) {
			count++
		}
		if reflect.DeepEqual(tp, tn2) {
			count++
		}
		if reflect.DeepEqual(tp, tn3) {
			count++
		}
	}
	if count < 3 {
		t.Error("Topic should be added to topics")
	}
}

func TestSubscribeaddQoS(t *testing.T) {
	s := &SubscribeMessage{}
	q1 := byte(1)
	q2 := byte(2)
	q3 := byte(3)
	s.addQoS(q1)
	s.addQoS(q2)
	s.addQoS(q3)

	qoSs := s.QoS()

	if !reflect.DeepEqual(qoSs[0], q1) {
		t.Error("QoS should be 1")
	}
	if !reflect.DeepEqual(qoSs[1], q2) {
		t.Error("QoS should be 2")
	}
	if !reflect.DeepEqual(qoSs[2], q3) {
		t.Error("QoS should be 3")
	}

	err := s.addQoS(4)
	if err == nil {
		t.Error("QoS sholud't add QoS larger than 3")
	}
}
