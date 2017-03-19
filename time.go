package bunq

import (
	"fmt"
	"strings"
	"time"
)

// Time represents a parsed timestamp returned by the bunq API.
type Time time.Time

// UnmarshalJSON is being used to parse timestamps from the bunq API.
// Bunq timestamps come into this format: "2015-06-13 23:19:16.215235"
func (ariTime *Time) UnmarshalJSON(input []byte) error {
	strInput := string(input)
	strInput = strings.Trim(strInput, `"`)
	var pattern string
	if len(strInput) == 10 {
		pattern = "2006-01-02"
	} else {
		pattern = "2006-01-02 15:04:05.000000"
	}
	newTime, err := time.Parse(pattern, strInput)
	if err != nil {
		return fmt.Errorf("Error parsing Time: %s", err)
	}
	*ariTime = Time(newTime)

	return nil
}

// MarshalText implements the encoding.MarshalText interface for custom time type.
func (ariTime *Time) MarshalText() ([]byte, error) {
	t := time.Time(*ariTime)
	return []byte(t.Format("2006-01-02 15:04:05.000000")), nil
}
