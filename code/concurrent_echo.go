package code

import (
	"io"
	"log"
	"net"
)

// const listenAddr = "localhost:4000"

func Concurrent_echo() {
	l, err := net.Listen("tcp", listenAddr)
	if err != nil {
		log.Fatal(err)
	}
	for {
		c, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go io.Copy(c, c)
	}
}

