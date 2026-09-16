package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(3)
	go func() { // main.func1
		println("Hello first Goroutine")
	}()
	go func() {

		a, b := 0, 1

		for range 10 {
			println(a)
			a, b = b, a+b
		}
	}()

	go func() {
		go func() {
			fmt.Println("Start time -->", time.Now())
			time.Sleep(time.Second * 3) // block the thread for 10 seconds
			fmt.Println("End Time -->", time.Now())
		}()
		i := 1
		for {

			if i%2 == 0 {
				println("Even:", i)
			}

			if i >= 20 {
				//return
				//break
				runtime.Goexit() // it exits gracefully
			}

			i++
		}

	}()
	println("Hello World")
	//time.Sleep(time.Millisecond * 1)
	runtime.Goexit() // would wait for all other gorountins to complete their execution before going exit
	// if you call it main, the main crashes ultimately
}

// main is also a goroutines
// main does not wait for other goroutine to completed their execution

// main runs on some M
// go creates a new G

// Each P contains a local queue
// a P selects the goroutine
// An OS thread(M) executes it
