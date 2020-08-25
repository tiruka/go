package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
Outerloop:
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("Boom!")
			break Outerloop
		default:
			fmt.Println(".")
			time.Sleep(100 * time.Millisecond)
		}
	}
	fmt.Println("after")
}
