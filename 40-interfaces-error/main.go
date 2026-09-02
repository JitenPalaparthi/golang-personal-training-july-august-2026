package main

import (
	"fmt"
	"os"
)

func main() {
	str := "Hello World! How are are you doing!"
	//fmt.Println(str)
	fmt.Fprintln(os.Stdout, str)

	//fops := NewFileOps("data.txt")

	var fops1 *FileOps
	_, err := fmt.Fprintln(fops1, str)
	if err != nil {
		fmt.Println(err)
	}

	fops2 := NewFileOps("")
	_, err = fmt.Fprintln(fops2, str)
	if err != nil {
		//fmt.Println(err)

		fileErr, ok := err.(*FileError)
		if ok {
			fmt.Println("These are the error details from the object")
			fmt.Println("Error code:", fileErr.Code)
			fmt.Println("Error Message:", fileErr.Msg)
		} else {
			fmt.Println(err.Error())
		}

	}
	//fmt.Fprintln(fops, "I am trying to learn Golang, Wish me the best!")
}

/*
type Writer interface {
	Write(p []byte) (n int, err error)
}

type error interface {
    Error() string
}
*/

type FileOps struct {
	FileName string
}

func NewFileOps(fn string) *FileOps {
	return &FileOps{fn}
}

func (fo *FileOps) Write(bytes []byte) (n1 int, err error) {

	if fo == nil {
		return 0, NewFileError(101, "nil File Obj")
	}

	if fo.FileName == "" {
		return 0, NewFileError(102, "invalid or empty file name")
	}

	file, err := os.OpenFile(fo.FileName, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	n1, err = file.Write(bytes)
	return n1, err
}

type FileError struct {
	Code int
	Msg  string
}

func NewFileError(code int, msg string) *FileError {
	return &FileError{code, msg}
}

func (fe *FileError) Error() string {
	return fmt.Sprintln("Code:", fe.Code, "Message:", fe.Msg)
}
