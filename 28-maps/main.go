package main

import "fmt"

func main() {

	var mymap map[string]string // declared but not instantiated, map is nil
	// What type can be a key -->?
	// any type that can implement == operator can be a key

	if mymap == nil {
		println("nil map")
	} else {
		println("non nil map")
	}

	mymap = make(map[string]string, 1000000)

	mymap["560086"] = "Bangalore-1"
	mymap["560096"] = "Bangalore-2"
	mymap["560034"] = "Bangalore-3"
	mymap["522001"] = "Guntur-1"
	mymap["522002"] = "Guntur-2"
	// 16 /5.0 -> 3.2 -> 32%

	for key, value := range mymap {
		fmt.Println("Key:", key, "Value:", value)
	}

	v := mymap["560086"]
	println(v)

	v, ok := mymap["5600861"]
	if ok {
		println(v)
	} else {
		println("key does not exist")
	}

	mymap1 := make(map[string]int)

	mymap1["Bangalore-1"] = 560086
	mymap1["Bangalore-2"] = 560096
	mymap1["Bangalore-3"] = 560034

	v1, ok := mymap1["Bangalore-11"]
	if ok {
		println(v1)
	} else {
		println("no key exists")
	}

}

// Map is not ordered
// Map is not thread safe by design
// Map can be nil
