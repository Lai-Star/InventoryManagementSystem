package admin

import (
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/LeonLow97/inventory-management-system-golang-react-postgresql/database/repository/dbrepo"
	"github.com/LeonLow97/inventory-management-system-golang-react-postgresql/utils"
	"github.com/golang-jwt/jwt/v4"
)

func Test_GetUsers(t *testing.T) {

	app := application{}
	app.DB = &dbrepo.TestDBRepo{}

	// Setting a request for testing
	req, _ := http.NewRequest(http.MethodGet, "/admin/create-user", nil)
	req.Header.Set("Content-Type", "application/json")

	// Setting and recording the response
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(utils.MakeHTTPHandler(app.GetUsers))

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

	expectedStatusCode := http.StatusOK
	if rr.Code != expectedStatusCode {
		t.Errorf("%s: returned wrong status code; expected %d but got %d", "Status Code", expectedStatusCode, rr.Code)
	}
	expectedBody := `{"Success":"Successfully get all users!","Status":200,"Result":null}`
	if strings.TrimSpace(rr.Body.String()) != expectedBody {
		t.Errorf("Unexpected response body: expected %v, got %v", rr.Body.String(), expectedBody)
	}

}
