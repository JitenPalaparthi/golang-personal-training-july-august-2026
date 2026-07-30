package main

func main() {

	var any1 any // interface{} // interface{}

	// type: nil
	// value: nil

	var num1 uint8 = 59

	any1 = num1 // num1 has 59 of type uint8

	var num2 uint8 = any1.(uint8) // type assertion

	//var num3 uint16 = any1.(uint16) // type assertion

	println(num2)

	var num3 uint16

	num3, ok := any1.(uint16)

	if ok == true {
		println(num3)
	} else {
		println("failed to assert to uint16")
	}

	any1 = 100 // int

	if ok {
		println(num3)
	} else {

		num4, ok := any1.(uint8)

		if ok {
			println(num4)
		} else {
			num4, ok := any1.(int)
			if ok {
				println(num4)
			} else {
				println("failed to assert to int, uint8 and uint16")
			}
		}
	}

	// a := 10
	// b := a
	// b++
	// println(a, b)

	// c := &a

	// *c++
	// fmt.Println(a, *c)

}

// any header
// DataPtr and TypePtr

// type casting
// type conversion
// type assertion
