package main

import "fmt"

type person struct {
	first   string
	last    string
	flavors []string
}

func main() {
	p1 := person{
		first:   "Michel",
		last:    "Ho",
		flavors: []string{"Chocolate", "Passion fruit"},
	}
	p2 := person{
		first:   "Châu",
		last:    "Nguyen",
		flavors: []string{"Pistacchio", "Vanilla"},
	}
	people := []person{p1, p2}

	for _, p := range people {
		fmt.Println(p.first, p.last)
		for _, flavor := range p.flavors {
			fmt.Printf("\t%v\n", flavor)
		}
	}

}
