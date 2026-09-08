package main

import (
	"fmt"
	"reflect"
)

func main() {

	var fn1 Fn = func(i1, i2 int) int {
		return i1 + i2
	}

	r := fn1(10, 20)
	println(r)

	s := fn1.ToString(10, 20)

	fmt.Println("result:", s, "Type:", reflect.TypeOf(s))

	fns := FuncStruct{a: 10, b: 20}

	fns.FnS = func() int {
		return fns.a + fns.b
	}

	rs := fns.ToString()

	fmt.Println("result:", rs, "Type:", reflect.TypeOf(rs))
}

type Fn func(int, int) int

func (f Fn) ToString(a int, b int) string {
	return fmt.Sprint(f(a, b))
}

type FnS func() int

func (f FnS) ToString() string {
	return fmt.Sprint(f())
}

type FuncStruct struct {
	a, b int
	FnS  //promoted filed
}
