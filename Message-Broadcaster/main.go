package main

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/VagueCoder/Random-Go-Snippets/Message-Broadcaster/broadcast"
)

// Global declaration
var (
	wg      *sync.WaitGroup = &sync.WaitGroup{}
	message interface{}
	count   int = 10
	ch      broadcast.Channel
)

func main() {
	// Message-1: String
	message = "Hi"
	ch = broadcast.NewBroadcaster(message)
	defer ch.Close()
	runConsumers()

	// Message-2: Struct
	message = struct {
		Fruit1 string
		Fruit2 string
	}{"APPLE", "BANANA"}
	ch.UpdateMessage(message)
	runConsumers()

	// Message-3: Pointer
	message = &http.Request{}
	ch.UpdateMessage(message)
	runConsumers()
}

// runConsumers runs given number of consumer goroutines to read messages.
func runConsumers() {
	for i := 0; i < count; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup, i int) {
			fmt.Printf("Message From Goroutine-%d: %v\n", i, ch.Read())
			wg.Done()
		}(wg, i)
	}
	wg.Wait()
}
