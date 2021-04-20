package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Main, nb goroutines:", runtime.NumGoroutine())
	go func() {
		fmt.Println("New. nb total goroutines:", runtime.NumGoroutine())
	}()
	go func() {
		fmt.Println("New. nb total goroutines:", runtime.NumGoroutine())
	}()
}
