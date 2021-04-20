package main

import "fmt"

func main() {
	defer fmt.Println("Awesome!")
	fmt.Println("I am...")
	fmt.Println("Wait for it...")
}
