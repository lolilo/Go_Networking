package code

import (
	"io"
	"log"
	"net"
	"fmt"
)

func TCP_chat() {
	l, err := net.Listen("tcp", listenAddr)
	if err != nil {
		log.Fatal(err)
	}
	for {
		c, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go match(c)
	}
}

var partner = make(chan io.ReadWriteCloser)

// decides which channel will hold two clients
func match(c io.ReadWriteCloser) {
	fmt.Fprint(c, "Waiting for a partner...")
	select {
	case partner <- c:
		//now handled by the other goroutine
	case p := <-partner:
		chat(p, c)
	}
}

// sends a greeting to each connection 
// copies data from one to the other, and vice versa
func chat(a, b io.ReadWriteCloser) {
	fmt.Fprintln(a, "Found one! Say hi.")
	fmt.Fprintln(b, "Found one! Say hi.")
	go io.Copy(a, b)
	io.Copy(b, a)
}
