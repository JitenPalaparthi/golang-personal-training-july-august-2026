package main

func main() {
	defer println("Done with main")

	ch := make(chan int)
	//sig := make(chan bool)
	sig := make(chan struct{})

	r := 100

	go func(r int) {
		for i := range 100 {
			ch <- (i + 1) * (i + 1)
		}
	}(r)

	go func(r int) {
		for range 100 {
			println(<-ch)
		}
		//	sig <- true
		sig <- struct{}{} // it is only a signal, not a value
	}(r)

	<-sig
	//println(v)
}
