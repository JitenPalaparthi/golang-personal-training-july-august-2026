package main

import (
	"errors"
	"fmt"
	"reflect"
)

func main() {
	if sum, err := SumOf(12, 13); err != nil {
		println(err.Error())
	} else {
		fmt.Printf("Sum:%.2f\n", sum)
	}

	if sum, err := SumOf(float32(12.34), float32(13.34)); err != nil {
		println(err.Error())
	} else {
		fmt.Printf("Sum:%v\n", sum)
	}

	if sum, err := SumOf(1324232.34, 1332434.34); err != nil {
		println(err.Error())
	} else {
		fmt.Printf("Sum:%.3f\n", sum)
	}

	if sum, err := SumOf(float32(12.34), float64(13.34)); err != nil {
		println(err.Error())
	} else {
		fmt.Printf("Sum:%.2f\n", sum)
	}

	if sum, err := SumOf(int8(12), int8(13)); err != nil {
		println(err.Error())
	} else {
		fmt.Printf("Sum:%.2f\n", sum)
	}

	if sum, err := SumOf(true, false); err != nil {
		println(err.Error())
	} else {
		fmt.Printf("Sum:%.2f\n", sum)
	}

	r1 := SumOfG(10, 20)
	fmt.Printf("Sum:%d\n", r1)

	r2 := SumOfG(10.4, 20.5)
	fmt.Printf("Sum:%.2f\n", r2)

	// r3 := SumOfG(true, false) // the compiler catches
	// fmt.Printf("Sum:%.2f\n", r3)

}

// Type earser

func SumOf(a, b any) (float64, error) {
	if !AreSame(a, b) {
		return 0, errors.New("a and b are different types")
	}
	if !IsNumber(a) {
		return 0, errors.New("a and b are not number types")
	}

	sum := 0.0
	switch v1 := a.(type) {
	case uint:
		sum = float64(v1 + b.(uint))
		return sum, nil
	case int:
		sum = float64(v1 + b.(int))
		return sum, nil
	case int8:
		sum = float64(a.(int8) + b.(int8))
		return sum, nil
	case uint8:
		sum = float64(a.(uint8) + b.(uint8))
		return sum, nil

	case int16:
		sum = float64(a.(int16) + b.(int16))
		return sum, nil
	case uint16:
		sum = float64(a.(uint16) + b.(uint16))
		return sum, nil

	case int32:
		sum = float64(a.(int32) + b.(int32))
		return sum, nil
	case uint32:
		sum = float64(a.(uint32) + b.(uint32))
		return sum, nil

	case int64:
		sum = float64(a.(int64) + b.(int64))
		return sum, nil
	case uint64:
		sum = float64(a.(uint64) + b.(uint64))
		return sum, nil

	case float32:
		sum = float64(a.(float32) + b.(float32))
		return sum, nil
	case float64:
		sum = a.(float64) + b.(float64)
		return sum, nil
	default:
		return 0, fmt.Errorf("type of a:%T and type of b:%T are not number types to perform addition", a, b)
	}

}

func SumOfG[T int | uint | int8 | uint8 | int16 | uint16 | int32 | uint32 | int64 | uint64 | float32 | float64](a, b T) T {
	return a + b
}

func SumOfGI[T NumberType](a, b T) T {
	return a + b
}

type NumberType interface {
	int | uint | int8 | uint8 | int16 | uint16 | int32 | uint32 | int64 | uint64 | float32 | float64
}

func IsNumber(a any) bool {
	// println(reflect.TypeOf(a))
	// fmt.Printf("%T\n", a)
	switch a.(type) {
	case int, uint, int8, uint8, int16, uint16, int32, uint32, int64, uint64, float32, float64:
		return true
	default:
		return false
	}
}
func AreSame(a, b any) bool {
	if reflect.TypeOf(a) == reflect.TypeOf(b) {
		return true
	}
	return false
}
