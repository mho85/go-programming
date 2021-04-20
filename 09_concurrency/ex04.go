package main

import (
	"fmt"
	"sync"
)

var counter4 = 0
var wg4 sync.WaitGroup
var mu sync.Mutex

func incrementor4(i int) {
	mu.Lock()
	_counter := counter4
	_counter++
	counter4 = _counter
	fmt.Printf("Goroutine #%v (end). Counter: %v\n", i, counter4)
	mu.Unlock()
	wg4.Done()
}

func main() {
	wg4.Add(6)
	go incrementor4(1)
	go incrementor4(2)
	go incrementor4(3)
	go incrementor4(4)
	go incrementor4(5)
	go incrementor4(6)
	wg4.Wait()
	fmt.Printf("Main: done. Counter: %v\n", counter4)
}
