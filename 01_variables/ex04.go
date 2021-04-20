package main

import "fmt"

type avenger int

var x4 avenger = 1

func main() {
	fmt.Printf("x4: %v (%T)\n", x4, x4)
	x4 = 42
	fmt.Printf("x4: %v (%T)\n", x4, x4)
}
