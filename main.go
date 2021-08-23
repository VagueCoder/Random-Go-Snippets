package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	calculatetime "github.com/VagueCoder/Random-Go-Snippets/Calculate-Time"
	fibonacci "github.com/VagueCoder/Random-Go-Snippets/Fast-Fibonacci"
	filterfiles "github.com/VagueCoder/Random-Go-Snippets/Files-Filtered-By-Time-Window"
	gcd "github.com/VagueCoder/Random-Go-Snippets/Find-GCD"
	formattedTime "github.com/VagueCoder/Random-Go-Snippets/Formatted-Time-Marshalling"
	hashmapOps "github.com/VagueCoder/Random-Go-Snippets/Go-HashMap-Operations"
	broadcast "github.com/VagueCoder/Random-Go-Snippets/Message-Broadcaster"
	randomStrings "github.com/VagueCoder/Random-Go-Snippets/Random-Strings"
)

func main() {

	// Calculate-Time
	fmt.Println("\n------ Calculate-Time ------")
	calculatetime.Start()
	fmt.Println("Wait time of 1 second")
	time.Sleep(time.Second)
	calculatetime.End()

	// Fast-Fibonacci
	fmt.Println("\n------ Fast-Fibonacci ------")
	calculatetime.Start()
	val := fibonacci.Fibonacci(50)
	fmt.Printf("Fibonacci(50): %v\n", val)
	calculatetime.End()

	// Files-Filtered-By-Time-Window
	fmt.Println("\n------ Files-Filtered-By-Time-Window ------")
	calculatetime.Start()
	// Start and end date should be in the format 02-Jan-2006 15:04:05 MST
	start := "01-Jul-2021 00:00:00 IST"
	end := "31-Jul-2021 00:00:00 IST"
	path := "/home/vaguecoder/Workspace/go/src/github.com/VagueCoder/Random-Go-Snippets/Files-Filtered-By-Time-Window/files"
	files, _ := filterfiles.FilesBetweenTime(path, start, end)
	fmt.Printf("Files in between time windows (%v, %v): %v\n", start, end, files)
	calculatetime.End()

	// Find-GCD
	fmt.Println("\n------ Find-GCD ------")
	calculatetime.Start()
	gcdval := gcd.FindGCD(54, 90, 180, 720, 3600)
	fmt.Printf("GCD(54, 90, 180, 720, 3600) = %v\n", gcdval)
	calculatetime.End()

	// Formatted-Time-Marshalling
	fmt.Println("\n------ Formatted-Time-Marshalling ------")
	calculatetime.Start()
	now := time.Now()
	nowString := fmt.Sprint(now)
	sample := struct {
		UnformattedTime time.Time                   `json:"unformatted_time"`
		FormattedTime   formattedTime.FormattedTime `json:"formatted_time"`
	}{
		UnformattedTime: now,
		FormattedTime:   formattedTime.FormattedTime(nowString),
	}
	// Marshal the struct object to JSON
	json.NewEncoder(os.Stdout).Encode(sample)
	calculatetime.End()

	// Go-HashMap-Operations
	fmt.Println("\n------ Go-HashMap-Operations ------")
	calculatetime.Start()
	hashmapOps.RunHashMapOperationsShort()
	calculatetime.End()

	// Message-Broadcaster
	fmt.Println("\n------ Message-Broadcaster  ------")
	calculatetime.Start()
	var message interface{}
	ch := broadcast.NewBroadcaster(message)
	defer ch.Close()
	message = struct {
		Fruit1 string
		Fruit2 string
	}{"APPLE", "BANANA"}
	ch.UpdateMessage(message)
	broadcast.RunConsumers(ch, 3)()
	calculatetime.End()

	// Random-Strings
	fmt.Println("\n------ Random-Strings ------")
	calculatetime.Start()
	var stringSize, stringsCount int = 20, 10
	rsch, cancel := randomStrings.RandomStrings(stringSize)
	defer cancel()
	var collection []string
	for i := 0; i < stringsCount; i++ {
		collection = append(collection, <-rsch)
	}
	fmt.Printf("Slice of 10 random strings: %v\n", collection)
	calculatetime.End()
}
