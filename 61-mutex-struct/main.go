package main

import (
	"sync"
)

var (
	Wg = new(sync.WaitGroup)
)

// ThreadSafe type
type CounterType struct {
	Counter int
	Mu      *sync.RWMutex
	// Wg      *sync.WaitGroup
}

func New(c int) *CounterType {
	return &CounterType{Counter: c, Mu: new(sync.RWMutex{})}
}

func (ct *CounterType) Increment(r int) {
	//ct.Wg.Add(1)
	//go func() {
	for range r {
		ct.Mu.Lock()
		ct.Counter++
		ct.Mu.Unlock()
	}
	//}()
	//ct.Wg.Done()
}

func (ct *CounterType) Decrement(r int) {
	//ct.Wg.Add(1)
	for range r {
		ct.Mu.Lock()
		ct.Counter--
		ct.Mu.Unlock()
	}
	//ct.Wg.Done()
}

func (ct *CounterType) Print() {
	//ct.Wg.Add(2)
	//ct.Mu.RLock()
	println(ct.Counter)
	//ct.Mu.RUnlock()
	//ct.Wg.Done()
}

func main() {
	ct := New(10)
	Wg.Add(2)
	go func() {
		ct.Increment(1000)
		Wg.Done()
	}()
	go func() {
		ct.Decrement(100)
		Wg.Done()
	}()

	Wg.Wait()
	ct.Print()
}
