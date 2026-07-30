package main

import (
	"fmt"
	"math/big"
)

func main() {

	var x big.Int
	x.SetString("234324234324234234434234234234234234234", 10)

	fmt.Printf("Value of x:%v Type of x:%T\n", x.Text(10), x)

	var y big.Int

	fmt.Println(y)

	x1 := int64(1131231231231231231)
	y1 := int64(1231231232131232132)
	//.         2362462463362463363
	//11312312312312312311231231232131232132
	// var z big.Int
	// z.SetInt64(x1).Add(&z, (&z).SetInt64(y1))

	//fmt.Printf("Value of z:%v Type of z:%T\n", z.Text(10), z)

	s1 := fmt.Sprint(x1) + fmt.Sprint(y1)
	y.SetString(s1, 10)

	fmt.Printf("Value of y:%v Type of y:%T\n", y.Text(10), y)

	a := new(big.Int)
	a.SetInt64(x1)
	b := new(big.Int)
	b.SetInt64(y1)
	c := new(big.Int)

	c.Add(a, b)

	fmt.Printf("Value of c:%v Type of c:%T\n", c.Text(10), c)

	var num1 int8 = 127
	var num2 int8 = 127

	var num3 int16 = int16(num1 + num2)

	println(num3)

}
