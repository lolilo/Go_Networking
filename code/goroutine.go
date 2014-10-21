package code

import (
	"fmt"
	"time"
)

func Goroutine() {
	go say("let's go!", 3)
	go say("ho!", 2)
	go say("hey!", 1)
	time.Sleep(4 * time.Second)
}

func say(text string, secs int) {
	time.Sleep(time.Duration(secs) * time.Second)
	fmt.Println(text)
}
