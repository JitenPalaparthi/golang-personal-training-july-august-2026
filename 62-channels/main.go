package main

func main() {
	var ch chan int // this is nil bcz the channel is not instantiated
	if ch == nil {
		println("nil channel")
		ch = make(chan int) // make is used to instantiate a slice, map and channel
	} else {
		println("not nil channel")
	}

	ch <- 42      // main blocked here
	println(<-ch) // does not come here becase it has been blocked at the previous line
}

// Chan is a kind of a queue
// Queue is first in first out
// Using channel multiple Goroutines can communicate with each other safely
// Buffered channels and unbuffered channels
// for a channel, there is a sender and there is a receiver, who are senders and receivers they are goroutines
// | | ==
// 		| |
// There is a sender , when it sends the value but the receiver does not receive it, the sender will block
// There is a receiver, when it receives the value but the sender does not send it, the receiver will block
// channel is always one directional, like water channel
// compiler cannot detect the deadlock
// There is nothing called the sender has to start first, the receiver has to start next
