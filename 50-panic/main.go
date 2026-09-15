package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	f, err := os.OpenFile("data.txt", os.O_RDONLY, 0644)

	if err != nil {
		panic(err.Error())
	}
	defer f.Close()

	bytes, err := io.ReadAll(f)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(string(bytes))

}

// panic is a builtin functin
// panic panics the callstack
// panic is caused by runtime / userdefined panic
// advantage of using panic, in contrast to fatal or os.Exit(2)
// -> panic can be recoved if required
// -> panic can call deffered functions
// -> cant do that using os.Exit(1) or Fatal
