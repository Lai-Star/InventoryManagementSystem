package admin

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/LeonLow97/inventory-management-system-golang-react-postgresql/database/repository/dbrepo"
	"github.com/LeonLow97/inventory-management-system-golang-react-postgresql/types"
	"github.com/LeonLow97/inventory-management-system-golang-react-postgresql/utils"
	"github.com/golang-jwt/jwt/v4"
)

func Test_CreateUser(t *testing.T) {

	app := application{}
	app.DB = &dbrepo.TestDBRepo{}

	var tests = []struct {
		name               string
		postedData         types.AdminUserJSON
		expectedBody       string
		expectedStatusCode int
	}{
		{
			name: "valid create user",
			postedData: types.AdminUserJSON{
				Username:         "lowjiewei",
				Password:         "Password0!",
				Email:            "lowjiewei@email.com",
				OrganisationName: "Golang",
				IsActive:         0,
			},
			expectedBody:       `{"Success":"[Admin] Successfully created 'lowjiewei' user!","Status":201}`,
			expectedStatusCode: 201,
		},
		{
			name: "blank fields",
			postedData: types.AdminUserJSON{
				Username:         "",
				Password:         "",
				Email:            "",
				OrganisationName: "",
				IsActive:         0,
			},
			expectedBody:       `{"Err":"Username cannot be blank","Status":400}`,
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
		req, _ := http.NewRequest(http.MethodPost, "/admin/create-user", reqBody)
		req.Header.Set("Content-Type", "application/json")

		// Setting and recording the response
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(utils.MakeHTTPHandler(app.AdminCreateUser))

		// Create a mock request with a cookie that contains the JWT token
		// Create a valid JWT token with an issuer claim
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
			Issuer: "lowjiewei",
		})
		signedToken, _ := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
		cookie := &http.Cookie{
			Name:  "leon-jwt-token",
			Value: signedToken,
		}
		req.AddCookie(cookie)

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
