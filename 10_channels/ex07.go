package main

import (
	"fmt"
)

func main() {

	c := make(chan int)

	// Launch 10 goroutines adding 10 values each
	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 10; j++ {
				c <- j
			}
			// We cannot close several times the same channel
		}()
	}

	// Receive
	for k := 0; k < 100; k++ { // therefore range is not possible
		fmt.Printf("(%v) %v\n", k, <-c)
	}

}
