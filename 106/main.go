package main

import (
	"fmt"
	"math"
)

type square struct {
	length float32
}

func (s square) area() float32 {
	return float32(math.Pow(float64(s.length), 2))
}

type circle struct {
	radius float32
}

func (c circle) area() float32 {
	return float32(math.Pow(float64(c.radius), 2)) * math.Pi
}

type shape interface {
	area() float32
}

func info(s shape) {
	fmt.Println(s.area())
}

func main() {
	c := circle{
		radius: 2,
	}
	info(c)

	s := square{
		length: 5,
	}
	info(s)
}
