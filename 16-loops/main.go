package main

import "math/rand/v2"

func main() {

	for i := 1; i <= 10; i++ {
		print(i, " ")
	}

	println()

	a, b := 0, 1

	i := 1

	for i <= 10 { // kind of while loop but using for syntx
		print(a, " ")
		a, b = b, a+b // swap in a single line
		i++
	}
	println()

	i = 1

	for ; i <= 10; i++ {

		if i%2 == 0 {
			print(i, " ")
		}
	}
	println()

	for i := 1; i <= 10; {
		if i%2 == 0 {
			i++
			continue
		}
		print(i, " ")
		i++
	}
	println()
	// unconditional loop
	i = 1
	for {

		num := rand.IntN(1000)
		if num%2 == 0 {
			println("Even:", num)
		} else {
			println("Odd:", num)
		}

		if num%13 == 0 {
			break
		}

		i++
	}

	for i, j := 1, 10; i < j && true; i, j = i+1, j-1 {
		println("i:", i, "j:", j)
	}

	println("nested loops")

	for i := 1; i <= 5; i++ {
		for j := 1; j <= 5; j++ {
			if j == 3 {
				break
			}
			println("i:", i, "j:", j)
		}
	}

	println("nested loops break both the loops")

	done := false
	for i := 1; i <= 5; i++ {
		if done {
			break
		}
		for j := 1; j <= 5; j++ {
			if j == 3 {
				done = true
				break
			}
			println("i:", i, "j:", j)
		}
	}

	println("nested loops break both the loops using a label")

exit:
	for i := 1; i <= 5; i++ {
		for j := 1; j <= 5; j++ {
			for k := 1; k <= 5; k++ {
				if k == 3 {
					break exit
				}
				println("i:", i, "j:", j, "k:", k)
			}

		}
	}
}

// 1.22

// only for loop
// for collections , maps and channles there is for range loop
