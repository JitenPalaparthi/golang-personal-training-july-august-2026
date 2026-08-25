package main

import (
	"fmt"
	"unsafe"
)

func main() {

	var any1 any = struct {
		No         int
		IsMarried  bool
		Name       string
		IsEmployed bool
		Email      string
	}{No: 101, IsMarried: true, Name: "Jiten", IsEmployed: true, Email: "Jiten"}

	switch e1 := any1.(type) {

	case struct {
		No         int
		Name       string
		Email      string
		IsMarried  bool
		IsEmployed bool
	}:
		fmt.Println(e1)

	case struct {
		No         int
		IsMarried  bool
		Name       string
		IsEmployed bool
		Email      string
	}:
		fmt.Println(e1)

	}

	e1 := struct {
		No         int
		IsMarried  bool
		Name       string
		IsEmployed bool
		Email      string
	}{No: 101, IsMarried: true, Name: "Jiten", IsEmployed: true, Email: "Jiten"}

	e2 := struct {
		No         int
		Email      string
		Name       string
		IsEmployed bool
		IsMarried  bool
	}{No: 101, IsMarried: true, Name: "Jiten", IsEmployed: true, Email: "Jiten"}

	fmt.Println("Size of e1:", unsafe.Sizeof(e1))
	fmt.Println("Size of e1:", unsafe.Sizeof(e2))

	var e3 T1 = T1{No: 101, IsMarried: true, Name: "Jiten", IsEmployed: true, Email: "Jiten"}
	var e4 T2 = T2{No: 101, IsMarried: true, Name: "Jiten", IsEmployed: true, Email: "Jiten"}

	fmt.Println("Size of e3, type of T1:", unsafe.Sizeof(e3))
	fmt.Println("Size of e4, type of T2:", unsafe.Sizeof(e4))
}

type T1 struct {
	No         int    // - - - - - - - - 8
	IsMarried  bool   // - - - - - - - - 8 --> 7 bytes are not usefull, we called as padded 7 bytes
	Name       string // - - - - - - - - 8 - - - - - - - - 8  sttuct{Ptr Len }
	IsEmployed bool   // - - - - - - - - 8 --> 7 bytes are not usefull, we called as padded 7 bytes
	Email      string // - - - - - - - - 8 - - - - - - - - 8
}

type T2 struct {
	No         int    // - - - - - - - - 8
	Email      string // - - - - - - - - 8 - - - - - - - - 8  sttuct{Ptr Len }
	Name       string // - - - - - - - - 8 - - - - - - - - 8  sttuct{Ptr Len }
	age        int32  //  - - - - - - - - 8
	IsEmployed bool   // 4 are remaining -->
	IsMarried  bool   // 3 bytes are padded , it does not create another 8 byte block, bcz the next field wants only 1 bytes it adjusts the padded bytes
	IsMajor    bool   // 2 bytes are padded
	CanVote    bool   // 1 byte is padded
}
