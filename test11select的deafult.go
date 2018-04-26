package main

import (
	"time"
	"fmt"
)

func main() {
	tick := time.Tick(10000 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			//break
			return
		//default:
		//	fmt.Print(".")
		//	time.Sleep(5000* time.Millisecond)
		}
	}
}