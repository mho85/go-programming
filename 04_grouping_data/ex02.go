package main

import "fmt"

func main() {
	s := []int{10, 11, 12, 13, 14, 15, 16, 17, 18, 19}
	for _, v := range s {
		fmt.Printf("%v (%T)\n", v, v)
	}
}
