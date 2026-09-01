package main

import (
	"errors"
	"fmt"
)

func main() {

	mymap1 := make(MyMap)

	mymap1["560086"] = "Bangalore-1"
	mymap1["560096"] = "Bangalore-2"
	mymap1["560034"] = "Bangalore-3"
	mymap1["522001"] = "Guntur-1"
	mymap1["522002"] = "Guntur-2"

	for k, v := range mymap1 {
		fmt.Println("Key:", k, "Value:", v)
	}

	if err := mymap1.Delete("nokey"); err != nil {
		println(err.Error())
	} else {
		println("successfully deleted")
	}

	keys, values := mymap1.GetKeysnValues()

	fmt.Println("Keys:", keys)

	fmt.Println("Values:", values)

	map2 := make(map[string]any)

	map2["560086"] = "Bangalore-1"
	map2["560096"] = "Bangalore-2"
	map2["560034"] = "Bangalore-3"
	map2["522001"] = "Guntur-1"
	map2["522002"] = "Guntur-2"

	println("normal map casting to MyMap")
	for k, v := range map2 {
		fmt.Println("Key:", k, "Value:", v)
	}

	if err := MyMap(map2).Delete("nokey"); err != nil {
		println(err.Error())
	} else {
		println("successfully deleted")
	}

	keys, values = MyMap(map2).GetKeysnValues()

	fmt.Println("Keys:", keys)

	fmt.Println("Values:", values)

}

type MyMap map[string]any // even though MyMap is a user defined bype, built in functions can be used on it
// make, delete etc.. range loop works
// which ever the built in functions can be called on map can also be called on MyMap

func (m MyMap) Delete(k string) error {
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

func (m MyMap) GetKeysnValues() (keys []string, values []any) {
	for k, v := range m {
		keys = append(keys, k)
		values = append(values, v)
	}
	return keys, values
}
