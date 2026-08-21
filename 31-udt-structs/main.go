package main

import (
	"fmt"
	"unsafe"
)

func main() {

	p1 := Person{Id: 1231, Name: "Jiten", Email: "Jitenp@outlook.com", Address: Address{City: "Mumbai", Pincode: "420011"}}
	fmt.Println(p1)
	p1.Address.Pincode = "400011"
	fmt.Println("Pincode:", p1.Address.Pincode)

	p2 := Person{Id: 100}

	fmt.Println(p2)

	e1 := Employee{Id: 1231, Name: "Jiten", Email: "Jitenp@outlook.com", Address: Address{City: "Mumbai", Pincode: "420011"}}

	fmt.Println(e1)

	e1.Pincode = "400011" // bcz Address is a promoted field

	fmt.Println(e1)
	fmt.Println(e1.City)

	var ep1 Empty
	var ep2 Empty
	fmt.Println(ep1, ep2)

	fmt.Printf("Size of ep1:%d address of ep1:%p Size of ep2:%d address of ep2:%p\n", unsafe.Sizeof(ep1), &ep1, unsafe.Sizeof(ep2), &ep2)

}

type Person struct {
	Id      int
	Name    string
	Email   string
	Address Address // can give the name of the filed is name of the type
}

type Address struct {
	City    string
	Pincode string
}

type Employee struct {
	Id      int
	Name    string
	Email   string
	Address // promoted field
}

type Empty struct{}
