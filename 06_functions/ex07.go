package main

import "fmt"

func main() {
	f := func() string {
		return "Hello"
	}
	fmt.Println(f())
}
