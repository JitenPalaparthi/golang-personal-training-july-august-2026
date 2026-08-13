package main

import (
	"fmt"
	"unsafe"
)

func main() {

	arr1 := [5]int{11111, 222222, 30, 40, 50}

	fmt.Printf("Ptr:%p\n", &arr1[0])
	fmt.Printf("Ptr:%d\n", &arr1[0])
	//fmt.Printf("Ptr:0x%x\n", &arr1[0])
	//ptr := &arr1[0]
	//ptr += 8 // This is not allowed in Go

	unsafePtr1 := unsafe.Pointer(&arr1[0]) // 1

	ptr1 := (*int)(unsafePtr1) // 2

	//println(*ptr1)

	uintptr1 := uintptr(unsafePtr1) // 4

	unsafePtr2 := unsafe.Pointer(uintptr1) // 3

	fmt.Println(unsafePtr1, ptr1, uintptr1, unsafePtr2)

	uptr1 := uintptr(unsafe.Pointer(&arr1[0])) // 1 and 4
	uptr1 += 8                                 // Pointer Arth
	rptr1 := *(*int)(unsafe.Pointer(uptr1))    // 2 and 3
	println(rptr1)

	{
		uptr1 := uintptr(unsafe.Pointer(&arr1[0]))
		for range 40 {
			rptr1 := *(*int)(unsafe.Pointer(uptr1))
			println(rptr1)
			uptr1 += unsafe.Sizeof(int(0))
		}
	}

	str1 := "Hello World"

	// Ptr -> 8
	// Len -> 8

	// [2]int

	//uint2 := uintptr(unsafe.Pointer(&str1))
	//strPtrs := (*[2]int)(unsafe.Pointer(uint2))

	strPtrArr1 := (*[2]int)(unsafe.Pointer(&str1))

	fmt.Println(strPtrArr1)
	strPtrArr1[1] = 100
	fmt.Println(len(str1), str1)

	str2 := "Hello World how are you doing"
	strPtrArr2 := (*[2]int)(unsafe.Pointer(&str2))
	strPtrArr1[0] = strPtrArr2[0]
	strPtrArr1[1] = strPtrArr2[1]
	fmt.Println("str1->", len(str1), str1)

	slice1 := make([]int, 5, 10)
	slice1[0], slice1[1], slice1[2], slice1[3], slice1[4] = 10, 20, 30, 40, 50

	uptr2 := uintptr(unsafe.Pointer(&slice1))
	slicePtr := (*[3]int)(unsafe.Pointer(uptr2))
	fmt.Println(slicePtr)
	slicePtr[1] = 10000
	slicePtr[2] = 10000
	fmt.Println(slice1, cap(slice1))
}

// unsafe.Pointer
// uintptr

// 1.A pointer value of any type can be converted to a Pointer.
// 2.A Pointer can be converted to a pointer value of any type.
// 3.A uintptr can be converted to a Pointer.
// 4.A Pointer can be converted to a uintptr.
