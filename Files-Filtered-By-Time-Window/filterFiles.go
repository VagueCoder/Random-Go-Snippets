package filterfiles

import (
	"fmt"
	"io/ioutil"
	"time"
)

func FilesBetweenTime(path, start, end string) ([]string, error) {
	layout := "02-Jan-2006 15:04:05 MST"
	startTime, err := time.Parse(layout, start)
	if err != nil {
		return nil, fmt.Errorf("Error at parsing start time: %v\nNote: Format the date similar to 02-Jan-2006 15:04:05 MST", err)
	}
	endTime, err := time.Parse(layout, end)
	if err != nil {
		return nil, fmt.Errorf("Error at parsing end time: %v\nNote: Format the date similar to 02-Jan-2006 15:04:05 MST", err)
	}

	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, fmt.Errorf("Error at reading directory: %v\n", err)
	}

	var filtered []string
	for _, file := range files {
		modifiedTime := file.ModTime()
		if modifiedTime.After(startTime) && modifiedTime.Before(endTime) {
			filtered = append(filtered, file.Name())
		}
	}

	return filtered, nil
}
