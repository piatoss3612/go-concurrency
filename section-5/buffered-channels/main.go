package main

import (
	"fmt"
	"time"
)

func listenToChan(ch chan int) {
	for {
		// print a received data message
		i := <-ch
		fmt.Println("Received", i, "from channel")

		// simulate doing a lot of work
		time.Sleep(1 * time.Second)
	}
}

func main() {
	// ch := make(chan int) // receive one int at a time
	ch := make(chan int, 10) // buffered channel: receive ten int at a time

	go listenToChan(ch)

	for i := 0; i <= 100; i++ {
		fmt.Println("sending", i, "to channel...")
		ch <- i
		fmt.Println("sent", i, "to channel!")
	}

	fmt.Println("Done!")
	close(ch)
}
