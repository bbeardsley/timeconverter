package main

import (
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode"
)

const dateLayout = "2006-01-02"
const hoursMinutesLayout = "15:04"
const secondsLayout = "05"

// Iso8601Replacer parses iso 8601 dates and formats them per format and location
type Iso8601Replacer struct {
	Regex *regexp.Regexp
}

func getTimePortion(dateTimeString string) string {
	dateTimeFields := strings.FieldsFunc(dateTimeString, func(c rune) bool {
		return c == 'T' || c == ' '
	})
	return dateTimeFields[1]
}

func getNanosLength(timePortion string) int {
	lastPeriod := strings.LastIndex(timePortion, ".")
	if lastPeriod < 1 {
		return 0
	}
	length := 0
	for i := lastPeriod + 1; i < len(timePortion); i++ {
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

func getNonUtcTimeLayout(timePortion string, hasT bool, nanosLen int) string {
	hasZ := strings.Contains(timePortion, "Z")
	offsetIndex := strings.LastIndexFunc(timePortion, func(c rune) bool {
		return c == '+' || c == '-'
	})
	timezoneLength := len(timePortion) - offsetIndex - 1

	var sb strings.Builder
	sb.WriteString(dateLayout)
	if hasT {
		sb.WriteString("T")
	} else {
		sb.WriteString(" ")
	}

	sb.WriteString(hoursMinutesLayout)
	hasSeconds := strings.Count(timePortion, ":") >= 2
	if hasSeconds {
		sb.WriteString(":")
		sb.WriteString(secondsLayout)
	}

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
	hasT := strings.Contains(dateString, "T")
	hasZ := strings.Contains(dateString, "Z")
	timePortion := getTimePortion(dateString)
	nanosLen := getNanosLength(timePortion)

	if strings.Contains(timePortion, "+") || strings.Contains(timePortion, "-") {
		return getNonUtcTimeLayout(timePortion, hasT, nanosLen)
	}

	var sb strings.Builder
	sb.WriteString(dateLayout)
	if hasT {
		sb.WriteString("T")
	} else {
		sb.WriteString(" ")
	}

	sb.WriteString(hoursMinutesLayout)
	hasSeconds := strings.Count(timePortion, ":") >= 2
	if hasSeconds {
		sb.WriteString(":")
		sb.WriteString(secondsLayout)
	}

	writeNanos(&sb, nanosLen)

	if hasZ {
		sb.WriteString("Z")
	}

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
		if format == UnixSeconds {
			return strconv.FormatInt(t.Unix(), 10)
		}
		return t.In(location).Format(format)
	})
}

// NewIso8601Replacer creates a new replacer with the regex initialized
func NewIso8601Replacer() *Iso8601Replacer {
	return &Iso8601Replacer{regexp.MustCompile(`\d{4}-\d{2}-\d{2}(T| )\d{2}:\d{2}(:\d{2}(\.\d{1,})?Z?((\+|\-)(\d{4}|\d{2}:\d{2}|\d{2}))?|Z)?`)}
}
