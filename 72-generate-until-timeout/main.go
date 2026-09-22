package main

import (
	"fmt"
	"time"
)

func main() {
	defer println("Done with main")
	//sig := make(chan bool)

	ch1 := Generate("gen-1", time.Millisecond*1)

	sig1 := Receive(ch1)

	<-sig1
	//println(v)
}

func Generate(name string, duration time.Duration) <-chan string { // This is called generator pattern
	ch := make(chan string)
	go func() {
		timeOut := Timeout(duration) //time.After(duration)
		i := 1
		for {
			select {
			// case ch <- fmt.Sprint(name, "--->", int(i+1)*int(i+1)):
			// 	i++
			case <-timeOut:
				close(ch)
				return
			default:
				ch <- fmt.Sprint(name, "--->", int(i+1)*int(i+1))
				i++
			}
		}
	}()

	return ch
}

func Timeout(duration time.Duration) <-chan struct{} {
	sig := make(chan struct{})
	go func() {
		time.Sleep(duration)
		sig <- struct{}{}
		close(sig)
	}()
	return sig
}

func Receive(ch <-chan string) <-chan struct{} {
	sig := make(chan struct{})
	go func() {
		for v := range ch {
			println(v)
		}
		sig <- struct{}{}
		close(sig)
	}()
	return sig
}
