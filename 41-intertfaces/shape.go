package main

import "fmt"

type IShape interface {
	Area() float64
	Perimeter() float64
	IWhat // can use another interface into IShape
}

type IWhat interface {
	What() string // definition
}

// consumer
// Dependency
// Concrete
func ShapeOf(ishape IShape) {
	fmt.Printf("Area of %s: %0.2f\n", ishape.What(), ishape.Area())
	fmt.Printf("Perimeter of %s :%0.2f\n", ishape.What(), ishape.Perimeter())
	println("------------------------------")
}
