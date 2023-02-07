package handlers_user

import (
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/LeonLow97/inventory-management-system-golang-react-postgresql/utils"
	"github.com/golang-jwt/jwt/v4"
)

type SignUpJson struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func RetrieveIssuer(w http.ResponseWriter, req *http.Request) bool {

	w.Header().Set("Content-Type", "application/json");

	cookie, err := req.Cookie("leon-jwt-token")
	if err != nil {
		utils.InternalServerError(w, "Internal Server Error in retrieving cookie: ", err);
		return false;
	}

	// Parses the cookie jwt claims using the secret key to verify
	token, err := jwt.ParseWithClaims(cookie.Value, &jwt.RegisteredClaims{}, func(*jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	if err != nil {
		utils.InternalServerError(w, "Internal Server Error in parsing cookie: ", err)
		return false
	}

	// To access the issuer, we need the token claims to be type RegisteredClaims
	claims := token.Claims.(*jwt.RegisteredClaims)

	// fmt.Println("Retrieved Issuer using claims.Issuer: ", claims.Issuer)
	w.Header().Set("username", claims.Issuer)
	// fmt.Println("Retrieved Issuer using w.Header().Get():" , w.Header().Get("username"))

	return true
}

func (user SignUpJson) UserFieldsTrimSpaces() (SignUpJson) {
	user.Username = strings.TrimSpace(user.Username)
	user.Password = strings.TrimSpace(user.Password)
	user.Email = strings.TrimSpace(user.Email)

	return user
}

func SignUpFormValidation(w http.ResponseWriter, user SignUpJson) bool {
	
	// Ensure username is not blank
	if !utils.CheckMinLength(user.Username, 0) {
		utils.ResponseJson(w, http.StatusBadRequest, "Username cannot be blank. Please try again.")
		return false
	}

	// Ensure username has a length of 5 - 50 characters
	if !utils.CheckLengthRange(user.Username, 5, 50) {
		utils.ResponseJson(w, http.StatusBadRequest, "Username must have a length of 5 - 50 characters. Please try again.")
		return false
	}

	// Ensure username does not have special characters (only underscore is allowed)
	if !CheckUsernameSpecialCharacter(user.Username) {
		utils.ResponseJson(w, http.StatusBadRequest, "Username cannot contain special characters other than underscore (_). Please try again.")
		return false
	}

	// Ensure password is not blank
	if !utils.CheckMinLength(user.Password, 0) {
		utils.ResponseJson(w, http.StatusBadRequest, "Password cannot be blank. Please try again.")
		return false
	}

	// Ensure password has a length of 8 - 20 characters 
	if !utils.CheckLengthRange(user.Password, 8, 20) {
		utils.ResponseJson(w, http.StatusBadRequest, "Password must have a length of 8 - 20 characters. Please try again.")
		return false
	}

	// Email cannot be blank
	if !utils.CheckMinLength(user.Email, 0) {
		utils.ResponseJson(w, http.StatusBadRequest, "Email address cannot be blank. Please try again.")
		return false
	}

	// Ensure email has a maximum length of 255 characters.
	if !utils.CheckMaxLength(user.Email, 255) {
		utils.ResponseJson(w, http.StatusBadRequest, "Email address has a maximum length of 255 characters. Please try again.")
		return false
	}

	// Ensure email address is in the correct format
	if !CheckEmailAddressFormat(user.Email) {
		utils.ResponseJson(w, http.StatusBadRequest, "Email address is not in the correct format. Please try again.")
		return false
	}

	return true
}

// Returns true if string does not contain special characters (underscore is allowed)
func CheckUsernameSpecialCharacter(str string) bool {
	specialCharRegex := regexp.MustCompile(`^[a-zA-Z0-9_]*$`)
	return specialCharRegex.MatchString(str)
}

// Returns true if email address is in the correct format
func CheckEmailAddressFormat (email string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(email)
}