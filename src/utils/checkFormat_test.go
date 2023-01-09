package utils

import "testing"

func Test_Length(t *testing.T) {
	lengthTests := []struct {
		testName string
		str string
		minLength int
		maxLength int
		expected bool
	} {
		{"More than maxLength", "TestPassword123", 8, 12, false},
		{"Less than minLength", "", 2, 10, false},
		{"Correct Length", "CorrectPass", 8, 12, true},
		{"Test with special characters", "#$@!123FAe", 8, 20, true},
	}

	for _, e := range lengthTests {
		result := CheckLength(e.str, e.minLength, e.maxLength)
		if e.expected && !result {
			t.Errorf("%s: expected true but got %v", e.testName, result)
		}
	}
}

func Test_CheckWhiteSpaces(t *testing.T) {
	
}