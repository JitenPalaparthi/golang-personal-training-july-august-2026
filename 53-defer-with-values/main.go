package main

func main() {

	num := 100

	defer func() {
		println("in func defer", num)
	}()

	defer func(num int) {
		println("in func withparam defer", num)
	}(num)

	defer println("in func with param defer", num)

	num += 1
	println("in main:", num)

	str := "Hello World!"

	for _, c := range str {
		defer println(string(c))
	}

}
