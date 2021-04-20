package main

import "fmt"

func main() {
	for i := 65; i <= 90; i++ {
		fmt.Printf("%d\n", i)
		fmt.Printf("\t%#U\n", i)
	}
}
