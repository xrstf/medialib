package convert

import (
	"fmt"
	"testing"
)

func TestSecondsToTimestamp(t *testing.T) {
	testcases := []struct {
		seconds  float64
		expected string
	}{
		{
			seconds:  0,
			expected: "00:00:00:000",
		},
		{
			seconds:  120,
			expected: "00:02:00:000",
		},
		{
			seconds:  254.954613,
			expected: "00:04:14:954",
		},
	}

	for _, testcase := range testcases {
		t.Run(fmt.Sprintf("%.5f", testcase.seconds), func(t *testing.T) {
			converted := secondsToTimestamp(testcase.seconds)

			if converted != testcase.expected {
				t.Fatalf("Expected %q, got %q.", testcase.expected, converted)
			}
		})
	}
}
