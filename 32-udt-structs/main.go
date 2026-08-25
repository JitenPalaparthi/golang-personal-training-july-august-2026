package main

import "fmt"

func main() {

	e1 := Employee{Id: 101, Name: "Jiten", Email: "JitenP@Outlook.Com", Address: struct {
		City    string
		Pincode string
		Lines   struct {
			Line1 string
			Line2 string
		}
	}{City: "Trivandrum", Pincode: "695011", Lines: struct {
		Line1 string
		Line2 string
	}{Line1: "Kappil Lane", Line2: "Next to Abcd stores"}}}

	fmt.Println(e1)
	fmt.Println(e1.Address)
	fmt.Println(e1.Address.Lines)

	var e2 struct {
		Id    int
		Name  string
		Email string
	}

	e2 = struct {
		Id    int
		Name  string
		Email string
	}{Id: 101, Name: "Jiten", Email: "Jitenp@outlook.com"}

	fmt.Println(e2)

	var a1 any = struct {
		Id    int
		Name  string
		Email string
	}{Id: 101, Name: "Jiten", Email: "Jitenp@outlook.com"}

	var e3 struct {
		Id    int
		Name  string
		Email string
	} = a1.(struct {
		Id    int
		Name  string
		Email string
	})

	fmt.Println(e3)

	switch e1 := a1.(type) {

	case struct {
		Id    int
		Name  string
		Email string
	}:
		fmt.Println(e1)
	}

}

type Employee struct {
	Id      int
	Name    string
	Email   string
	Address struct { // embedded struct
		City    string
		Pincode string
		Lines   struct { // embedded struct
			Line1 string
			Line2 string
		}
	}
}
