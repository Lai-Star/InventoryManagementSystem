package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/LeonLow97/inventory-management-system-golang-react-postgresql/utils"
	"github.com/golang-jwt/jwt/v4"
)

// User login details: username, password
type LoginJson struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(w http.ResponseWriter, req *http.Request) {
	// Extract username and password from database
	dbUsername := "Leon"
	dbPassword := "$2a$10$YIofk.GOJ0jBlpjw1YzSKOQxcbN3bpyMRx82F4EDj/sBRpRkQlgDu"

	// Set Headers
	w.Header().Set("Content-Type", "application/json");
	var user LoginJson
	
	// Reading the request body and UnMarshal the body to the LoginJson struct
	bs, _ := io.ReadAll(req.Body);
	if err := json.Unmarshal(bs, &user); err != nil {
		utils.InternalServerError(w, "Internal Server Error in UnMarshal JSON body in Login route", err)
		return;
	}

	isValidPassword := utils.CompareHash(dbPassword, user.Password);
	if (!isValidPassword || dbUsername != user.Username) {
		utils.ResponseJson(w, http.StatusUnauthorized, "You have entered an incorrect username and/or password. Please try again.")
		return;
	}

	// Generate Token with Claims
	tokenExpireTime := time.Now().Add(1 * time.Hour)
	generateToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims {
		Issuer: dbUsername,
		ExpiresAt: jwt.NewNumericDate(tokenExpireTime), // 1 hour
	})

	// Signing the jwt token with a secret key 
	signedToken, err := generateToken.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		utils.InternalServerError(w, "Internal Server Error in signing jwt token: ", err)
		return;
	}

	c := &http.Cookie {
		Name: "leon-jwt-token",
		Value: signedToken,
		MaxAge: 3600,
		Path: "/",
		Domain: "localhost",
		Secure: false,
		HttpOnly: true,
	}

	// Setting a cookie with the signed jwt-token
	http.SetCookie(w, c)

	utils.ResponseJson(w, http.StatusOK, "Successfully Logged In!");
}

func Logout(w http.ResponseWriter, req *http.Request) {
	RetrieveIssuer(w, req)
	fmt.Println("Retrieved Issuer", w.Header().Get("username"))

	w.Header().Set("Content-Type", "application/json")
	cookie, err := req.Cookie("leon-jwt-token")
	if err != nil {
		utils.InternalServerError(w, "Error retrieving cookie: ", err)
		return;
	}
	cookie.Value = ""

	cookie = &http.Cookie {
		Name: "leon-jwt-token",
		Value: "",
		MaxAge: -1,
		Path: "",
		Domain: "",
		Secure: false,
		HttpOnly: true,
	}

	http.SetCookie(w, cookie)

	utils.ResponseJson(w, http.StatusOK, "Successfully Logged Out!")
}

func RetrieveIssuer(w http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("leon-jwt-token")
	if err != nil {
		utils.InternalServerError(w, "Internal Server Error in retrieving cookie: ", err);
		return;
	}

	// Parses the cookie jwt claims using the secret key to verify
	token, err := jwt.ParseWithClaims(cookie.Value, &jwt.RegisteredClaims{}, func(*jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	if err != nil {
		utils.InternalServerError(w, "Internal Server Error in parsing cookie: ", err)
		return
	}

	// To access the issuer, we need the token claims to be type RegisteredClaims
	claims := token.Claims.(*jwt.RegisteredClaims)

	fmt.Println("Retrieved Issuer using claims.Issuer: ", claims.Issuer)
	w.Header().Set("username", claims.Issuer)
	fmt.Println("Retrieved Issuer using w.Header().Get():" , w.Header().Get("username"))
}