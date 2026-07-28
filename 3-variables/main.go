package main

import (
	"fmt"
	"reflect"
)

func main() {

	// what about chars

	str1 := "Hello World" // 11 chars
	str2 := "Hello 世界"    // 8 chars

	// len and size
	println("len of str1:", len(str1))
	println("len of str2:", len(str2))

	// Len is not going to give , how many chars , it gives how many bytes
	// each char can be 1 -4 bytes based on the UTF-8 chars
	// string is always a collection of bytes

	// what is a char in go

	var char1 uint8 = 'A'  // <=255
	var char2 uint16 = '世' // 19990
	var char3 int32 = '💙'  // 4 byte chars
	println(char1, char2, char3)

	//rune

	var char4 rune = 'A'
	char4 = '世'
	char4 = '💙'

	println(char4)
	fmt.Println("Type of char4:", reflect.TypeOf(char4))

	var char5 char = '💙'
	fmt.Println("Type of char5:", reflect.TypeOf(char5))

	var age uint8 = 42

	var bage byte = 42

	var sbyte single = 42

	fmt.Println("Type of age", reflect.TypeOf(age), "Type of bage:", reflect.TypeOf(bage), "Type of sbyte:", reflect.TypeOf(sbyte))
}

type char = int32 // creating an alias
type single = uint8
