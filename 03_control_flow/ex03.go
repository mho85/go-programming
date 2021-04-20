package main

import "fmt"

func main() {
	yearOfBirth := 1985
	currentYear := 2021
	for yearOfBirth <= currentYear {
		fmt.Println(yearOfBirth)
		yearOfBirth++
	}
}
