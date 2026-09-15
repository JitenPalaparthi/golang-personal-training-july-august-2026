package main

func main() {

	func() { // func1
		func() { // func1.1
			var num = 0
			println(100 / num) // divide by zero
		}()
	}()

	func() {
		arr := [5]int{10, 43, 56, 98, 1}

		for i := 0; i <= len(arr); i++ {
			println(arr[i])
		}
	}()

	func() {
		var ptr *int

		*ptr = 100

		println(*ptr)
	}()

}

// need to comment the code that causes divide by zero to enable index out of range panic
// similarly need to comment both divide by zero and index out of range to enable invalid memory or nil pointer dereference panic
