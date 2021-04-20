package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type person2 struct {
	First   string
	Last    string
	Sayings []string
}

func main() {
	p1 := person2{
		First:   "James",
		Last:    "Bond",
		Sayings: []string{"Shaken, not stirred", "Any last wishes?", "Never say never"},
	}

	bs, err := toJSON(p1)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(bs))

}

// toJSON needs to return an error also
func toJSON(a interface{}) ([]byte, error) {
	bs, err := json.Marshal(a)
	return bs, fmt.Errorf("Error in toJSON: %v\n", err)
}
