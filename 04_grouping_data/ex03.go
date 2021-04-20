package main

import "fmt"

func main() {
	s := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	s0 := s[:5]
	s1 := s[5:]
	s2 := s[2:7]
	s3 := s[1:6]
	ss := [][]int{s0, s1, s2, s3}

	for _, s := range ss {
		fmt.Printf("%v\n", s)
	}
}
