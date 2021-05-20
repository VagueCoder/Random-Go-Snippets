package main

import (
	"fmt"
	"os"

	filterfiles "github.com/VagueCoder/Random-Go-Snippets/Files-Filtered-By-Time-Window/filterFiles"
)

func main() {
	// Start and end date should be in the format 02-Jan-2006 15:04:05 MST
	start := "21-May-2021 00:30:00 IST"
	end := "21-May-2021 00:40:00 IST"

	if len(os.Args) == 1 {
		fmt.Println("Arguments Error: Path should be provided as first argument.")
		os.Exit(0)
	}

	path := os.Args[1] // Full or relative

	files, err := filterfiles.FilesBetweenTime(path, start, end)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	fmt.Println(files)
}
