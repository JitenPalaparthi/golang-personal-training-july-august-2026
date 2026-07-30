package main

import (
	"fmt"
	"reflect"
)

func main() {

	c1 := 213.34 + 4i // type called complex number

	// complex128 and complex64 , built in function complex

	c2 := complex(123.123, 4.3)

	fmt.Println("type of c1:", reflect.TypeOf(c1))
	fmt.Println("type of c2:", reflect.TypeOf(c2))

	c3 := complex(float32(123.123), float32(4.3))
	fmt.Println("type of c3:", reflect.TypeOf(c3))

	c4 := c1 + c2 + complex128(c3)
	fmt.Println(c4)
	fmt.Printf("%.2f\n", c4)

	fmt.Printf("real of c4: %.3f imaginary of c4: %.2f\n", real(c4), imag(c4))
}
