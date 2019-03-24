package main

import (
	"testing"
	"time"
)

type testDefinition struct {
	Input    string
	Expected string
}

const format = "Mon 2006 Jan 02 03:04pm MST"

func (definition testDefinition) execute(t *testing.T) {
	location, _ := time.LoadLocation("America/Denver")
	replacer := NewIso8601Replacer()

	result := replacer.ReplaceDates(definition.Input, format, location)
	if result != definition.Expected {
		t.Error(
			"For", definition.Input,
			"expected", definition.Expected,
			"got", result,
		)
	}
}

// TestUtcDates should work
func TestUtcDates(t *testing.T) {
	testDefinitions := []testDefinition{
		{"2019-03-17T01:02:03Z", "Sat 2019 Mar 16 07:02pm MDT"},
		{"2019-03-17T01:02:03.004Z", "Sat 2019 Mar 16 07:02pm MDT"},
	}
	for _, definition := range testDefinitions {
		definition.execute(t)
	}
}

func TestUtcPlusOffset(t *testing.T) {
	testDefinitions := []testDefinition{
		{"2019-03-17T01:02:03Z+05", "Sat 2019 Mar 16 02:02pm MDT"},
		{"2019-03-17T01:02:03Z+05:30", "Sat 2019 Mar 16 01:32pm MDT"},
		{"2019-03-17T01:02:03Z+0530", "Sat 2019 Mar 16 01:32pm MDT"},

		{"2019-03-17T01:02:03.004Z+05", "Sat 2019 Mar 16 02:02pm MDT"},
		{"2019-03-17T01:02:03.004Z+05:30", "Sat 2019 Mar 16 01:32pm MDT"},
		{"2019-03-17T01:02:03.004Z+0530", "Sat 2019 Mar 16 01:32pm MDT"},

		{"2019-03-17T01:02:03+05", "Sat 2019 Mar 16 02:02pm MDT"},
		{"2019-03-17T01:02:03+0530", "Sat 2019 Mar 16 01:32pm MDT"},
		{"2019-03-17T01:02:03+05:30", "Sat 2019 Mar 16 01:32pm MDT"},

		{"2019-03-17T01:02:03.004+05", "Sat 2019 Mar 16 02:02pm MDT"},
		{"2019-03-17T01:02:03.004+0530", "Sat 2019 Mar 16 01:32pm MDT"},
		{"2019-03-17T01:02:03.004+05:30", "Sat 2019 Mar 16 01:32pm MDT"},
	}
	for _, definition := range testDefinitions {
		definition.execute(t)
	}
}

func TestUtcMinusOffset(t *testing.T) {
	testDefinitions := []testDefinition{
		{"2019-03-17T01:02:03Z-05", "Sun 2019 Mar 17 12:02am MDT"},
		{"2019-03-17T01:02:03Z-05:30", "Sun 2019 Mar 17 12:32am MDT"},
		{"2019-03-17T01:02:03Z-0530", "Sun 2019 Mar 17 12:32am MDT"},

		{"2019-03-17T01:02:03.004Z-05", "Sun 2019 Mar 17 12:02am MDT"},
		{"2019-03-17T01:02:03.004Z-05:30", "Sun 2019 Mar 17 12:32am MDT"},
		{"2019-03-17T01:02:03.004Z-0530", "Sun 2019 Mar 17 12:32am MDT"},

		{"2019-03-17T01:02:03-05", "Sun 2019 Mar 17 12:02am MDT"},
		{"2019-03-17T01:02:03-05:30", "Sun 2019 Mar 17 12:32am MDT"},
		{"2019-03-17T01:02:03-0530", "Sun 2019 Mar 17 12:32am MDT"},

		{"2019-03-17T01:02:03.004-05", "Sun 2019 Mar 17 12:02am MDT"},
		{"2019-03-17T01:02:03.004-05:30", "Sun 2019 Mar 17 12:32am MDT"},
		{"2019-03-17T01:02:03.004-0530", "Sun 2019 Mar 17 12:32am MDT"},
	}
	for _, definition := range testDefinitions {
		definition.execute(t)
	}
}

func TestMultipleDates(t *testing.T) {
	testDefinitions := []testDefinition{
		{"Date 1: 2019-03-17T01:02:03Z-05 and Date 2: 2019-03-17T01:02:03Z-05.", "Date 1: Sun 2019 Mar 17 12:02am MDT and Date 2: Sun 2019 Mar 17 12:02am MDT."},
	}
	for _, definition := range testDefinitions {
		definition.execute(t)
	}
}
func TestNonRecognized(t *testing.T) {
	testDefinitions := []testDefinition{
		{"This is a test string with no valid dates", "This is a test string with no valid dates"},
		{"This is a test string with an invalid date 2019-03-17T01:02:03.", "This is a test string with an invalid date 2019-03-17T01:02:03."},
	}
	for _, definition := range testDefinitions {
		definition.execute(t)
	}
}
