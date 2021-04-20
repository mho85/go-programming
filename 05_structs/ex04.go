package main

import "fmt"

func main() {
	movie := struct {
		title string
		year  int
	}{
		title: "Back to the future",
		year:  1985,
	}

	fmt.Println(movie)
}
