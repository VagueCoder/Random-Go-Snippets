package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	ft "github.com/VagueCoder/Random-Go-Snippets/Formatted-Time-Marshalling/formattedTime"
)

// Sample structure to demonstrate formatted and unformatted time objevts
type Sample struct {
	Name            string           `json:"name"`
	ID              int              `json:"rank"`
	UnformattedTime time.Time        `json:"unformatted_time"`
	FormattedTime   ft.FormattedTime `json:"formatted_time"`
}

func main() {
	now := time.Now()
	nowString := fmt.Sprint(now)
	sample := &Sample{
		Name:            "Bob",
		ID:              1,
		UnformattedTime: now,
		FormattedTime:   ft.FormattedTime(nowString),
	}

	// Marshal the struct object to JSON
	err := json.NewEncoder(os.Stdout).Encode(sample)
	if err != nil {
		panic(err)
	}
}

/*
Sample Output:
{
	"name": "Bob",
	"rank": 1,
	"unformatted_time": "2021-05-21T23:20:08.4207751+05:30",
	"formatted_time": "21-May-2021 23:20:08 IST"
}
*/
