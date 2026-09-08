package main

import "fmt"

func main() {

	funcMap := make(map[string]any)

	funcMap["add"] = func(a, b int) int {
		return a + b
	}

	funcMap["sub"] = sub

	funcMap["greet"] = func() {
		fmt.Println("Hello World")
	}

	funcMap["value"] = 10000

	for k, v := range funcMap {
		a, b := 20, 10
		switch v.(type) {
		case func(int, int) int:
			r := v.(func(int, int) int)(a, b)
			println("Result of ", k, ":", r)

		case func():
			println("Executing func", k)
			v.(func())()
		default:
			println("untracked type")
		}

	}

}

func sub(i, j int) int {
	return i - j
}
