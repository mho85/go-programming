package main

import "fmt"

func main() {
	f := foo8()
	fmt.Println(f())
}

func foo8() func() int {
	return func() int {
		return 1985
	}
}
