package randomStrings

import (
	"math/rand"
)

// randInt is local function that returns a random integer between provided min, max.
func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

// RandomString returns random string with the requested size.
func RandomString(size int) string {
	bytes := make([]byte, size)
	for i := 0; i < size; i++ {
		bytes[i] = byte(randInt(65, 90))
	}
	return string(bytes)
}
