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
		result := CheckLengthRange(e.str, e.minLength, e.maxLength)
		if e.expected && !result {
			t.Errorf("%s: expected true but got %v", e.testName, result)
		}
	}
}

func Test_HasWhiteSpaces(t *testing.T) {
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
		result := HasWhiteSpaces(e.str)
		if e.expected && !result {
			t.Errorf("%s: expected true but got %v", e.testName, result)
		}
	}
}

func Test_IsBlankField(t *testing.T) {
	blankFieldTests := []struct {
		testName string
		str string
		expected bool
	} {
		{"Empty String", "", true},
		{"String with blank spaces but no words", "   ", false},
		{"String with blank spaces", "  Hello   ", false},
		{"Regular string with no blank spaces", "Hello world", false},
	}

	for _, e := range blankFieldTests {
		result := IsBlankField(e.str)
		if e.expected && !result {
			t.Errorf("%s: expected true but got %v", e.testName, result)
		}
	}
}

func Test_CheckUsernameSpecialChar(t *testing.T) {
	usernameSpecialCharTests := []struct {
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

	for _, e := range usernameSpecialCharTests {
		result := CheckUsernameSpecialChar(e.str)
		if e.expected && !result {
			t.Errorf("%s: expected true but got %v", e.testName, result)
		}
	}
}

func Test_CheckPasswordSpecialChar(t *testing.T) {
	passwordSpecialCharTests := []struct {
		testName string
		str string
		expected bool 
	} {
		{"Contains lowercase, uppercase, number, and special character", "p@ssw0rdTest", true},
		{"Contains lowercase and uppercase letters only", "passwordTest", false},
		{"Contains lowercase letters and numbers only", "password123", false},
		{"Contains uppercase letters and special characters only", "PASSWORD!@#", false},
		{"Missing special character", "Password123", false},
	}

	for _, e := range passwordSpecialCharTests {
		result := CheckPasswordSpecialChar(e.str)
		if e.expected != result {
			t.Errorf("%s: expected %v but got %v", e.testName, e.expected, result)
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