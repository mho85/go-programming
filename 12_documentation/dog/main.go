// Package dog is all about dogs
package dog

// Years calculate the equivalent human age for dogs
func Years(x int) int {
	return x * 7
}

// Dog has an age
type Dog struct {
	Age int
}
