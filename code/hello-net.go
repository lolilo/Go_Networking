// listens on a TCP port
// when a socket comes in, say "Hello" 
// then shut the socket down

package main

import (
	"fmt"
	"log"
	"net"
)

const listenAddr = "localhost:4000"

// func main() {
	
func hello_net() {
	l, err := net.Listen("tcp", listenAddr) // returns "l", a net listener
	if err != nil {
		log.Fatal(err)
	}

	// loop forever in this for-loop (Go has no while-loops)
	for {
		c, err := l.Accept() // accept method on listener, which will block until a connection comes in
		
		// bail if an error occurs
		if err != nil {
			log.Fatal(err)
		}
		
		fmt.Fprintln(c, "Hello!")
		c.Close()
	}
}

