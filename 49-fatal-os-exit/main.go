package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	f, err := os.OpenFile("data.txt", os.O_RDONLY, 0644)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
		// log.Fatalln(err.Error())
		// fatal error
		// println(err.Error())
		// return
	}
	defer f.Close()

	bytes, err := io.ReadAll(f)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(string(bytes))

}
