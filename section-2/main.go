package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	words := []string{"alpha", "beta", "delta", "pi", "zeta", "eta", "theta", "epsilon"}

	wg.Add(12)

	for i, word := range words {
		go printSomething(fmt.Sprintf("%d: %s", i, word), &wg)
	}

	wg.Wait()

	// wg.Add(1)

	printSomething("This is the second thing to be printed!", &wg)
}

func printSomething(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(s)
}
