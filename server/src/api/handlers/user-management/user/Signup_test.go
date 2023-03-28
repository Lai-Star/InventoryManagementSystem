package user

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/LeonLow97/inventory-management-system-golang-react-postgresql/database/repository/dbrepo"
	"github.com/LeonLow97/inventory-management-system-golang-react-postgresql/types"
	"github.com/LeonLow97/inventory-management-system-golang-react-postgresql/utils"
)

func Test_SignUpTransaction(t *testing.T) {

	app := application{}
	app.DB = &dbrepo.TestDBRepo{}

	var tests = []struct {
		name               string
		postedData         types.SignUpJSON
		expectedBody       string
		expectedStatusCode int
	}{
		{
			name: "valid signup",
			postedData: types.SignUpJSON{
				Username: "leonlow",
				Password: "Password0!",
				Email:    "leonlow@email.com",
			},
			expectedBody:       `{"Success":"User leonlow was successfully created!","Status":201}`,
			expectedStatusCode: 201,
		},
		{
			name: "invalid signup",
			postedData: types.SignUpJSON{
				Username: "a",
				Password: "Password",
				Email:    "a@email.com",
			},
			expectedBody:       `{"Err":"Username must have a length of 5 - 50 characters. Please try again.","Status":400}`,
			expectedStatusCode: 400,
		},
	}

	for _, e := range tests {
		jsonStr, err := json.Marshal(e.postedData)
		if err != nil {
			t.Fatal(err)
		}

		// Setting a request for testing
		reqBody := bytes.NewBuffer(jsonStr)
		req, _ := http.NewRequest(http.MethodPost, "/signup", reqBody)
		req.Header.Set("Content-Type", "application/json")

		// Setting and recording the response
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(utils.MakeHTTPHandler(app.SignUp))

		handler.ServeHTTP(rr, req)
		// SignUp(rr, req)

		if rr.Code != e.expectedStatusCode {
			t.Errorf("%s: returned wrong status code; expected %d but got %d", e.name, e.expectedStatusCode, rr.Code)
		}

		if strings.TrimSpace(rr.Body.String()) != e.expectedBody {
			t.Errorf("Unexpected response body: expected %v, got %v", rr.Body.String(), e.expectedBody)
		}
	}
}
