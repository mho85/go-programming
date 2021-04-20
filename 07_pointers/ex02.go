package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func main() {
	me := person{"Michel", "Ho", 36}
	fmt.Println("Before", me.first, me.last, me.age)
	changeMe(&me)
	fmt.Println("After", me.first, me.last, me.age)

}

func changeMe(p *person) {
	(*p).first = "Nâu"
	(*p).last = "Nâu"
	(*p).age = 27
}
