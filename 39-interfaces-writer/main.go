package main

import (
	"fmt"
	"os"
)

func main() {
	str := "Hello World! How are are you doing!"
	//fmt.Println(str)
	fmt.Fprintln(os.Stdout, str)
	fops := NewFileOps("data.txt")
	fmt.Fprintln(fops, str)

	fmt.Fprintln(fops, "I am trying to learn Golang, Wish me the best!")
}

/*
type Writer interface {
	Write(p []byte) (n int, err error)
}*/

type FileOps struct {
	FileName string
}

func NewFileOps(fn string) *FileOps {
	return &FileOps{fn}
}

func (fo *FileOps) Write(bytes []byte) (n1 int, err error) {
	file, err := os.OpenFile(fo.FileName, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	n1, err = file.Write(bytes)
	return n1, err
}
