// Package dog is all about dogs
package main

import (
	"fmt"
	"github.com/mho85/go-programming/12_documentation/dog"
)

func main() {
	theDog := dog.Dog{7}
	fmt.Printf("Rantanplan is %v years old or %v human years old\n",
		theDog.Age,
		dog.Years(theDog.Age))
}
