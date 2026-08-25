package main

import "fmt"

func main() {

	r1 := Rect{100.12, 123.2}
	a1 := Area(r1)
	p1 := Perimeter(r1)

	fmt.Printf("Area:%.2f Perimeter:%0.2f\n", a1, p1)

	a1 = r1.Area()
	p1 = r1.Perimeter()

	fmt.Printf("Area:%.2f Perimeter:%0.2f\n", a1, p1)

}

type Rect struct {
	L, B float32
}

func Area(r Rect) float64 {
	return float64(r.L * r.B)
}

func Perimeter(r Rect) float64 {
	return 2 * float64(r.L+r.B)
}

func (r Rect) Area() float64 { // r is a receiver
	return float64(r.L * r.B)
}
