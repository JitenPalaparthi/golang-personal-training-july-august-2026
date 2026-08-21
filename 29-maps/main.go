package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
)

func main() {

	mymap1 := make(map[int]int)

	slice1 := make([]int, 100)

	for i := range len(slice1) {
		slice1[i] = rand.IntN(100)
	}
	fmt.Println(slice1)

	for _, v := range slice1 {

		_, ok := mymap1[v]

		if ok {
			mymap1[v] = mymap1[v] + 1
		} else {
			mymap1[v] = 1
		}

	}

	fmt.Println(mymap1)

	mymap2 := make(map[string]any)

	mymap2["560086"] = "Bangalore-1"
	mymap2["560096"] = "Bangalore-2"
	mymap2["560034"] = "Bangalore-3"
	mymap2["522001"] = "Guntur-1"
	mymap2["522002"] = "Guntur-2"

	fmt.Println(mymap2)

	delete(mymap2, "560086")
	println("------------------")
	fmt.Println(mymap2)

	delete(mymap2, "5600862")
	println("------------------")
	fmt.Println(mymap2)

	if err := Delete(mymap2, "560096"); err != nil {
		println(err.Error())
	} else {
		println("successfully deleted")
	}

	clear(mymap2) // it removes all keys and values

	fmt.Println(mymap2)

	var mymap3 map[string]any = map[string]any{"560086": "Blr-1", "560096": "Blr-2"}
	println(mymap3)
}

func Delete(m map[string]any, k string) error {
	if m == nil {
		return errors.New("input map is nil")
	}

	_, ok := m[k]
	if !ok {
		return errors.New("no key existed in the map")
	}

	delete(m, k)
	return nil
}
