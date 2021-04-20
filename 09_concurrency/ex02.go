package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

/*// Pointer receiver
func (p *person) speak() {
	fmt.Println("I speak")
}*/

// Value receiver
func (p person) speak() {
	fmt.Println("I speak")
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	p1 := person{
		first: "Michel",
		last:  "Ho",
		age:   36,
	}

	/*// Pointer-receiver
	saySomething(&p1) // OK
	// saySomething(p1) // NOK*/

	// Value-receiver
	saySomething(&p1) // OK
	saySomething(p1)  // OK
}
