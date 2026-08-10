package main

import (
	"fmt"
	"math/rand/v2"
	"unsafe"
)

func main() {

	var slice1 []int // no need to give the size. ptr:nil len:0 cap:0
	// slice1 is nil since it is just declared but not instantiated.
	if slice1 == nil {
		println("Slice1 is nil")
	}
	// to instantiate a slice, there are multiple ways.

	slice1 = make([]int, 10) // ptr:<some address> len:10 cap:10
	// what are the values of slice now..

	fmt.Printf("slice1: %v slice address:%p len:%d cap:%d ptr:%p\n", slice1, &slice1, len(slice1), cap(slice1), &slice1[0])

	slice2 := make([]int, 10, 20) // ptr:<some address> len:10 cap:10
	// what are the values of slice now..

	fmt.Printf("slice2: %v slice address:%p len:%d cap:%d ptr:%p\n", slice2, &slice2, len(slice2), cap(slice2), &slice2[0])

	for i := range 10 {
		slice1[i] = rand.IntN(100)
	}

	fmt.Println(slice1)

	sum := SumOf(slice1)
	println("Sum:", sum)

	for i := range slice2 {
		slice2[i] = rand.IntN(100)
	}
	fmt.Println(slice2)

	sum = SumOf(slice2)
	println("Sum:", sum)

	slice1 = append(slice1, 200)
	fmt.Printf("slice1: %v slice address:%p len:%d cap:%d ptr:%p\n", slice1, &slice1, len(slice1), cap(slice1), &slice1[0])

	slice2 = append(slice2, 300)
	fmt.Printf("slice2: %v slice address:%p len:%d cap:%d ptr:%p\n", slice2, &slice2, len(slice2), cap(slice2), &slice2[0])

	// for i := range 20 {
	// 	slice2 = append(slice2, i+10)
	// }
	for i := range 20 {
		slice2 = append(slice2, i+10)
		//fmt.Printf("slice address:%p len:%d cap:%d ptr:%p\n", &slice2, len(slice2), cap(slice2), &slice2[0])
	}
	fmt.Printf("slice:%v slice address:%p len:%d cap:%d ptr:%p\n", slice2, &slice2, len(slice2), cap(slice2), &slice2[0])

	slice3 := []int{10, 20, 30, 40, 50} // short hand declaration of slice
	fmt.Printf("slice3:%v slice address:%p len:%d cap:%d ptr:%p\n", slice3, &slice3, len(slice3), cap(slice3), &slice3[0])

	slice4 := []int{} // is slice nil or not?, this is considered as instantiation in go
	// ptr: ? len:0 cap:0
	if slice4 == nil {
		println("yes ")
	} else {
		println("not nil")
	}

	//var ptr1 *int = runtime.zerobase

	var slice5 []int

	slice5 = append(slice5, 10, 20, 30, 40, 50) // 1. append can append any number of elements
	// append can append even the slice is nil

	sum = SumOf(slice5)
	println("sum:", sum)

	fmt.Println(slice5)
	clear(slice5)
	fmt.Println(slice5)

	var t1 T
	fmt.Printf("%p size of T:%d", &t1, unsafe.Sizeof(t1))
}

type T struct{}

// sequence of elements stored together and the length is not fixed.
// can extend the slice at runtime, up to what ever extent you can (subjective to total memory available and allocatable)
// where slice is allocated, in stack or in heap --> any thing in go not only slice, can be allocated on slice or heap based on the compiler algo

/* SliceHeader
ptr:
len:
cap:
*/

func SumOf(slice []int) int {
	sum := 0

	for _, v := range slice {
		sum += v
	}
	return sum
}
