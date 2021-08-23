package formattedTime

import (
	"fmt"
	"regexp"
	"time"
)

type FormattedTime string

func (f FormattedTime) MarshalJSON() ([]byte, error) {
	// Curtom regex as per time string . Eg.: "2021-05-21 23:15:46.3598134 +0530 IST m=+0.000204301"
	pattern := regexp.MustCompile(`\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}.\d+ [-+]\d{4} [A-Z]{3}`)

	// Extract & validate part of time string to parse
	b := pattern.Find([]byte(fmt.Sprint(f)))
	if len(b) == 0 {
		return []byte(""), fmt.Errorf("error at Regex Match: Couldn't find time string in pattern XXXX-XX-XX XX:XX:XX")
	}

	// Parse time string int time.Time object
	t, err := time.Parse("2006-01-02 15:04:05 -0700 MST", string(b))
	if err != nil {
		return []byte(""), fmt.Errorf("error at FormattedTime Marshal: %v", err)
	}

	// Again, format time.Time object into required time format
	timeString := fmt.Sprintf("%q", t.Format("02-Jan-2006 15:04:05 MST"))

	return []byte(timeString), nil
}
