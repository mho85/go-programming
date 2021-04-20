package main

import "fmt"

func main() {
	vaccinatedRateThreshold := 0.9
	x := 0.3

	if x > vaccinatedRateThreshold {
		fmt.Println("OK. Immunity rate is reached")
	} else {
		fmt.Println("Not yet... Keep vaccinating")
	}
}
