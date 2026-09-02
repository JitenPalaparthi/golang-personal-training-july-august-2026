package main

func main() {

	r1 := NewRectC(123.43, 76.45)
	r2 := NewRectC(123.43, 76.45)
	r3 := NewRect(123.43, 76.45)
	s1 := NewSquare(56.6)
	s2 := NewSquare(13.6)
	s3 := NewSquare(89.6)

	ShapeOf(r1)
	ShapeOf(r2)
	ShapeOf(r3)
	ShapeOf(s1)
	ShapeOf(s2)
	ShapeOf(s3)

	rect := NewRectC(43.45, 34.54)
	rect.Something()

	ishape1 := NewRect(43.45, 34.54)
	rect1, ok := ishape1.(*Rect)
	if ok {
		rect1.Something()
	}
}

// 1. Always go with interface driven approach
// 2. Easy for unit testing
// 3. Better design
