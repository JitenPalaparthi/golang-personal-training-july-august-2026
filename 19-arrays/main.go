package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	var arr1 [5]int // zero value for array
	fmt.Println(arr1)

	// var arr2 [4]int

	// if arr1 == arr2 {

	// }

	arr1[0], arr1[1], arr1[2], arr1[3], arr1[4] = 10, 20, 30, 40, 50

	sum := 0
	for i := 0; i < len(arr1); i++ {
		sum += arr1[i]
		println(arr1[i])
	}
	println(sum)

	arr2 := [5]int{10, 20, 30, 40, 50}
	arr3 := [...]int{23, 54, 34, 5, 34, 8, 9, 45, 6, 8, 8, 4}

	for i, v := range arr2 {
		println("index:", i, "value:", v)
	}

	for i, v := range arr3 {
		if v%2 != 0 {
			arr3[i] = rand.IntN(100)
		}
	}
	fmt.Println(arr3)

	sum = SumOf(arr1)
	println("Sum:", sum)

	sum = SumOf(arr2)
	println("Sum:", sum)

	// sum = SumOf(arr3)
	// println("Sum:", sum)

	arr2d := [2][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}}

	for i := 0; i < len(arr2d); i++ {
		for j := 0; j < len(arr2d[i]); j++ {
			print(arr2d[i][j], " ")
		}
		println()
	}

	for _, arr := range arr2d {
		for _, v := range arr {
			print(v, " ")
		}
		println()
	}

	arr3d := [2][2][2]int{{{1, 2}, {3, 4}}, {{5, 6}, {7, 8}}}

	for _, arr2d := range arr3d {
		for _, arr := range arr2d {
			for _, v := range arr {
				print(v, " ")
			}
			println()
		}
		println()
	}

	arra := [...]any{123.23, 34, true, "hello World", 'H'}
	fmt.Println(arra)

}

// the length is fixed and to be known to the compiler(by compile time)
//  is there any header for the array --> No
// how does it understand the length: the type of the array is including its length
// cannot type cast array
// unless arrays are used inside your data structures , we dont write functions with arrays as parameters

func SumOf(arr [5]int) int {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	return sum
}
