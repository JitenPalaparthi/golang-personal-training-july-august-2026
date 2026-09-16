package main

import (
	"fmt"
	"sync"
)

var wg = new(sync.WaitGroup)

func main() {
	//defer wg.Wait() // it is asked to wait ? How long ? till the state becomes 0
	wg.Add(1) // the state of wg = 1
	go func() {
		defer wg.Done() // decrement the state
		a, b := 0, 1
		for i := range 10 {
			println("---->", i+1, "--->", a)
			a, b = b, a+b
		}
	}()

	wg.Add(1) // the state of wg = 2
	go func() {
		for i := range 10 {
			if i%2 == 0 {
				wg.Add(1)
				go func() {
					println("Even-->", i, "  ")
					wg.Done()
				}()
			}
		}
		wg.Done()
	}()

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		PrimeNumbers(10)
		wg.Done()
	}(wg)

	println("Hello World")
	// time.Sleep(time.Microsecond * 10)
	wg.Wait()
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n <= 3 {
		return true
	}

	if n%2 == 0 || n%3 == 0 {
		return false
	}
	for i := 5; i*i <= n; i += 6 {

		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true
}

func PrimeNumbers(n int) {
	count := 0
	candidate := 2
	for count < n {
		if isPrime(candidate) {
			count++
			fmt.Printf("%d: %d\n", count, candidate)
		}
		candidate++
	}
}
