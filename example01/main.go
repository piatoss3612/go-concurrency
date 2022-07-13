package main

import (
	"fmt"
	"time"
)

func main() {
	words := []string{"alpha", "beta", "delta", "pi", "zeta", "eta", "theta", "epsilon"}

	for i, word := range words {
		go printSomething(fmt.Sprintf("%d: %s", i, word))
	}

	time.Sleep(10 * time.Nanosecond)

	printSomething("This is the second thing to be printed!")
}

func printSomething(s string) {
	fmt.Println(s)
}
