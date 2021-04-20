package main

import "fmt"

func main() {
	fmt.Println(translate("VN", hello))
}

func translate(l string, f func(l string) string) string {
	return f(l)
}

func hello(l string) string {
	switch l {
	case "FR":
		return "bonjour"
	case "ENG":
		return "hello"
	case "DE":
		return "Güten Tag"
	case "VN":
		return "Xin chào"
	default:
		return "Unknown translation"
	}
}
