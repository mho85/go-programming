package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Printf("OS: %v\n", runtime.GOOS)
	fmt.Printf("ARCH: %v\n", runtime.GOARCH)
}
