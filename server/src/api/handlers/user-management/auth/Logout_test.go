package auth

import (
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/LeonLow97/inventory-management-system-golang-react-postgresql/database/repository/dbrepo"
	"github.com/golang-jwt/jwt/v4"
)

func Test_Logout(t *testing.T) {

	app := application{}
	app.DB = &dbrepo.TestDBRepo{}

	rr := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/logout", nil)

	tokenExpireTime := time.Now().Add(1 * time.Hour)
	generateToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Issuer:    "leon",
		ExpiresAt: jwt.NewNumericDate(tokenExpireTime), // 1 hour
	})
	signedToken, _ := generateToken.SignedString([]byte(os.Getenv("SECRET_KEY")))

	// Create a mock request with a cookie that contains the JWT token
	cookie := &http.Cookie{
		Name:     "leon-jwt-token",
		Value:    signedToken,
		MaxAge:   3600,
		Path:     "/",
		Domain:   "localhost",
		Secure:   false,
		HttpOnly: true,
	}
	req.AddCookie(cookie)

	app.Logout(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("expected status code 200 but got %d", status)
	}

	expectedResponse := `{"Success":"Successfully Logged out!","Status":200}`
	if strings.TrimSpace(rr.Body.String()) != expectedResponse {
		t.Errorf("unexpected response body: expected %v but got %v", expectedResponse, rr.Body.String())
	}

	actualCookie := rr.Header().Get("Set-Cookie")
	expectedCookie := "leon-jwt-token=; Max-Age=0; HttpOnly"
	if actualCookie != expectedCookie {
		t.Errorf("handler did not modify cookie correctly: expected %v but got %v ", expectedCookie, actualCookie)
	}

}
