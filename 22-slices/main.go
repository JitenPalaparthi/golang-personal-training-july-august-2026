package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {

	slice1 := make([]int, 10) // Ptr: 0x689090 Len:10 Cap:10
	fmt.Printf("Outside Func: %v address:%p ptr:%p Len:%d Cap:%d\n", slice1, &slice1, &slice1[0], len(slice1), cap(slice1))
	FillSlice(slice1) // The header is copied to the slice in the func
	fmt.Println(slice1)
	println()
	slice2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} // len = 10 and cap =10
	fmt.Printf("Outside Func: %v address:%p ptr:%p Len:%d Cap:%d\n", slice2, &slice2, &slice2[0], len(slice2), cap(slice2))
	AddElemDouble(slice2, 11, 12, 13, 14, 15)
	fmt.Println(slice2)
	println()
	slice3 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} // len = 10 and cap =10
	fmt.Printf("Outside Func: %v address:%p ptr:%p Len:%d Cap:%d\n", slice3, &slice3, &slice3[0], len(slice3), cap(slice3))
	AddElemDoubleP(&slice3, 11, 12, 13, 14, 15)
	fmt.Println(slice3)
}

func FillSlice(slice []int) {
	for i := range slice {
		slice[i] = rand.IntN(1000)
	}
	fmt.Printf("Inside Func: %v address:%p ptr:%p len:%d cap:%d\n", slice, slice, &slice[0], len(slice), cap(slice))
}

func AddElemDouble(slice []int, nums ...int) { // 0x72d4e3fb80a0
	for _, v := range nums {
		slice = append(slice, v) // 0x72d4e3fc2000// either changes the len or changes all three ptr,len,cap)
	}

	for i, v := range slice {
		slice[i] = v * 2
	}
	fmt.Printf("Inside Func: %v address:%p ptr:%p len:%d cap:%d\n", slice, &slice, &slice[0], len(slice), cap(slice))
}

func AddElemDoubleP(slice *[]int, nums ...int) { // 0x72d4e3fb80a0
	for _, v := range nums {
		*slice = append(*slice, v) // 0x72d4e3fc2000// either changes the len or changes all three ptr,len,cap)
	}

	for i, v := range *slice {
		(*slice)[i] = v * 2
	}
	fmt.Printf("Inside Func: %v address:%p ptr:%p len:%d cap:%d\n", slice, &slice, &(*slice)[0], len(*slice), cap(*slice))
}

// address:0x69e78ae36000 ptr:0x69e78ae34050 Len:10 Cap:10
// address:0x69e78ae36030 ptr:0x69e78ae34050 len:10 cap:10
