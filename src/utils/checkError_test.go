package utils

import (
	"bytes"
	"errors"
	"log"
	"os"
	"testing"
	"time"
)

func Test_CheckError(t *testing.T) {
	// redirect log output to testing.T
	var buf bytes.Buffer
	log.SetOutput(&buf)

	defer func() {
		log.SetOutput(os.Stderr)
	} ()

	// Generate an error
	err := errors.New("test error")

	CheckError(err)
	actualOutput := buf.String()

	// Get date time to compare with logs
	dt := time.Now()
	dtString := dt.Format("2006/01/02 15:04:05")
	expectedOutput := dtString + " Internal Server Error: test error\n"
	
	if actualOutput != expectedOutput {
		t.Errorf("Incorrect CheckError: expected %q but got %q", expectedOutput, actualOutput)
	}
}

func Test_CheckErrorDatabase(t *testing.T) {
	// redirect log output to testing.T
	var buf bytes.Buffer // creates an in-memory buffer to write data to the buffer, can process the data later
	log.SetOutput(&buf) // redirect the log output to the buffer

	defer func() {
		// reset the output destination of the log package to the default output stream (os.Stderr)
		log.SetOutput(os.Stderr) // os.Stderr: standard error output stream
	} ()

	// Generate an error
	err := errors.New("test error database")

	CheckErrorDatabase(err)
	actualOutput := buf.String()

	// Get date time to compare with logs
	dt := time.Now()
	dtString := dt.Format("2006/01/02 15:04:05")
	expectedOutput := dtString + " PostgreSQL Internal Server Error: test error database\n"
	
	if actualOutput != expectedOutput {
		t.Errorf("Incorrect CheckError: expected %q but got %q", expectedOutput, actualOutput)
	}
}
