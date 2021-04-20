package main

import (
	"fmt"
	"runtime"
	"sync"
)

var counter = 0
var wg sync.WaitGroup

func incrementor(i int) {
	fmt.Printf("New Goroutine #%v. Counter: %v\n", i, counter)
	_counter := counter
	runtime.Gosched()
	_counter++
	counter = _counter
	wg.Done()
}

func main() {
	wg.Add(6)
	go incrementor(1)
	go incrementor(2)
	go incrementor(3)
	go incrementor(4)
	go incrementor(5)
	go incrementor(6)
	wg.Wait()
	fmt.Println("Main: done.")
}
