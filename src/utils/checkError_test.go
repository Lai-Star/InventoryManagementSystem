package utils

import (
	"errors"
	"fmt"
	"io"
	"os"
	"testing"
)

func Test_CheckError(t *testing.T) {
	// Generate an error
	err := errors.New("test error")

	// Capture the output of CheckError
	out := captureOutput(func() {
		CheckError(err)
	})

	// Compare the output to the expected output
	expectedOutput := fmt.Sprintf("Internal Server Error: %s\n", err)
	if out != expectedOutput {
		t.Errorf("Incorrect CheckError: expected %q but got %q", expectedOutput, out)
	}
}

func Test_CheckErrorDatabase(t *testing.T) {
	// Generate an error
	err := errors.New("test error database")

	// Capture the output of CheckErrorDatabase
	out := captureOutput(func() {
		CheckErrorDatabase(err)
	})

	// Compare the output to the expected output
	expectedOutput := fmt.Sprintf("PostgreSQL Internal Server Error: %s\n", err)
	if out != expectedOutput {
		t.Errorf("Incorrect CheckError: expected %q but got %q", expectedOutput, out)
	}
}

func captureOutput(f func()) string {
	// Save a copy of os.Stdout
	oldOut := os.Stdout

	// Create a read and write pipe
	r, w, _ := os.Pipe()

	// Set os.Stdout to the write pipe
	os.Stdout = w

	// Call the function that is being tested
	f()

	// Close the writer (prevent resource leak)
	_ = w.Close()

	// Reset os.Stdout to what it was before
	os.Stdout = oldOut

	// Read the output from the read pipe
	out, _ := io.ReadAll(r)

	return string(out)
}
