package main

import "math/rand/v2"

func main() {
	day := uint(rand.IntN(10))

	println(day)

	switch day {
	case 1:
		println("Sunday")
	case 2:
		println("Monday")
	case 3:
		println("Tuesday")
	case 4:
		println("Wednesday")
	case 5:
		println("Thursday")
	case 6:
		println("Friday")
	case 7:
		println("Saturday")
	default:
		println("no day")
	}

}
