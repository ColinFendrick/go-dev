package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}

type square struct {
	side float64
}

type shape interface {
	area() float64
}

func main() {
	s := square{
		side: 12,
	}

	c := circle{
		radius: 4,
	}

	info(s)
	info(c)
}

func (s square) area() float64 {
	return s.side * s.side
}

func (c circle) area() float64 {
	return math.Pi * c.radius
}

func info(s shape) {
	fmt.Println("Area is:", s.area())
}
