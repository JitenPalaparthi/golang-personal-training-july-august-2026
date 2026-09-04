package main

import "unsafe"

func main() {

	var any1 any = 100
	var num int = any1.(int)
	println(num)

	println("size of any1:", unsafe.Sizeof(any1))

	var empty1 Empty = 100
	var num1 int = empty1.(int)
	println(num1)
	println("size of empty1:", unsafe.Sizeof(empty1))

	T1{}.Greet()
	T2{}.Greet()

}

type Empty interface {
}

// like any --> data pointer and type pointer

type T1 struct{}
type T2 struct{}

func (T1) Greet() {
	println("Hello World, I am from T1")
}

func (T2) Greet() {
	println("Hello World, I am from T2")
}

// a.Area(100)
// a.Area1(100,200)

// int Area(int a)
// int Area(int a,int b)

// // Monomorphization
// int Area(int a)
// int Area1(int a, int b)
