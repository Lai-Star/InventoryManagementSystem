package auth

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/LeonLow97/inventory-management-system-golang-react-postgresql/types"
	"github.com/LeonLow97/inventory-management-system-golang-react-postgresql/utils"
	"github.com/golang-jwt/jwt/v4"
)

func (app application) Login(w http.ResponseWriter, req *http.Request) error {

	if req.Method != http.MethodPost {
		return utils.ApiError{Err: "Invalid Method", Status: http.StatusMethodNotAllowed}
	}

	var authUser types.LoginJSON

	if err := authUser.ReadJSON(req.Body); err != nil {
		log.Println("authUser.ReadJSON:", err)
		return utils.ApiError{Err: "Internal Server Error", Status: http.StatusInternalServerError}
	}

	// Setting timeout to follow SLA
	ctx := req.Context()
	ctx, cancel := context.WithTimeout(ctx, 2*time.Minute)
	defer cancel()

	authUser.AuthFieldsTrimSpaces()

	// Sanitise username and password input values to mitigate SQL Injection
	if err := authUser.AuthFormValidation(w); err != nil {
		return err
	}

	// Check if username exists in database
	userCount, err := app.DB.GetCountByUsername(ctx, authUser.Username)
	if err != nil {
		log.Println("app.DB.GetCountByUsername:", err)
		return utils.ApiError{Err: "Internal Server Error", Status: http.StatusInternalServerError}
	}
	if userCount == 0 {
		return utils.ApiError{Err: "You have entered an incorrect username and/or password. Please try again.", Status: http.StatusUnauthorized}
	}

	// Compare password with hashed password in database
	dbPassword, err := app.DB.GetPasswordByUsername(ctx, authUser.Username)
	if err != nil {
		log.Println("app.DB.GetPasswordByUsername:", err)
		return utils.ApiError{Err: "Internal Server Error", Status: http.StatusInternalServerError}
	}
	if !utils.CompareHash(dbPassword, authUser.Password) {
		return utils.ApiError{Err: "You have entered an incorrect username and/or password. Please try again.", Status: http.StatusUnauthorized}
	}

	// Check User Account Status (active/inactive)
	dbIsActive, err := app.DB.GetIsActiveByUsername(ctx, authUser.Username)
	if err != nil {
		log.Println("app.DB.GetIsActiveByUsername:", err)
		return utils.ApiError{Err: "Internal Server Error", Status: http.StatusInternalServerError}
	}
	if dbIsActive == 0 {
		return utils.ApiError{Err: "This account has been disabled. Please contact the System Administrator", Status: http.StatusForbidden}
	}

	// Generate Token with Claims
	tokenExpireTime := time.Now().Add(1 * time.Hour)
	generateToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Issuer:    authUser.Username,
		ExpiresAt: jwt.NewNumericDate(tokenExpireTime), // 1 hour
	})

	// Signing the jwt token with a secret key
	signedToken, err := generateToken.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		log.Println("generateToken.SignedString:", err)
		return utils.ApiError{Err: "Internal Server Error", Status: http.StatusInternalServerError}
	}

	c := &http.Cookie{
		Name:     "leon-jwt-token",
		Value:    signedToken,
		MaxAge:   3600,
		Path:     "/",
		Domain:   "localhost",
		Secure:   false,
		HttpOnly: true,
	}

	// Setting a cookie with the signed jwt-token
	http.SetCookie(w, c)

	return utils.WriteJSON(w, http.StatusOK, utils.ApiSuccess{Success: "Login Successful!", Status: http.StatusOK})

	// // Retrieve user's email to send OTP
	// // dbEmail, _ := database.GetEmailByUsername(user.Username)
	// // go utils.SMTP(user.Username, dbEmail, utils.Generate2FA())
}
