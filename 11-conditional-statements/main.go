package main

import (
	"errors"
	"fmt"
	"math/big"
	"reflect"
	"strconv"
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
		fmt.Printf("Sum:%v\n", sum)
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

	if r, err := SumOf(true, true); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Sum of bool:", r)
	}

	if r, err := SumOf("123.23", "765.34"); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Sum of strings of floats:", r)
	}

	if r, err := SumOf("123", "765"); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Sum of strings of ints:", r)
	}

}

// SumOf(a , b any) (any,error) {
func SumOf(a any, b any) (sum any, err error) { // named return type
	// arthimetic operation
	if reflect.TypeOf(a) != reflect.TypeOf(b) {
		return 0, errors.New("a and b are two different types.So invalid for Sum")
	}
	// check both of them are number types
	if v, ok := a.(int); ok {
		sum = float64(v + b.(int))
		return sum, nil
	} else if _, ok := a.(uint); ok {
		sum = float64(a.(uint) + b.(uint))
		return sum, nil
	} else if _, ok := a.(uint8); ok {
		sum = float64(a.(uint8) + b.(uint8))
		return sum, nil
	} else if _, ok := a.(uint16); ok {
		sum = float64(a.(uint16) + b.(uint16))
		return sum, nil
	} else if _, ok := a.(uint32); ok {
		sum = float64(a.(uint32) + b.(uint32))
		return sum, nil
	} else if _, ok := a.(uint64); ok {
		bi1, bi2 := new(big.Int), new(big.Int)
		bi1.SetUint64(a.(uint64))
		bi2.SetUint64(b.(uint64))
		biSum := new(big.Int)
		biSum.Add(bi1, bi2)
		return biSum.Text(10), nil
	} else if _, ok := a.(int8); ok {
		sum = float64(a.(int8) + b.(int8))
		return sum, nil
	} else if _, ok := a.(int16); ok {
		sum = float64(a.(int16) + b.(int16))
		return sum, nil
	} else if _, ok := a.(int32); ok {
		sum = float64(a.(int32) + b.(int32))
		return sum, nil
	} else if _, ok := a.(int64); ok {
		bi1, bi2 := new(big.Int), new(big.Int)
		bi1.SetInt64(a.(int64))
		bi2.SetInt64(b.(int64))
		biSum := new(big.Int)
		biSum.Add(bi1, bi2)
		return biSum.Text(10), nil
	} else if _, ok := a.(float32); ok {
		//sum = float64(a.(float32) + b.(float32))
		//return sum, nil
		bi1, bi2 := new(big.Float), new(big.Float)
		bi1.SetFloat64(float64(a.(float32)))
		bi2.SetFloat64(float64(b.(float32)))
		biSum := new(big.Float)
		biSum.Add(bi1, bi2)
		return biSum.String(), nil
	} else if _, ok := a.(float64); ok {
		bi1, bi2 := new(big.Float), new(big.Float)
		bi1.SetFloat64(a.(float64))
		bi2.SetFloat64(b.(float64))
		biSum := new(big.Float)
		biSum.Add(bi1, bi2)
		return biSum.String(), nil
	} else {
		if reflect.TypeOf(a).String() == "string" {
			if a1, err := strconv.ParseFloat(a.(string), 64); err == nil {
				if b1, err := strconv.ParseFloat(b.(string), 64); err == nil {
					sum = a1 + b1
					return sum, nil
				} else {
					return 0, err
				}

			} else {
				if a1, err := strconv.Atoi(a.(string)); err == nil {
					if b1, err := strconv.Atoi(b.(string)); err == nil {
						sum = float64(a1 + b1)
						return sum, nil
					} else {
						return 0, err
					}
				}
			}
		}

		if _, ok := a.(bool); ok {
			if a.(bool) == true && b.(bool) == true {
				return float64(2), nil
			} else if a.(bool) == false && b.(bool) == true {
				return float64(1), nil
			} else if a.(bool) == true && b.(bool) == false {
				return float64(1), nil
			} else {
				return 0, nil
			}
		}

		return 0, fmt.Errorf("invalid type(s)")
	}
}
