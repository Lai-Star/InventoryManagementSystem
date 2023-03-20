package auth

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/LeonLow97/inventory-management-system-golang-react-postgresql/database"
	"github.com/LeonLow97/inventory-management-system-golang-react-postgresql/utils"
	"github.com/golang-jwt/jwt/v4"
)

// User login details: username, password
type LoginJson struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(w http.ResponseWriter, req *http.Request) {

	// Set Headers
	w.Header().Set("Content-Type", "application/json")
	var user LoginJson

	// Reading the request body and UnMarshal the body to the LoginJson struct
	bs, _ := io.ReadAll(req.Body)
	if err := json.Unmarshal(bs, &user); err != nil {
		utils.ResponseJson(w, http.StatusInternalServerError, "Internal Server Error")
		log.Println("Internal Server Error in UnMarshal JSON body in Login route:", err)
		return
	}

	// Check if username exists in database
	if !database.GetUsername(user.Username) {
		utils.ResponseJson(w, http.StatusUnauthorized, "You have entered an incorrect username and/or password. Please try again.")
		return
	}

	// Compare password with hashed password in database
	dbPassword, err := database.GetPasswordByUsername(user.Username)
	if err != nil {
		utils.ResponseJson(w, http.StatusInternalServerError, "Internal Server Error")
		log.Println("Internal server error in getting password in login route:", err)
		return
	}
	isValidPassword := utils.CompareHash(dbPassword, user.Password)
	if !isValidPassword {
		utils.ResponseJson(w, http.StatusUnauthorized, "You have entered an incorrect username and/or password. Please try again.")
		return
	}

	// Check User Status (active/inactive)
	dbStatus, err := database.GetActiveStatusByUsername(user.Username)
	if err != nil {
		utils.ResponseJson(w, http.StatusInternalServerError, "Internal Server Error")
		log.Println("Internal server error in getting active status in login route:", err)
		return
	}
	if dbStatus != 1 {
		utils.ResponseJson(w, http.StatusForbidden, "Your account has been disabled. Please contact the IMS Administrator.")
		return
	}

	// Generate Token with Claims
	tokenExpireTime := time.Now().Add(1 * time.Hour)
	generateToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Issuer:    user.Username,
		ExpiresAt: jwt.NewNumericDate(tokenExpireTime), // 1 hour
	})

	// Signing the jwt token with a secret key
	signedToken, err := generateToken.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		utils.ResponseJson(w, http.StatusInternalServerError, "Internal Server Error")
		log.Println("Internal Server Error in signing jwt token:", err)
		return
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

	utils.ResponseJson(w, http.StatusOK, "Successfully Logged In!")

	// Retrieve user's email to send OTP
	// dbEmail, _ := database.GetEmailByUsername(user.Username)
	// go utils.SMTP(user.Username, dbEmail, utils.Generate2FA())

}
