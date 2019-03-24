package main

import (
	"regexp"
	"strings"
	"time"
	"unicode"
)

const dateTimeLayout = "2006-01-02T15:04:05"

// Iso8601Replacer parses iso 8601 dates and formats them per format and location
type Iso8601Replacer struct {
	Regex *regexp.Regexp
}

func getTimePortion(dateTimeString string) string {
	dateTimeFields := strings.FieldsFunc(dateTimeString, func(c rune) bool {
		return c == 'T'
	})
	return dateTimeFields[1]
}

func getNanosLength(timePortion string) int {
	lastPeriod := strings.LastIndex(timePortion, ".")
	if lastPeriod < 1 {
		return 0
	}
	length := 0
	for i := lastPeriod + 1; i <= len(timePortion); i++ {
		if unicode.IsDigit(rune(timePortion[i])) {
			length++
		} else {
			break
		}
	}
	return length
}

func writeNanos(sb *strings.Builder, nanosLen int) {
	if nanosLen > 0 {
		sb.WriteString(".")
		for i := 1; i <= nanosLen; i++ {
			sb.WriteString("0")
		}
	}
}

func getNonUtcTimeLayout(timePortion string, nanosLen int) string {
	hasZ := strings.Contains(timePortion, "Z")
	offsetIndex := strings.LastIndexFunc(timePortion, func(c rune) bool {
		return c == '+' || c == '-'
	})
	timezoneLength := len(timePortion) - offsetIndex - 1

	var sb strings.Builder
	sb.WriteString(dateTimeLayout)

	writeNanos(&sb, nanosLen)

	if hasZ {
		sb.WriteString("Z")
	}
	switch timezoneLength {
	case 2:
		sb.WriteString("-07")
	case 4:
		sb.WriteString("-0700")
	case 5:
		sb.WriteString("-07:00")
	default:
		panic("Unsupported non UTC time layout " + timePortion)
	}
	return sb.String()
}

func getTimeLayout(dateString string) string {
	timePortion := getTimePortion(dateString)
	nanosLen := getNanosLength(timePortion)

	if strings.Contains(timePortion, "+") || strings.Contains(timePortion, "-") {
		return getNonUtcTimeLayout(timePortion, nanosLen)
	}

	var sb strings.Builder
	sb.WriteString(dateTimeLayout)

	writeNanos(&sb, nanosLen)

	sb.WriteString("Z")

	return sb.String()
}

// ReplaceDates replaces multiple instances in the input string
func (replacer Iso8601Replacer) ReplaceDates(input string, format string, location *time.Location) string {
	return replacer.Regex.ReplaceAllStringFunc(input, func(dateString string) string {
		layout := getTimeLayout(dateString)
		t, err := time.Parse(layout, dateString)
		if err != nil {
			panic(err.Error())
		}
		return t.In(location).Format(format)
	})
}

// NewIso8601Replacer creates a new replacer with the regex initialized
func NewIso8601Replacer() *Iso8601Replacer {
	return &Iso8601Replacer{regexp.MustCompile(`\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(\.\d{1,3})?(Z?(\+|\-)(\d{4}|\d{2}:\d{2}|\d{2})|Z)`)}
}
