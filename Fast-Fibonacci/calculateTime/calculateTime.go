package calculatetime

import (
	"fmt"
	"time"
)

var (
	start time.Time
	end   time.Duration
)

func Start() {
	start = time.Now()
}

func End() {
	end = time.Duration(time.Duration(time.Since(start)))
	fmt.Printf("\nTotal time of execution: %v\n", end)
}
