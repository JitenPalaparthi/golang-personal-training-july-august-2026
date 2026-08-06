package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

var (
	greetfunc func() = greet
)

func main() {

	var num1 int = 100
	var ptr1 *int = &num1

	var byte2 byte = 56
	var ptr2 *byte = &byte2

	fmt.Printf("Address that ptr1 holds:%p\n", ptr1)
	fmt.Printf("Value that ptr1 holds:%d\n", *ptr1)
	fmt.Printf("Type of  ptr1 :%T\n", ptr1)
	fmt.Printf("Size of  ptr1 :%d\n", unsafe.Sizeof(ptr1))
	*ptr1 = 200
	*(&num1) = 300
	fmt.Printf("Value that ptr1 holds:%d\n", *ptr1)

	fmt.Printf("Address that ptr2 holds:%p\n", ptr2)
	fmt.Printf("Value that ptr2 holds:%d\n", *ptr2)
	fmt.Printf("Type of  ptr2 :%T\n", ptr2)
	fmt.Printf("Size of  ptr2 :%d\n", unsafe.Sizeof(ptr2))

	var ptr3 *string // this holds the address of string

	// what is the zero value of pointer?

	if ptr3 == nil {
		println("nil pointer", ptr3)
	}

	// *ptr3 = "Hello World, how are you" // nil pointer dereference

	str1 := "Hello World"
	ptr3 = &str1
	ptr3 = nil // technically it is not free, bcz might have not been allocated in help, generally the word free is used when heap is allocated
	//*ptr3 = "Hello World, how are you"

	ptr4 := new(int) // it allocates the memory to the size of int and assigns a zero value to it
	ptr5 := new(12312.123)
	println(*ptr4, *ptr5)
	*ptr5 = 23432.54

	arr1 := [5]int{10, 76, 4, 59, 8}
	fmt.Println("Type of arr1:", reflect.TypeOf(arr1))

	var ptrarr1 *[5]int = &arr1
	var arrptr1 [5]*int

	fmt.Println(*ptrarr1)

	sum := 0
	for i := 0; i < len(*ptrarr1); i++ {
		sum += (*ptrarr1)[i]
	}

	println("sum:", sum)

	fmt.Println(arrptr1)

	numarr1 := 100
	numarr2 := 200
	ptrarr3 := new(300)
	ptrarr4 := new(int)
	*ptrarr4 = 400
	numarr5 := float64(87665.435)

	arrptr1[0] = &numarr1
	arrptr1[1] = &numarr2
	arrptr1[2] = ptrarr3
	arrptr1[3] = ptrarr4
	num := (int)(numarr5)
	arrptr1[4] = &num

	for i := 0; i < len(arrptr1); i++ {
		fmt.Printf("address: %d value: %d\n", arrptr1[i], *arrptr1[i])
	}

	greetfunc()

	//arr2 := [4]int{10, 76, 4, 59}
	// if ([4]int)(arr1) == arr2 {

	// }

	// var num5 = 100
	// var num6 uint64 = 100

	// if uint64(num5) == num6 {

	// }

	var ptr6 **int = &ptr1
	var ptr7 ***int = &ptr6

	**ptr6 = 45645645
	println(num1)
	***ptr7 = 6456456
	println(num1)

	if greet == nil {
		println("it is nil func")
	} else {
		println("greet is a named function and it is not nil")
	}

	funcptr := greet

	funcptr()
	greet()

	addfunc := add

	add(10, 20)
	addfunc(10, 20)

	fmt.Printf("address of greet:  %p\n", greet)
	fmt.Printf("address of funcptr:%p\n", funcptr)

}

func greet() {
	println("Hello World")
}

func add(a, b int) int {
	return a + b
}

// *ptr is called dereference

// nil pointer dereference
// use after free
