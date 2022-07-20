package main

import "sync"

var philosopher = []string{"Plato", "Socrates", "Aristotle", "Pascal", "Locke"}
var wg sync.WaitGroup

func diningProblem() {
	defer wg.Done()
}

func main() {
	// print intro

	wg.Add(len(philosopher))

	// spawn one goroutine for each philosopher
	for i := 0; i < len(philosopher); i++ {
		// call goroutine
		go diningProblem()
	}

	wg.Wait()
}
