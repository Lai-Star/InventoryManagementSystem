package auth

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

func Test_Login(t *testing.T) {

	app := application{}
	app.DB = &dbrepo.TestDBRepo{}

	var tests = []struct {
		name               string
		postedData         types.LoginJSON
		expectedBody       string
		expectedStatusCode int
	}{
		{
			name: "valid login",
			postedData: types.LoginJSON{
				Username: "lowjiewei",
				Password: "Password0!",
			},
			expectedBody:       `{"Success":"Login Successful!","Status":200}`,
			expectedStatusCode: 200,
		},
		{
			name: "invalid login",
			postedData: types.LoginJSON{
				Username: "leonlow",
				Password: "Password1!",
			},
			expectedBody:       `{"Err":"You have entered an incorrect username and/or password. Please try again.","Status":401}`,
			expectedStatusCode: 401,
		},
		{
			name: "Blank Fields",
			postedData: types.LoginJSON{
				Username: "",
				Password: "",
			},
			expectedBody:       `{"Err":"Please fill in all fields - Username and Password.","Status":400}`,
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
		handler := http.HandlerFunc(utils.MakeHTTPHandler(app.Login))

		handler.ServeHTTP(rr, req)
		// SignUp(rr, req)

		if rr.Code != e.expectedStatusCode {
			t.Errorf("%s: returned wrong status code; expected %d but got %d", e.name, e.expectedStatusCode, rr.Code)
		}

		if strings.TrimSpace(rr.Body.String()) != e.expectedBody {
			t.Errorf("%s: Unexpected response body: expected %v, got %v", e.name, rr.Body.String(), e.expectedBody)
		}

		// check that the cookie was set correctly
		cookies := rr.Result().Cookies()
		if len(cookies) > 0 {
			cookie := rr.Result().Cookies()[0]
			if cookie.Name != "leon-jwt-token" {
				t.Errorf("expected cookie name %q but got %q", "leon-jwt-token", cookie.Name)
			}
			if cookie.Path != "/" {
				t.Errorf("expected cookie path %q but got %q", "/", cookie.Path)
			}
			if cookie.HttpOnly != true {
				t.Errorf("expected HttpOnly cookie but got non-HttpOnly cookie")
			}
		}
	}
}

func Test_RetrieveIssuer(t *testing.T) {

	rr := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)

	// Create a valid JWT token with an issuer claim
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Issuer: "lowjiewei",
	})
	signedToken, _ := token.SignedString([]byte(os.Getenv("SECRET_KEY")))

	// Create a mock request with a cookie that contains the JWT token
	cookie := &http.Cookie{
		Name:  "leon-jwt-token",
		Value: signedToken,
	}
	req.AddCookie(cookie)

	// Call the function being tested
	RetrieveIssuer(rr, req)

	// Check that the function set the username header correctly
	if rr.Header().Get("username") != "lowjiewei" {
		t.Errorf("Expected username header to be set to %q, but got %q", "lowjiewei", rr.Header().Get("username"))
	}

}
