package user

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func Test_SignUpTransaction(t *testing.T) {
	// Set up mock database
	// mockDB := &database.MockDBRepo{}
	// database := mockDB

	var tests = []struct {
		name               string
		postedData         SignUpJson
		expectedStatusCode int
	}{
		{
			name: "valid login",
			postedData: SignUpJson{
				Username: "leonlow",
				Password: "Password0!",
				Email:    "leonlow@email.com",
			},
			expectedStatusCode: 200,
		},
	}

	for _, e := range tests {
		jsonStr, err := json.Marshal(e.postedData)
		if err != nil {
			t.Fatal(err)
		}

		// Setting a request for testing
		req, _ := http.NewRequest(http.MethodPost, "/signup", strings.NewReader(string(jsonStr)))
		req.Header.Set("Content-Type", "application/json")

		// Setting and recording the response
		res := httptest.NewRecorder()
		handler := http.HandlerFunc(SignUp)

		handler.ServeHTTP(res, req)
		// SignUp(res, req)

		if res.Code != e.expectedStatusCode {
			t.Errorf("%s: returned wrong status code; expected %d but got %d", e.name, e.expectedStatusCode, res.Code)
		}
	}
}
