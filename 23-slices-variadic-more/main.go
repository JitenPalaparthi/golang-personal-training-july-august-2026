package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	slice1 := make([]int, 10)
	FillSlice(slice1)

	slice2 := make([]int, 10)
	copy(slice2, slice1) // Shallow copy, element by element copy
	// 1. slice2 should not be nil
	// 2. if slice1 len is greater than slice2, only the len for slice2 are copied
	// 3. is slice1 len is smaller than slice2, till the len of slice1 are copied and then the rest are default values

	slice3 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	slice4 := make([]int, 5)
	copy(slice4, slice3)
	fmt.Println(slice4)

	slice5 := make([]int, 20)
	copy(slice5, slice3)
	fmt.Println(slice5)

	println()
	println("Copy named function")

	slice6 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	slice7 := make([]int, 5)
	copy(slice7, slice6)
	fmt.Println(slice7)

	slice8 := make([]int, 20)
	copy(slice8, slice6)
	fmt.Println(slice8)

	// var nums... int //cant create variaic

	sum := SumOf("Some Sum func:")
	println(sum)

	sum = SumOf("Some Sum func:", 10, 20)
	println(sum)

	sum = SumOf("Some Sum func:", 10, 20, 34, 34, 45, 978, 6, 6, 45, 76)
	println(sum)

	sum = SumOf("Sum of Slice:", slice3...)
	println(sum)

	arr1 := [5]int{10, 20, 30, 40, 50}

	slice9 := make([]int, 5)
	Copy(slice9, arr1[:]) // This is a elem by elem copy

	var slice10 []int // slice10 is nil but append would instantiate the slice when it is nil

	slice10 = append(slice10, slice3...)
	fmt.Println("slice10", slice10)

	clear(slice10)
	fmt.Println("slice10", slice10)
}

func Copy(dst, src []int) {
	for i := range min(len(dst), len(src)) {
		dst[i] = src[i]
	}
}

func FillSlice(slice []int) {
	for i := range slice {
		slice[i] = rand.IntN(1000)
	}
	fmt.Printf("Inside Func: %v address:%p ptr:%p len:%d cap:%d\n", slice, slice, &slice[0], len(slice), cap(slice))
}

// Variadle parameter must be the last parameter in the function
// func SumOf( nums ...int,des string) int { // This does not work bcz the variadic must the be the last one
// the variadic parameter must be used only in functions or methods, cannot be userd as normal variable or field etc
func SumOf(des string, nums ...int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	println(des, sum)
	return sum
}
