package main

import (
	"fmt"
	"reflect"
)

func main() {

	var num1 int = 100

	var num2 MyInt1 = 120

	var num3 MyInt2 = 140

	var num4 MyInt3 = 150

	var float1 float64 = 150.15

	str1 := MyInt1(num1).ToString()
	println(str1)

	sq1 := MyInt2(num1).Sq()
	println(sq1)

	cb1 := MyInt3(num1).Cube()
	println(cb1)

	str2 := num2.ToString()
	println(str2)

	sq2 := MyInt2(num2).Sq()
	println(sq2)

	cb2 := MyInt3(num2).Cube()
	println(cb2)

	str3 := MyInt1(num3).ToString()
	println(str3)

	sq3 := num3.Sq()
	println(sq3)

	cb3 := MyInt3(num3).Cube()
	println(cb3)

	str4 := MyInt1(num4).ToString()
	println(str4)

	sq4 := MyInt2(num4).Sq()
	println(sq4)

	cb4 := num4.Cube()
	println(cb4)

	str5 := MyInt1(float1).ToString()
	println(str5)

	sq5 := MyInt2(float1).Sq()
	println(sq5)

	cb5 := MyInt3(float1).Cube()
	println(cb5)

	// not possible bcz a bool cannot be type casted to any number type
	var ok1 bool = true

	// MyInt1(ok1).ToString()

	str6 := MyInt1(BToI(ok1)).ToString()
	fmt.Println("type:", reflect.TypeOf(str6), str6)

}

func BToI(b bool) int {
	if b {
		return 1
	}
	return 0
}

type MyInt1 int

type MyInt2 int

type MyInt3 MyInt1

//type FuncType func()

func (m MyInt1) ToString() string {
	return fmt.Sprint(m)
}

func (m MyInt2) Sq() int {
	return int(m * m)
}

func (m MyInt3) Cube() int {
	return int(m * m * m)
}
