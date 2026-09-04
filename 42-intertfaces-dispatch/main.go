package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {

	r1 := NewRect(123.43, 76.45)
	r2 := NewRect(123.43, 76.45)
	r3 := NewRect(123.43, 76.45)
	s1 := NewSquare(56.6)
	s2 := NewSquare(13.6)
	s3 := NewSquare(89.6)

	fmt.Printf("ptr of Rect Area:%p\n", r1.Area)
	fmt.Printf("ptr of Square Area:%p\n", s1.Area)

	//var ishape1 IShape = NewRect(45.4, 87.45)

	// ITab

	// ITab *abi.ITab. -> Itab for IShape,Rect)
	// Data unsafe.Pointer -> Rectangle Value
	/*
		     type ITab struct{
			 Inter *InterfaceType
			 Type *Type
			 Func [1]uintptr
			 }

			 IShape + Rect --> ITab A
			 IShape + Square --> ITab B

			 IWhat + Rect --> ITab C
			 IWhat + Square --> ITab D

	*/

	shapeSlice := []IShape{r1, r2, r3, s1, s2, s3}

	// for _, shape := range shapeSlice {
	// 	ShapeOf(shape)
	// }

	for range 6 {
		i := rand.IntN(6)
		ShapeOf(shapeSlice[i])
	}
}

// ITab is used for the dynamic dispatch
// For every concrete type, (Rect and Square) a table is created in mermory
// the table maintains the type info and also func pointers etc/
// what a method is called on a interface object, it finds the
// original concret type object function and executes
// The most imp thing is at runtime

// 1000a0ef0 T main.(*Rect).Area
// 1000a1210 T main.(*Rect).Area-fm
// 1000a1110 T main.(*Square).Area
// 1000a1230 T main.IShape.Area-fm
// 1000a10b0 T main.Square.Area
