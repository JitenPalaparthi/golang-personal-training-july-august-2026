package main

func main() {

	defer func() { // func1.1
		println("Hello World! main ends")
	}()

	println("Hello World, main starts")

	func() { // func1
		defer func() { // func1.1
			println("Hello World! func1 ends")
		}()
		defer func() { // func1.2
			println("Hello World! func1 ends")
		}()
		println("Hello World, func1 starts")
	}() // main.main.func1.func1.1

}

// defer is a keyword
// defer defers the execution of a function/method to the end of its caller
// defers maintain its own stackframes
