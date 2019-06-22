package main

import (
	"testing"
)

var containsTest = []struct {
	query    string
	expected bool
}{
	{"one", true},
	{"four", false},
}

func TestContains(t *testing.T) {
	slice := []string{"one", "two", "three"}

	for _, tt := range containsTest {
		result := Contains(slice, tt.query)

		if result != tt.expected {
			t.Errorf("The expected value was %v and get %v", true, result)
		}
	}
}
