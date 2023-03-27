package user

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/LeonLow97/inventory-management-system-golang-react-postgresql/types"
	"github.com/LeonLow97/inventory-management-system-golang-react-postgresql/utils"
)

var app application

func Test_SignUpTransaction(t *testing.T) {

	var tests = []struct {
		name               string
		postedData         types.SignUpJSON
		expectedStatusCode int
	}{
		{
			name: "valid signup",
			postedData: types.SignUpJSON{
				Username: "leonlow",
				Password: "Password0!",
				Email:    "leonlow@email.com",
			},
			expectedStatusCode: 201,
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
		handler := http.HandlerFunc(utils.MakeHTTPHandler(app.SignUp))

		handler.ServeHTTP(res, req)
		// SignUp(res, req)

		if res.Code != e.expectedStatusCode {
			t.Errorf("%s: returned wrong status code; expected %d but got %d", e.name, e.expectedStatusCode, res.Code)
		}
	}
}
