package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/VagueCoder/Random-Go-Snippets/Random-Strings/randomStrings"
)

var err error

func main() {
	var count int
	if len(os.Args) > 1 {
		count, err = strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Printf("Argument Error: %v\n", err)
			os.Exit(1)
		}
	} else {
		fmt.Println("Argument Error: Count not provided in argument 1. Running default: count=10.")
		count = 10
	}

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < count; i++ {
		fmt.Println(randomStrings.RandomString(10))
	}
}
