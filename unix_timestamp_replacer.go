package main

import (
	"regexp"
	"strconv"
	"time"
)

// UnixTimestampReplacer parses unix timestamps formats them per format and location
type UnixTimestampReplacer struct {
	Regex *regexp.Regexp
}

// ReplaceDates replaces multiple instances in the input string
func (replacer UnixTimestampReplacer) ReplaceDates(input string, format string, location *time.Location) string {
	return replacer.Regex.ReplaceAllStringFunc(input, func(timestampMatch string) string {
		i, err := strconv.ParseInt(timestampMatch, 10, 64)
		if err != nil {
			panic(err)
		}
		t := time.Unix(i, 0)
		if format == unixSeconds {
			return strconv.FormatInt(t.Unix(), 10)
		}
		return t.In(location).Format(format)
	})
}

// NewUnixTimestampReplacer creates a new replacer with the regex initialized
func NewUnixTimestampReplacer() *UnixTimestampReplacer {
	return &UnixTimestampReplacer{regexp.MustCompile(`\d+`)}
}
