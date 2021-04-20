package main

import "fmt"

func main() {
	arr := [5]int{10, 11, 12, 13, 14}
	for _, v := range arr {
		fmt.Printf("%v (%T)\n", v, v)
	}
}
