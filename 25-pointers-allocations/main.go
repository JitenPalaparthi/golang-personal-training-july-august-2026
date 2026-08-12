package main

import (
	"fmt"
)

func main() {

	a := 10
	{
		b := 20
		c := add(a, b)
		//println(c)
		fmt.Println(c)
	}
	//println(c)

	c := new(int)

	addO(10, 20, c)

	println(*c)
	// slog.Info("Hello World")
	// slog.Info("Hello World", *c)
	cptr := addP(10, 20)
	println(*cptr)

	arr1 := [100]int{}
	//println(arr1)
	fmt.Println(arr1)

	arr1[0] = 100

	arr2 := [99999]int{}
	arr2[0] = 10

	slice1 := make([]int, 10)
	slice1[0] = 3214

	slice2 := make([]int, 99999)
	slice2[0] = 3214

}

//go:noinline
func add(a, b int) int {
	return a + b
}

//go:noinline
func addO(a, b int, out *int) {
	if out != nil {
		*out = a + b
	}
}

//go:noinline
func addP(a, b int) *int {
	//c := new(a + b)
	c := new(int)
	*c = a + b
	return c // dangling pointer handling in Go
}

// Nil pointer dereference
// Dangling Pointer -> Is a very big problam in c,c++ and not allowed in Rust(safe)

// These below problems are handled by GC
// Use after free
// Double Free
// Memory leak
