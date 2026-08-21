package main

import (
	"fmt"
)

func main() {

	var p1 Person // declaere a struct varialbe

	fmt.Println(p1) // zero values

	p1.Id = 2312
	p1.Email = "Jitenp@outlook.com"
	p1.Name = "Jiten"

	fmt.Println(p1)

	p2 := Person{Id: 1231, Name: "Jiten"}
	fmt.Println(p2)

	p3 := new(Person)

	p3.Id = 123
	p3.Name = "Jiten"
	(*p3).Email = "JitenP@Outlook.Com" // no need to deref and assign

	fmt.Println(*p3)
	//slog.Info(fmt.Sprint(*p3))

	p4 := &Person{Id: 1231, Name: "Jiten", Email: "JitenP@outlook.com"}
	fmt.Println(p4)

	p5 := new(Person{Id: 1231, Name: "Jiten", Email: "JitenP@outlook.com"}) // after go version 1.24

	fmt.Println(p5)

	cc1 := ColourCode{int: 9999, string: "red", astring: "red but not very red"}
	fmt.Println(cc1)

	var cc2 ColourCode

	cc2.int = 433
	cc2.string = "blue"
	cc2.astring = "some blue"

	fmt.Println(cc2)
}

type Person struct {
	Id    int
	Name  string
	Email string
}

type ColourCode struct { // struct with anonymous fields
	int
	string
	astring
}

type astring = string // not a new type but just alias to the existing type
