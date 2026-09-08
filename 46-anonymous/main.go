package main

//go:noinline
func main() {

	r1 := calc(10, 20, func(i1, i2 int) int {

		return i1 + i2
	})
	println(r1)

	r4 := calc(3, 5, func(i1, i2 int) int {
		return func(a int) int {
			return a * a
		}(i1) + func(a int) int {
			return a * a
		}(i2)
	})
	println(r4)

	r2 := calc(20, 10, sub)

	println(r2)

	var fn func(int, int) int = func(i1, i2 int) int {
		return i1 * i2
	}

	r3 := calc(30, 4, fn)

	println(r3)

}

//go:noline
func calc(a, b int, fn func(int, int) int) int {
	return fn(a, b)
}

func sub(i, j int) int {
	return i - j
}
