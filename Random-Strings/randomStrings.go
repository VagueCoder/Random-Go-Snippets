package randomStrings

import (
	"context"
	"math/rand"
	"time"
)

// randInt is local function that returns a random integer between provided min, max.
func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

// randomString is local function that processes random strings \
// continuously until cancel function is called.
func randomString(ctx context.Context, ch chan<- string, size int) {
	for {
		bytes := make([]byte, size)
		for i := 0; i < size; i++ {
			bytes[i] = byte(randInt(65, 90))
		}
		ch <- string(bytes)
	}
}

// RandomStrings initiates the randomString process \
// returns channel to read from, and cancel function to stop manually.
func RandomStrings(size int) (chan string, context.CancelFunc) {
	ctx, cancel := context.WithCancel(context.Background())

	// Channel to exchange strings between randomString goroutine and caller. \
	// Buffered channel's size can be modified as per need. Here, 100.
	ch := make(chan string, 100)

	// Seed so the strings are created random
	rand.Seed(time.Now().UnixNano())

	// Initiate the randomString process in concurrent goroutine
	go randomString(ctx, ch, size)

	return ch, cancel
}
