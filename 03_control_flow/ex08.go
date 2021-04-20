package main

import "fmt"

func main() {
	switch {
	case 1 == 1:
		fmt.Println("One")
		fallthrough
	case 2 == 3:
		fmt.Println("Two")
	default:
		fmt.Println("Three")
	}
}
