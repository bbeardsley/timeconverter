package timeconverter

import (
	"testing"
	"time"
)

type unixTimestampTestDefinition struct {
	Input    string
	Expected string
}

func (definition unixTimestampTestDefinition) execute(t *testing.T) {
	location, _ := time.LoadLocation("America/Denver")
	replacer := NewUnixTimestampReplacer()

	result := replacer.ReplaceDates(definition.Input, format, location)
	if result != definition.Expected {
		t.Error(
			"For", "\""+definition.Input+"\"",
			"expected", "\""+definition.Expected+"\"",
			"got", "\""+result+"\"",
		)
	}
}

// TestUnixTimestamps should work
func TestUnixTimestamps(t *testing.T) {
	testDefinitions := []unixTimestampTestDefinition{
		{"1405544146", "Wed 2014 Jul 16 02:55pm MDT"},
		{"The time is 1405544146 and the time will never be 1405544146 again", "The time is Wed 2014 Jul 16 02:55pm MDT and the time will never be Wed 2014 Jul 16 02:55pm MDT again"},
		{"Time marches on", "Time marches on"},
	}
	for _, definition := range testDefinitions {
		definition.execute(t)
	}
}
