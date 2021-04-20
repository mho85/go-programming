package main

import "fmt"

func main() {
	x := func(a ...int) int {
		sum := 0
		for _, v := range a {
			sum += v
		}
		return sum
	}(2, 3, 5)
	fmt.Println(x)
}
