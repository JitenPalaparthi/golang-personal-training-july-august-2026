package main

import "math/rand/v2"

func main() {

	r1 := NewRect(123.43, 76.45)
	r2 := NewRect(123.43, 76.45)
	r3 := NewRect(123.43, 76.45)
	s1 := NewSquare(56.6)
	s2 := NewSquare(13.6)
	s3 := NewSquare(89.6)

	shapeSlice := []IShape{r1, r2, r3, s1, s2, s3}

	// for _, shape := range shapeSlice {
	// 	ShapeOf(shape)
	// }

	for range 6 {
		i := rand.IntN(6)
		ShapeOf(shapeSlice[i])
	}
}

// ITable is used for the dynamic dispatch
// For every concrete type, (Rect and Square) a table is created in mermory
// the table maintains the type info and also func pointers etc/
// what a method is called on a interface object, it finds the
// original concret type object function and executes
// The most imp thing is at runtime
