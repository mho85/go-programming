package main

import "fmt"

func main() {
	yearOfBirth := 1985
	currentYear := 2021
	year := yearOfBirth

	for {
		if year > currentYear {
			break
		}
		fmt.Println(year)
		year++
	}
}
