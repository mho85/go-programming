package main

import "fmt"

func main() {
	jb := []string{"James", "Bond", "Shaken, not stirred"}
	mp := []string{"Miss", "Moneypenny", "Helloooooo, James."}
	s := [][]string{jb, mp}

	for i, character := range s {
		fmt.Println(i)
		for _, data := range character {
			fmt.Printf("\t%v\n", data)
		}
	}
}
