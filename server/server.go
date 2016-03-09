package server

import (
	"log"
	"net"
)

const (
	PORT = "8080"
)

func handle(c *net.TCPConn) {
	log.Println("one connection is coming...")
}

func Serve() {
	rAddr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:"+PORT)
	if err != nil {
		panic(err)
	}

	listener, err := net.ListenTCP("tcp", rAddr)
	if err != nil {
		panic(err)
	}
	log.Println("start listen on", PORT, "...")
	defer listener.Close()

	for {
		c, err := listener.AcceptTCP()
		if err != nil {
			log.Println(err)
			continue
		}
		go handle(c)
	}
}
