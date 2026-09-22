package main

import (
	"sync"
)

func main() {
	defer println("Done with main")
	//sig := make(chan bool)

	ch1 := Generate(10)

	workers := 50
	wg := new(sync.WaitGroup)
	for w := range workers {
		wg.Add(1)
		go Receive(wg, w+1, ch1)
	}
	wg.Wait()
}

func Generate(r uint) <-chan int { // This is called generator pattern
	ch := make(chan int)
	go func() {
		for i := range r {
			ch <- int(i+1) * int(i+1)
		}
		close(ch)
	}()

	return ch
}

func Receive(wg *sync.WaitGroup, worker int, ch <-chan int) {
	for v := range ch {
		println("--------------")
		println("worker", worker)
		switch {
		case v%8 == 0:
			println(v, "divisible by 8")
			fallthrough
		case v%4 == 0:
			println(v, "divisible by 4")
			fallthrough
		case v%2 == 0:
			println(v, "divisible by 2")
		default:
			println(v, "is not divisible by 2, 4 or 8")
		}
		println("--------------")
	}
	wg.Done()
}
