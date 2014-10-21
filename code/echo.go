package main

import (
	"io"
	"log"
	"net"
)

const linstenAddr = "localhost:4000"


// not concurrent. can only have one connection. 
func echo() {
	l, err := net.Listen("tcp", listenAddr)
	if err != nil {
		log.Fatal(err)
	}
	for {
		c, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		io.Copy(c, c) // copies from a reader to a writer 
	}
}



