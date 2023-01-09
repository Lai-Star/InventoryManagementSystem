package utils

import (
	"testing"
)

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
	whiteSpacesTests := []struct {
		testName string
		str string
		expected bool
	} {
		{"Left whitespace", "    Hello", false},
		{"No whitespace", "Daniel", true},
		{"Right whitespace", "Whatever     ", false},
		{"Middle whitespace", "Good Morning", false},
		{"Combination", "   Good Afternoon     ", false},
	}

	for _, e := range whiteSpacesTests {
		result := CheckWhiteSpaces(e.str)
		if e.expected && !result {
			t.Errorf("%s: expected true but got %v", e.testName, result)
		}
	}
}

func Test_CheckSpecialChar(t *testing.T) {
	specialCharTests := []struct {
		testName string
		str string
		expected bool
	} {
		{"only lowercase", "abcdefg!", false},
		{"only uppercase", "ABCDEF!@#G", false},
		{"only numerical", "1234567", true},
		{"only underscore", "abcd_", true},
		{"all except underscore", "abcABC123", true},
		{"correct format", "abcABC1234_", true},
	}

	for _, e := range specialCharTests {
		result := CheckSpecialChar(e.str)
		if e.expected && !result {
			t.Errorf("%s: expected true but got %v", e.testName, result)
		}
	}
}

func Test_CheckEmailAddress(t *testing.T) {
	emailAddressTests := []struct {
		testName string
		email string
		expected bool
	} {
		{"Correct Email Format", "myself@email.com", true},
		{"Correct Email Format", "myself_123_haha@email.com", true},
		{"Missing @ symbol", "nosymbolemail.com", false},
		{"Missing email domain", "nodomain@email", false},
		{"Missing email domain", "nodomain@", false},
	}

	for _, e := range emailAddressTests {
		result := CheckEmailAddress(e.email)
		if e.expected && !result {
			t.Errorf("%s: expected true but got %v", e.testName, result)
		}
	}
}