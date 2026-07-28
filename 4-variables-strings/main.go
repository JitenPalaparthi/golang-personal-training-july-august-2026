package main

import (
	"fmt"
	"math/rand/v2"
	"unsafe"
)

var (
	n = 24 // DataSegment
)

const (
	TICK = 60        // RO
	MIN  = TICK * 60 // RO
	HOUR = MIN * 60  // RO
	N    = 24
	DAY  = N * HOUR // RO
)

func main() {

	var str1 string // empty string which is zero value // ptr:nil len:0
	if str1 == "" {
		println("empty string")
	}

	// if str1 == nil {

	// }

	println("size of str1:", unsafe.Sizeof(str1), "len:", len(str1))

	var str2 = "Hello World" // type inference
	// "Hello World" is called string literal
	println("size of str2:", unsafe.Sizeof(str2), "len:", len(str2))

	str2 = "Hello Universe"

	str3 := "Hello World " // short hand declaration

	println("size of str3:", unsafe.Sizeof(str3), "len:", len(str3))

	str4 := fmt.Sprint(rand.IntN(99999)) // formatter func
	println("size of str4:", unsafe.Sizeof(str4), "len:", len(str4))

	str2 = fmt.Sprint(rand.IntN(99999))

	str2 = fmt.Sprint("Hello World")

	Greet()
	//println("Hello World")
	str5 := str1 + str2 + str3 + str4 // concatination, this concatination happens at runtime
	println("size of str5:", unsafe.Sizeof(str5), "len:", len(str5))
	fmt.Printf("%p %p ", fmt.Sprint, Greet)
}

//go:noinline
func Greet() {
	println("Hello World")
}

/*
String Header
-------------
ptr // 8 bytes
len // 8 bytes
*/

// Go does not have null , but nil
// nil is very curious type go

// string, any, interface , *int , func , slice, map , chan
// internally they contain pointer

// string literals , constants and global variabls are strored in data segment
