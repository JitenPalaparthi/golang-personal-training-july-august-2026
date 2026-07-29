package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {

	var num1 uint8 = 42 //0x2A //0b00101010

	var num2 uint64 = uint64(num1) // need to type cast

	var num3 uint16 = 31232 // 0b1111010 00000000

	fmt.Printf("num2:%b num3:%b\n", num2, num3)

	num1 = uint8(num3)

	fmt.Printf("num1:%b->%d\n", num1, num1)

	num3 = 43534 // 10101010 00001110

	fmt.Printf("num3:%d num3:%b\n", num3, num3)

	num1 = uint8(num3)

	fmt.Printf("%d ", 0b00001110)
	println(num1)

	var char1 rune = 'A'

	var num4 uint64 = uint64(char1)

	println(num4)

	var float1 = 4343.5456 // float64
	num4 = uint64(float1)
	println(num4)

	float1 = float64(num3)

	float2 := float32(float1) // type is inferred by casting

	println(float2)

	var char2 rune = '世' // 19990

	println(char2, string(char2))

	ok1 := true // 1 byte

	//num1 = uint8(ok1) // cannot type caste
	println(ok1, "cannot be casted to any number type")
	// if ok1 {
	// 	num1 = 1
	// } else {
	// 	num1 = 0
	// }

	num5 := uint32(19990) // "19990"

	str1 := string(num5)
	// cast to string

	println(str1, string(num5)) // casting from number to string is it tries to caste it to UTF-8 char

	str2 := strconv.Itoa(int(num5)) //fmt.Sprint(num5) // kind of a formatter

	fmt.Println(str2, "type of str2:", reflect.TypeOf(str2))

	a, s := calc(20, 10)
	println(a, s)
	a, _ = calc(20, 10) // _ blank identifier
	println(a)
	_, s = calc(20, 10) // _ blank identifier
	println(s)
	calc(234, 34)
	_, _ = calc(234, 34)

	str3 := "7657657a"

	num6, err := strconv.Atoi(str3)

	if err != nil {
		fmt.Println(err.Error(), num6)
	} else {
		println(num6)
	}
}

func calc(a, b int) (int, int) {
	return a + b, a - b
}

// strconv.ParseFloat , strconv.ParseBool, strconv.ParseInt(another function)
// Please look into these functions try to conver from string to float, strong to bool
// and string to int
