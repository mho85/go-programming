package main

import "fmt"

func main() {

	c := make(chan int)

	// Send to channel
	populate(c)
	fmt.Println("Finished populating channel")

	// Receive from channel
	pull(c)
	fmt.Println("End of main program")

}

func populate(c chan<- int) {
	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()
}

func pull(c <-chan int) {
	for v := range c {
		fmt.Println(v)
	}
}
