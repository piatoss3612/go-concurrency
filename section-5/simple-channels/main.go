package main

import (
	"fmt"
	"strings"
)

// shout has two paraneters ping and pong
// ping is receive only channel
// pong is send only channel
func shout(ping <-chan string, pong chan<- string) {
	for {
		// goroutine waits here until something is recevied on the channel 'ping'
		s, ok := <-ping
		if !ok {
			// do something
		}

		// convert s to uppercase and send it to channel 'pong'
		pong <- fmt.Sprintf("%s!!!", strings.ToUpper(s))
	}
}

func main() {
	// create two channels
	ping := make(chan string)
	pong := make(chan string)

	// start a goroutine
	go shout(ping, pong)

	fmt.Println("Type something and press ENTER (enter Q to quit)")

	for {
		// print a prompt
		fmt.Print("-> ")

		// get upser input
		var userInput string
		_, _ = fmt.Scanln(&userInput)

		if strings.ToLower(userInput) == "q" {
			break
		}

		ping <- userInput

		// wait for a response
		response := <-pong
		fmt.Println("Response:", response)
	}

	fmt.Println("All done! closing channels...")

	close(ping)
	close(pong)
}
