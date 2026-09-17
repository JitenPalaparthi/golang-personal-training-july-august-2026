package main

import (
	"sync"
)

var (
	Counter = 1 // Global variable , where is it be created , Data Segment
	wg      = new(sync.WaitGroup)
	mu      = new(sync.RWMutex)

// It is a process level variable
)

func main() {
	wg.Add(1)
	go func() {
		for range 1000 {
			wg.Add(1)
			go func() {
				mu.Lock()
				Counter++
				mu.Unlock()
				wg.Done()
			}()
		}
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		for range 1000 {
			//time.Sleep(time.Millisecond * 5)
			wg.Add(1)
			go func() {
				//time.Sleep(time.Millisecond * 1)
				mu.RLock()
				print(Counter, " >> ")
				mu.RUnlock()
				wg.Done()
			}()
		}
		wg.Done()
	}()

	wg.Wait()
	println(Counter)
}
