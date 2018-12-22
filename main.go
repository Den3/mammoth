package main

import (
	"log"

	"github.com/Den3/mammoth/server"
)

func main() {
	s := &server.Server{}

	err := s.Listen()
	if err != nil {
		log.Fatal(err)
	}
}
