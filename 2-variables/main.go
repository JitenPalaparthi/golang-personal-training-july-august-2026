package main

import (
	"fmt"
	"reflect"
)

const PI float32 = 3.14

func main() {

	// const PI float32 = 3.14
	println("Constant PI:", PI)

	var num1 int // once declared it must be used
	println("num1:", num1)
	num1 = 100
	println("num1:", num1)

	var num2 uint8 = 56
	println("num2:", num2)

	// Group of variables
	var (
		num3   uint16 = 12332
		num4   int32  = 233232
		num5   int64  = -213213323232
		num6   uint
		float1 float32 = 3343.34
		float2 float64 = -123213.123123
		ok1    bool    // zero value false
		str1   string  = "Hello World!"
	)

	fmt.Println("num3:", num3, "num4:", num4, "num5:", num5, "num6:", num6, "flaot1:", float1, "float2:", float2,
		"ok1:", ok1, "str1L", str1)

	// type inference , go compiler
	var num7 = 123213 // int
	var num8 = 10     // int
	var float3 = 334343.244234324434
	var float4 = 3.14  // float64
	var age uint8 = 42 // automatically take 8 bytes

	var (
		ok2  = true
		str2 = "Hello Golang folks"
	)

	fmt.Printf("num7:%d num8:%d float3:%.3f float4:%.2f ok2:%v str2:%s age:%d\n", num7, num8, float3, float4, ok2, str2, age)

	fmt.Printf("type of num7:%T type of num8:%T type of float3:%T\n", num7, num8, float3)
	fmt.Println("type of float4:", reflect.TypeOf(float4))

	// shorthand declaration

	a := 10
	b, c, d, e := 20, 13.45, true, "Hello World"

	println(a, b, c, d, e)

	a1, b1 := 10, 20 // created new variable called a1 and b1 with values, type inference
	t1 := a1         // t1 is again a new variable, so t1:=
	a1 = b1
	b1 = t1

	println(a1, b1)

	a1, b1 = b1, a1 //swap
	println(a1, b1)

	a2, b2, c2 := 10, 20, 30 // var (a2 int = 10 b2 int =20 c2 int =30)
	// var (
	// 	a2 int = 10
	// 	b2 int = 20
	// 	c2 int = 30
	// )

	a2, b2, c2 = b2, c2, a2 // swap of three variable
	println(a2, b2, c2)

}

func PIMultiplyBy2() {
	println("py * 2", PI*2)
}

// constants -> const --> immutable

// type of the variable

// Numbers -> Zero value is 0
// int,uint,int8,uint8,int16,uint16,int32,uint32,int64,uint64,float32,float64
// alias types: rune(int32), byte(uint8)

// Bool --> Zero value is false
// bool -> true | false

// string --> collection of bytes , Zero value is ""

// interface type --> Zero value is nil
// interface{} or any

// github.com/josharian/impl@v1.4.0
