package main

import (
	"fmt"

	calculatetime "github.com/VagueCoder/Random-Go-Snippets/Fast-Fibonacci/calculateTime"
	"github.com/VagueCoder/Random-Go-Snippets/Fast-Fibonacci/fibonacci"
)

func main() {
	calculatetime.Start()
	defer calculatetime.End()

	val := fibonacci.Fibonacci(50)
	fmt.Printf("Fibonacci(50): %v\n", val)
}
