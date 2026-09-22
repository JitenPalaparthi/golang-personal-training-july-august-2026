package main

func main() {
	defer println("Done with main")

	ch := make(chan int, 10)
	//sig := make(chan bool)
	sig := make(chan struct{})

	go sender(ch, 100)
	go receiver(ch, sig)

	<-sig
	//println(v)
}

func sender(ch chan<- int, r int) {
	for i := range r {
		ch <- (i + 1) * (i + 1)
	}
	close(ch)
}

func receiver(ch <-chan int, sig chan<- struct{}) {
	for v := range ch {
		println(v)
	}
	sig <- struct{}{}
}
