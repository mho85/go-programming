package main

import "fmt"

func main() {
	s1 := foo2(2, 3)
	sx := []int{4, 7}
	s2 := bar2(sx)
	fmt.Println("s1:", s1)
	fmt.Println("s2:", s2)
}

func foo2(x ...int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}

func bar2(sx []int) int {
	sum := 0
	for _, v := range sx {
		sum += v
	}
	return sum
}
