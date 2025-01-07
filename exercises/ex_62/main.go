package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
}

type circle struct {
	radius float64
}

type square struct {
	side float64
}

func main() {
	fmt.Println("Practicing Interfaces!")
	c1 := circle{
		radius: 6.77,
	}
	s1 := square{
		side: 5,
	}

	printArea(s1)
	printArea(c1)

}

func printArea(s shape) {
	fmt.Println(s.area())
}

func (c circle) area() float64 {
	area := math.Pi * math.Pow(c.radius, 2)
	return area
}

func (s square) area() float64 {
	area := math.Pow(s.side, 2)
	return area
}
