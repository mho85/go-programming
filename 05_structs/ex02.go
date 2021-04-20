package main

import "fmt"

type person2 struct {
	first   string
	last    string
	flavors []string
}

func main() {
	p1 := person2{
		first:   "Michel",
		last:    "Ho",
		flavors: []string{"Chocolate", "Passion fruit"},
	}
	p2 := person2{
		first:   "Châu",
		last:    "Nguyen",
		flavors: []string{"Pistacchio", "Vanilla"},
	}

	m := map[string]person2{
		p1.last: p1,
		p2.last: p2,
	}

	for _, p := range m {
		fmt.Println(p.first, p.last)
		for _, flavor := range p.flavors {
			fmt.Printf("\t%v\n", flavor)
		}
	}

}
