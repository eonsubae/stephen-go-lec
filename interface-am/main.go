package main

import "fmt"

type shape interface {
	getArea() float64
}

type triangle struct {
	base   float64
	height float64
}

type square struct {
	sideLength float64
}

func (t *triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func (s *square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

func main() {
	t1 := triangle{
		base:   3.5,
		height: 10.75,
	}

	s1 := square{sideLength: 30.0}

	printArea(&t1)
	printArea(&s1)
}
