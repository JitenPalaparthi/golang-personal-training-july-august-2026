package main

import (
	"fmt"
	"os"
	"sync"
)

var wg = new(sync.WaitGroup)

func main() {
	defer wg.Wait()
	wg.Add(1)
	go func() {
		for i := 1; i <= 2000000; i++ {
			wg.Add(1)
			go func() {
				file, err := os.OpenFile("data.txt", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
				if err != nil {
					return
				}
				defer file.Close()
				if i%2 == 0 {
					print("even-->", i)
					//fmt.Fprintln(file, "even-->", i)
					file.WriteString(fmt.Sprintln("even-->", i))
				} else {
					print("odd-->", i)
					//fmt.Fprintln(file, "odd-->", i)
					file.WriteString(fmt.Sprintln("odd-->", i))
				}
				wg.Done()
			}()
		}
		wg.Done()
	}()
}
