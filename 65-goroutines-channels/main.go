package main

import "sync"

var (
	wg = new(sync.WaitGroup)
)

func main() {
	defer println("Done with main")
	defer wg.Wait()
	ch := make(chan int)

	r := 100

	wg.Add(1)
	go func(r int) {
		for i := range 100 {
			ch <- (i + 1) * (i + 1)
		}
		wg.Done()
	}(r)

	wg.Add(1)
	go func(r int) {
		for range 100 {
			println(<-ch)
		}
		wg.Done()
	}(r)

}
