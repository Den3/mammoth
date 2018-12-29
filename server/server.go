package server

import (
	"log"
	"net"

	"github.com/Den3/mammoth/message"
)

const (
	// Port is that MQTT listens on port 1883 for TCP
	Port = "1883"

	// AcceptInterval is connection acception interval in micro second
	AcceptInterval = 10
)

// Server is listening on port 1883 only
type Server struct{}

// handleConn judges its MQTT type
func (s *Server) handleConn(c net.Conn) {
	m := message.NewConnackMessage()
	buf := make([]byte, 3)
	m.Encode(buf)
	_, err := c.Write(buf)
	if err != nil {
		log.Fatal(err)
	}
}

// Listen Listen on port 1883 only
func (s *Server) Listen() error {
	log.Println("Server starting...")
	ln, err := net.Listen("tcp", "0.0.0.0:"+Port)
	log.Println("Server lisening on " + Port + "...")
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
		s.handleConn(c)
	}
}
