package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func main() {
	me := person{
		first: "Michel",
		last:  "Ho",
		age:   36,
	}
	me.speak()
}

func (p person) speak() {
	fmt.Printf("I'm %v %v, %d years old.\n", p.last, p.first, p.age)
}
