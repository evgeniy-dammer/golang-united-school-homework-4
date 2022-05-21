package string_sum

import (
	"testing"
)

type testStringSum struct {
	name  string
	input string
	want  string
}

var testCases = []testStringSum{
	{"one", "-3+5", "2"},
	{"two", "3+5", "8"},
	{"three", " 5 + - 3 ", "2"},
	{"four", " -20 - 50 ", "-70"},
}

func TestStringSumMy(t *testing.T) {
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			actual, err := StringSum(testCase.input)

			if err != nil {
				t.Errorf("%s", err)
			}

			if actual != testCase.want {
				t.Errorf("got: %s; want: %s", actual, testCase.want)
			}
		})
	}
}
