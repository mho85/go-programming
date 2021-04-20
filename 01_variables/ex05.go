package main

import "fmt"

type avenger5 int

var x5 avenger5 = 1
var y5 int

func main() {
	x5 = 42
	fmt.Printf("x5: %v (%T) \n", x5, x5)
	y5 = int(x5)
	fmt.Printf("y5: %v (%T) \n", y5, y5)
}
