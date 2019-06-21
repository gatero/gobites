package main

import (
	"testing"
)

func TestAssign(t *testing.T) {
	firstMap := map[string]interface{}{
		"one": "First value",
		"two": "Second value",
	}

	secondMap := map[string]interface{}{
		"one":   "overwrited",
		"three": "third value",
		"four":  "fourth value",
	}

	mergedMaps := Assign(firstMap, secondMap)

	if !(len(mergedMaps) > 0) {
		t.Errorf("The expected length id 4 and got %v", len(mergedMaps))
	}

	expectedKeys := map[string]interface{}{
		"one":   "overwrited",
		"two":   "Second value",
		"three": "third value",
		"four":  "fourth value",
	}
	for key, value := range mergedMaps {
		if expectedKeys[key] != nil {
			if expectedKeys[key] != value {
				t.Errorf("The extected value was %v and got %v", expectedKeys[key], value)
			}
		} else {
			t.Errorf("The extected value was %v and got %v", key, expectedKeys[key])
		}
	}
}
