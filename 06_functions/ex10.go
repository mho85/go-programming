package main

import "fmt"

func main() {

	{
		message := "I love you in this closure"
		fmt.Println(message)
	}

	// Unknown variable outside the closure
	// fmt.Println(message)
}
