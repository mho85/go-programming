package main

import "fmt"

func main() {
	favSports := "Football"
	switch favSports {
	case "Football":
		fmt.Println("PSG")
		fallthrough
	case "Basketball":
		fmt.Println("Lakers")
	default:
		fmt.Println("Other")
	}
}
