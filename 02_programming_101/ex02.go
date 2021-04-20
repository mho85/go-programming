package main

import "fmt"

func main() {
	x := 2
	y := 3
	exp1 := x == y
	exp2 := x <= y
	exp3 := x >= y
	exp4 := x != y
	exp5 := x < y
	exp6 := x > y

	fmt.Println(exp1, exp2, exp3, exp4, exp5, exp6)
}
