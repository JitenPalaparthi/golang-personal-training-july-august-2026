package main

import (
	"fmt"
	"strconv"
)

func main() {

	var (
		num1 int     = -100
		num2 uint    = 234
		num3 uint16  = 2343
		num4 int16   = -3434
		num5 uint64  = 23423424342
		num6 float32 = 3243.344
		num7 float64 = 2343244.234234
		str1         = "23423.34"
		str2         = "342342"
		ok1          = true
		ok2          = false
		any1         = any(num4)
		any2         = any(num7)
		sum  float64
	)

	sum = float64(num1) + float64(num2) + float64(num3) + float64(num4) + float64(num5) + float64(num6) + num7

	s1, _ := strconv.ParseFloat(str1, 63) // if error the value of s will be 0, adding 0 has not effect on sum
	sum += s1

	s2, _ := strconv.ParseInt(str2, 10, 64)

	sum += float64(s2)

	if ok1 { // if true add 1
		sum += 1
	}

	if ok2 {
		sum += 1
	}

	sum += float64(any1.(int16))

	sum += any2.(float64)

	//println("sum = ", sum)
	fmt.Printf("Sum: %.3f\n", sum)
}
