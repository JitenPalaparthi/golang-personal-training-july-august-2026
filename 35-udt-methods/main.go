package main

import "fmt"

func main() {

	r1 := Rect{L: 100.12, B: 123.2}

	r2 := Rect{100.12, 123.2, 0, 0}

	// call by value bcz the receiver is not a pointer
	a1 := r1.Area()
	p1 := r1.Perimeter()

	fmt.Printf("Area:%.2f Perimeter:%0.2f\n", a1, p1)
	fmt.Printf("Area:%.2f Perimeter:%0.2f\n", r1.A, r1.P)

	// call by reference bcz the receiver is a pointer
	a2 := r2.AreaR()
	p2 := r2.PerimeterR()

	fmt.Printf("Area:%.2f Perimeter:%0.2f\n", a2, p2)
	fmt.Printf("Area:%.2f Perimeter:%0.2f\n", r2.A, r2.P)

}

type Rect struct {
	L, B float32
	A, P float64
}

// by default or design all methods are call by values unless pointer receiver is used
func (r Rect) Area() float64 { // r is a receiver
	r.A = float64(r.L * r.B)
	return r.A
}

func (r Rect) Perimeter() float64 { // r is a receiver
	r.P = 2 * float64(r.L+r.B)
	return r.P
}

// pointer receiver is nothing but call by reference

func (r *Rect) AreaR() float64 { // r is a pointer receiver
	r.A = float64(r.L * r.B)
	return r.A
}

func (r *Rect) PerimeterR() float64 { // r is a pointer receiver
	r.P = 2 * float64(r.L+r.B)
	return r.P
}
