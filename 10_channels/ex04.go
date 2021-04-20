package main

import (
	"fmt"
)

func main() {
	q := make(chan int)
	c := gen4(q)

	receive4(c, q)

	fmt.Println("about to exit")
}

func gen4(q chan<- int) <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		q <- 1
		close(c)
	}()

	return c
}

func receive4(c, q <-chan int) {
	for {
		select {
		case v := <-c:
			fmt.Println(v)
		case <-q:
			return
		}
	}
}
