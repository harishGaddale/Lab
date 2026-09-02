package main

import "fmt"

type triange struct {
	base   float64
	height float64
}

type square struct {
	side float64
}

type shape interface {
	getArea() float64
}

func (t triange) getArea() float64 {
	return 0.5 * t.base * t.height
}

func (s square) getArea() float64 {
	return s.side * s.side
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}
func main() {
	t := triange{5, 5}
	s := square{5}
	printArea(t)
	printArea(s)
}
