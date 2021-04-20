package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var counter5 int64 = 0
var wg5 sync.WaitGroup

func incrementor5(i int) {
	_counter := atomic.LoadInt64(&counter5)
	_counter++
	atomic.StoreInt64(&counter5, _counter)
	fmt.Printf("Goroutine #%v (end). Counter: %v\n", i, atomic.LoadInt64(&counter5))
	wg5.Done()
}

func main() {
	wg5.Add(6)
	go incrementor5(1)
	go incrementor5(2)
	go incrementor5(3)
	go incrementor5(4)
	go incrementor5(5)
	go incrementor5(6)
	wg5.Wait()
	fmt.Printf("Main: done. Counter: %v\n", counter5)
}
