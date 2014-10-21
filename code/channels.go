// Goroutines communicate via channels.
// A channel is a typed conduit, like a Unix pipe.
// Channels may be synchronous (unbuffered) or asynchronous (buffered).

// Channels are the pipes that connect concurrent goroutines. 
// You can send values into channels from one goroutine and 
// receive those values into another goroutine.

// channel syntax <-
// data flows in the direction of the arrow

package code

import "fmt"

func Channels() {
	ch := make(chan int) // channels are typed by the values they carry
	go fibs(ch)
	for i := 0; i < 20; i++ {
		fmt.Println(<-ch) // receive from the channel and Println
	}
}

// loops forever, sending fibonacci numbers down the given channel
func fibs(ch chan int) {
	i, j := 0, 1
	for {
		ch <- j // blocks until there is a receiver on the other side
		// isn't spinning forever -- 
		// waiting and producing 
		// kind of like a generator in this way
		i, j = j, i+j
	}
}

