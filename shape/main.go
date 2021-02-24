package main

import (
	"fmt"
	"math"
)

type triangle struct {
	height float64
	base   float64
}
type square struct {
	sideLength float64
}

type shape interface {
	getArea() float64
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func (s square) getArea() float64 {
	return math.Pow(s.sideLength, 2)
}

func printArea(s shape) {
	fmt.Println("Area ", s.getArea())
}
func main() {
	aSquare := square{
		sideLength: 5,
	}
	aTriangle := triangle{
		height: 2.1,
		base:   3.2,
	}

	printArea(aSquare)
	printArea(aTriangle)
}
