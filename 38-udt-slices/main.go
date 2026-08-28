package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {

	si1 := make(SliceInt, 10)

	si1.FillRand(1000)

	fmt.Println(si1)

	max := si1.Max()
	println("Max:", max)

	min := si1.Min()
	println("Min:", min)

	slice1 := []int{34, 42, 49, 19, 5, 72, 91, 39, 36, 0}

	fmt.Println(slice1)

	max = SliceInt(slice1).Max()
	println("Max:", max)

	min = SliceInt(slice1).Min()
	println("Min:", min)

}

type SliceInt []int

func (si SliceInt) FillRand(r int) {
	for i := range si {
		si[i] = rand.IntN(r)
	}
}

func (si SliceInt) Max() int {
	if len(si) == 0 {
		return 0
	}
	max := si[0]
	for _, v := range si {
		if v > max {
			max = v
		}
	}
	return max
}

func (si SliceInt) Min() int {
	if len(si) == 0 {
		return 0
	}
	min := si[0]
	for _, v := range si {
		if v < min {
			min = v
		}
	}
	return min
}
