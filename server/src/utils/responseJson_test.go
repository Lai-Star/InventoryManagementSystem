package utils

import (
	"bytes"
	"errors"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"
)

func Test_ResponseJson(t *testing.T) {
	responseJsonTests := []struct {
		testName string
		code int
		message string
		expectedJsonResponse string
	} {
		{"Successful Login", http.StatusOK, "Successfully Logged In!", `{"code":200,"message":"Successfully Logged In!"}`},
		{"Duplicate username", http.StatusBadRequest, "Username already exists. Please try again.", `{"code":400,"message":"Username already exists. Please try again."}`},
	}

	for _, e := range responseJsonTests {
		// Creating a mock ResponseWriter
		w := httptest.NewRecorder()

		ResponseJson(w, e.code, e.message)

		// Read the response body as a string
		body, _ := io.ReadAll(w.Result().Body)
		actual := string(body)

		expected := e.expectedJsonResponse
		if actual != expected {
			t.Errorf("%s: expected %s but got %s", e.testName, e.expectedJsonResponse, actual)
		}
	}
}

func Test_InternalServerError(t *testing.T) {
	// redirect log output to testing.T
	var buf bytes.Buffer
	log.SetOutput(&buf)

	defer func() {
		log.SetOutput(os.Stderr)
	} ()

	// Generate an error
	err := errors.New("test error")

	w := httptest.NewRecorder()

	// Calling the function
	InternalServerError(w, "Internal Server Error (testing):", err)
	actualOutput := buf.String()

	// Get date time to compare with logs
	dt := time.Now()
	dtString := dt.Format("2006/01/02 15:04:05")
	expectedOutput := dtString + " Internal Server Error (testing): test error\n"
	
	if actualOutput != expectedOutput {
		t.Errorf("Incorrect Internal Server Error: expected %q but got %q", expectedOutput, actualOutput)
	}
}