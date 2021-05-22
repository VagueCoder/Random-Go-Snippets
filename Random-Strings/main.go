package main

import (
	"fmt"

	rs "github.com/VagueCoder/Random-Go-Snippets/Random-Strings/randomStrings"
)

var err error

func main() {
	var stringSize, stringsCount int = 20, 10000

	// RandomStrings
	ch, cancel := rs.RandomStrings(stringSize)

	// deferring cancel function to call before end of current scope.
	defer cancel()

	// Reading strings from channel
	for i := 0; i < stringsCount; i++ {
		fmt.Println(<-ch)
	}
}
