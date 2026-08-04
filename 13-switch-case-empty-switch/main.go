package main

import "math/rand/v2"

func main() {

	num := rand.IntN(500)
	//num = -200 //mutate the variable
	switch { // empty switch
	case num >= 0 && num <= 100:
		println(num, "num is between 0 to 100")
		EvenOrOdd(num)
	case num >= 101 && num <= 200:
		println(num, "num is between 101 to 200")
		EvenOrOdd(num)
	case num >= 201 && num <= 300:
		println(num, "num is between 201 to 300")
		EvenOrOdd(num)
	case num >= 301:
		println(num, "num is greater than or equal to 301")
		EvenOrOdd(num)
	default:
		println(num, "is a negative number")
		EvenOrOdd(num)
	}

	char := '世'

	switch char {
	case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
		println(string(char), "is vowel character")
	default:
		println(string(char), "is either consonent or a special utf-8 char")
	}

}

func EvenOrOdd(num int) {
	if num%2 == 0 {
		println(num, "is even number")
	} else {
		println(num, "is odd number")
	}
}
