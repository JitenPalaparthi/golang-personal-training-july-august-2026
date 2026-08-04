package main

func main() {

	str1 := "Hello World" // collection of bytes

	for i := 0; i < len(str1); i++ {
		print(string(str1[i]), " ")
	}
	println()

	str2 := "Hello 世界"

	for i := 0; i < len(str2); i++ {
		print(string(str2[i]), " ")
	}
	println()

	for i, v := range str1 {
		println("index:", i, "value:", string(v))
	}

	println()

	for i, v := range str2 {
		println("index:", i, "value:", string(v))
	}

	for _, v := range str2 {
		println("value:", string(v))
	}
	for i := range str2 {
		// for i, _ := range str2 {
		println("index:", i)
	}

	for i := range 10 {
		if i%2 == 0 {
			println(i, "is even")
		} else {
			println(i, "is odd")
		}
	}
	a, b := 0, 1

	for range 10 {
		print(a, " ")
		a, b = b, a+b
	}

	println()

}
