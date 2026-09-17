package main

import (
	"sync"
)

var (
	Counter = 1 // Global variable , where is it be created , Data Segment
	wg      = new(sync.WaitGroup)
	mu      = new(sync.Mutex)

// It is a process level variable
)

func main() {
	wg.Add(1)
	go func() {
		for range 1000 {
			//time.Sleep(time.Millisecond * 5)
			mu.Lock()
			Counter++
			mu.Unlock()

		}
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		for range 1000 {
			//time.Sleep(time.Millisecond * 5)
			mu.Lock()
			Counter--
			mu.Unlock()
		}
		wg.Done()
	}()

	wg.Wait()
	println(Counter)
}
