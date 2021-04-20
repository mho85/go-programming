package main

import "fmt"

var x3 = 42
var y3 = "James Bond"
var z3 = true

func main() {
	s := fmt.Sprintln(x3, y3, z3)
	fmt.Println(s)
}
