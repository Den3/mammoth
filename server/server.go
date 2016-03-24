package server

import (
	"log"
	"net"
	"time"
)

const (
	// Port is that MQTT listens on port 1883 for TCP
	Port = "1883"

	// AcceptInterval is connection acception interval in micro second
	AcceptInterval = 10
)

// Server is listening on port 1883 only
type Server struct {
}

// getControlPacketType gets Contorl Packet type
func (s *Server) getControlPacketType(t byte) string {
	t = t >> 4
	switch t {
	case 1:
		return "CONNECT"
	case 2:
		return "CONNACK"
	case 3:
		return "PUBLISH"
	case 4:
		return "PUBACK"
	case 5:
		return "PUBREC"
	case 6:
		return "PUBREL"
	case 7:
		return "PUBCOMP"
	case 8:
		return "SUBSCRIBE"
	case 9:
		return "SUBACK"
	case 10:
		return "UNSUBSCRIBE"
	case 11:
		return "UNSUBACK"
	case 12:
		return "PINGREQ"
	case 13:
		return "PINGRESP"
	case 14:
		return "DISCONNECT"
	}
	return ""
}

// handleConn judges its MQTT type
func (s *Server) handleConn(c net.Conn) {
	switch c.(type) {
	case *net.TCPConn:
		log.Println("TCP connection")
	}

}

// Listen Listen on port 1883 only
func (s *Server) Listen() error {
	log.Println("Server starting...")
	ln, err := net.Listen("tcp", "0.0.0.0:"+Port)
	if err != nil {
		return err
	}
	defer ln.Close()

	for {
		c, err := ln.Accept()
		if err != nil {
			log.Println("accept conn error:", err)
			continue
		}
		go s.handleConn(c)
		time.Sleep(AcceptInterval * time.Microsecond)
	}
}
