package main

import (
	"fmt"
	"sync"
)

var (
	msg string
	wg  sync.WaitGroup
)

func updateMessage(s string) {
	defer wg.Done()

	msg = s
}

func main() {
	msg = "Hello, World!"

	wg.Add(2)

	go updateMessage("Hello, Universe!")
	go updateMessage("Hello, Cosmos!")

	wg.Wait()

	fmt.Println(msg)
}

// func updateMessage(s string, m *sync.Mutex) {
// 	defer wg.Done()

// 	m.Lock()
// 	msg = s
// 	m.Unlock()
// }

// func main() {
// 	msg = "Hello, World!"

// 	var mutex sync.Mutex

// 	wg.Add(2)

// 	go updateMessage("Hello, Universe!", &mutex)
// 	go updateMessage("Hello, Cosmos!", &mutex)

// 	wg.Wait()

// 	fmt.Println(msg)
// }
