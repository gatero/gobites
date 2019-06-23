package main

import (
	"errors"
	"testing"
)

var findTests = []struct {
	query    string
	expected int
	err      error
}{
	{"one", 0, nil},
	{"two", 1, nil},
	{"three", 2, nil},
	{"four", 3, nil},
	{"five", 0, errors.New("'five' was not found")},
}

func TestFunc(t *testing.T) {
	slice := []string{"one", "two", "three", "four"}

	for _, tt := range findTests {
		result, err := Find(slice, tt.query)

		if result != tt.expected {
			t.Errorf("The expected value was \"%v\" and get \"%v\"", tt.expected, result)
		}

		if err != nil && err.Error() != tt.err.Error() {
			t.Errorf("The expected value was \"%v\" and get \"%v\"", tt.err, err)
		}
	}
}
