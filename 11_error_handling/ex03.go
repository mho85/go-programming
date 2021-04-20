package main

import "fmt"

type customErr struct {
}

func (ce *customErr) Error() string {
	return "Error: Oh la la"
}

func foo(err error) {
	fmt.Println(err.Error())
}

func main() {
	ce := customErr{}
	foo(&ce)
}
