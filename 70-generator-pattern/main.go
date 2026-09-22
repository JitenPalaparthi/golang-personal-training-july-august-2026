package main

import "fmt"

func main() {
	defer println("Done with main")
	//sig := make(chan bool)

	ch1 := Generate("gen-1", 50)
	ch2 := Generate("gen-2", 50)

	sig1 := Receive(ch1)
	sig2 := Receive(ch2)

	<-sig1
	<-sig2
	//println(v)
}

func Generate(name string, r uint) <-chan string { // This is called generator pattern
	ch := make(chan string)
	go func() {
		for i := range r {
			ch <- fmt.Sprint(name, "--->", int(i+1)*int(i+1))
		}
		close(ch)
	}()

	return ch
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
