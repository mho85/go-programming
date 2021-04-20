package main

import (
	"fmt"
	"math"
)

type SQUARE struct {
	l float64
	w float64
}

type CIRCLE struct {
	r float64
}

type SHAPE interface {
	area() float64
}

func main() {
	sq1 := SQUARE{
		l: 0.1,
		w: 0.2,
	}
	c1 := CIRCLE{
		r: 0.5,
	}
	info(sq1)
	info(c1)
}

func (sq SQUARE) area() float64 {
	return sq.l * sq.w
}

func (c CIRCLE) area() float64 {
	return math.Pi * math.Pow(c.r, 2)
}

func info(sh SHAPE) {
	fmt.Println("Area:", sh.area())
}
