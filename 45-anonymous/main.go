package main

import "fmt"

func main() {

	func() {
		println("Hello World!")
	}() // () executor

	r := func(a, b int) int {
		return a + b
	}(10, 20)

	println(r)

	rf := func(a, b int) int {
		return a + b
	}
	println(rf(10, 20))

	var genfibs func(uint) []int

	genfibs = func(r uint) []int {
		a, b := 0, 1
		fibs := []int{} // it is not nil, stil the len is 0
		for range r {
			fibs = append(fibs, a)
			a, b = b, a+b
		}
		return fibs
	}

	figs := genfibs(10)
	fmt.Println(figs)

}

// inline
// named funcs
// anonymous func // closures
