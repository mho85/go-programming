package main

import "fmt"

const (
	currentYear = 2021 + iota
	year1       = currentYear + iota
	year2       = currentYear + iota
	year3       = currentYear + iota
	year4       = currentYear + iota
)

func main() {
	fmt.Println(currentYear, year1, year2, year3, year4)
}
