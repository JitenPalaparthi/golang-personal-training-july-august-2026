package main

func main() {
	defer println("Done with main")

	ch := make(chan int, 10)
	//sig := make(chan bool)
	sig := make(chan struct{})

	r := 100

	go func(r int) {
		for i := range 100 {
			ch <- (i + 1) * (i + 1)
		}
		close(ch)
	}(r)

	go func() {
		for v := range ch { // it iterates over the channel until it is closed
			println(v)
		}
		//	sig <- true
		sig <- struct{}{} // it is only a signal, not a value
	}()

	<-sig
	//println(v)
}
