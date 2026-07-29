package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	var any1 any // Zero value of any is nil

	if any1 == nil {
		println("Yes any1 is nil")
	}

	fmt.Println("Value of any1:", any1, "Type of any1:", reflect.TypeOf(any1), "Size of any1:", unsafe.Sizeof(any1))

	any1 = 100

	fmt.Println("Value of any1:", any1, "Type of any1:", reflect.TypeOf(any1), "Size of any1:", unsafe.Sizeof(any1))

	any1 = 1232.23
	fmt.Println("Value of any1:", any1, "Type of any1:", reflect.TypeOf(any1), "Size of any1:", unsafe.Sizeof(any1))

	any1 = "Hello World"
	fmt.Println("Value of any1:", any1, "Type of any1:", reflect.TypeOf(any1), "Size of any1:", unsafe.Sizeof(any1))

	any1 = true

	fmt.Println("Value of any1:", any1, "Type of any1:", reflect.TypeOf(any1), "Size of any1:", unsafe.Sizeof(any1))

	var any2 interface{} = "Hello Go folks!" // string

	any1 = any2

	fmt.Println("Value of any1:", any1, "Type of any1:", reflect.TypeOf(any1), "Size of any1:", unsafe.Sizeof(any1))

}

// any or interface{}

// any can store any kind of data
// any does not contain its own type but it contains the type of the data which is assigned to it

// All the types --> numbers, string, bool and user defined data types are loaded into memory.

// 100005 fe0

/*
any header
----------
DataPtr -> 8 bytes The pointer to the original data assigned to any
TypePtr -> 8 bytes The pointer to the type that the data belongs to
*/
