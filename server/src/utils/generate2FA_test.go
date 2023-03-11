package utils

import (
	"regexp"
	"testing"
)

func TestGenerate2FA(t *testing.T) {
	// Call the function to generate an OTP
	otp := Generate2FA()

	// Check that the OTP string is of the correct format
	matched, err := regexp.MatchString(`^[A-Za-z]{4}-\d{6}$`, otp)
	if err != nil {
		t.Errorf("Error matching OTP string format: %v", err)
	}

	if !matched {
		t.Errorf("Expected format to be 4 alphabets, followed by a dash and 6 numerical digits but got %s", otp)
	}

	// Check that the OTP string length is correct
	if len(otp) != 11 {
		t.Errorf("Invalid OTP length: expected 11, got %d", len(otp))
	}

	// Print the generated OTP for debugging purposes
	t.Logf("Generated OTP: %s", otp)
}
