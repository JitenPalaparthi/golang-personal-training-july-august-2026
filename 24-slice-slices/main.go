package main

import "fmt"

func main() {

	arr1 := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("address of 0th elem of array:%p\n", &arr1[0]) // 0x2ebf316ba000
	slice1 := arr1[:]
	fmt.Printf("Ptr of slice1:%p\n", &slice1[0]) //  0x2ebf316ba000

	slice2 := slice1[5:]
	slice3 := slice1[:5] // but not 5 from 0th index to the 4th index
	slice4 := slice1[3:8]
	slice5 := slice1 // The header of slice1 is copied to the header of slice5

	fmt.Println(slice2)
	fmt.Println(slice3)
	fmt.Println(slice4)
	fmt.Println(slice5)

	slice5[0] = 1000
	slice2[0] = 5000
	fmt.Println(slice1)

	// slice:=make([]int,5,8)
	// 1. It creates an array with 8 elements
	// 2. But the iteration is only for 5 elements
	// 3. upon append, if there is elements left in the cap it fills and changes the len
	// 4. if not, it would recrete an array and fill all previous array elements to the new array
	// 5. Assigns the new ptr, changes len and changes the cap
	// 6. Same process, again when append is called
}
