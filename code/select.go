// A select statement is like a switch, 
// but it selects over channel operations, 
// choosing exactly one of them.


// channels on their own are useful, 
// but made more powerful with select statements

// select statement will block until one of these channels are ready
package code

import ( 
	"fmt"
	"time"
)

func Select() {
	// ticker and boom are channels

	// returns a channel that ticks every .25 seconds
	ticker := time.NewTicker(time.Millisecond * 250) 
	// returns a channel that sends a single message after specified time period
	boom := time.After(time.Second * 1) 
	for {
		select {
		case <-ticker.C:
			fmt.Println("tick")
		case <-boom:
			fmt.Println("boom!")
			return
		}
	}
}
