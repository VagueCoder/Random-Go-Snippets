package fibonacci

import (
	"sync"
)

func Fibonacci(n int) int {

	var wg sync.WaitGroup
	var mu sync.Mutex
	var a, b int = 0, 1

	wg.Add(n)

	for i := 0; i < n; i++ {
		go func() {
			mu.Lock()
			defer mu.Unlock()
			b += a
			a = b - a
			wg.Done()
		}()
	}

	wg.Wait()

	return a
}
