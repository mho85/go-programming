package main

import "fmt"

func main() {
	vaccinatedRateThreshold := 0.9
	x := 0.3

	if x > vaccinatedRateThreshold {
		fmt.Println("OK. Immunity rate is reached")
	} else if x > 0.5 {
		fmt.Println("Almost...")
	} else {
		fmt.Println("Not yet... Keep vaccinating")
	}
}
