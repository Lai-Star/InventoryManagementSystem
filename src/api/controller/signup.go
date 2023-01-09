package controller

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/LeonLow97/inventory-management-system-golang-react-postgresql/database"
	"github.com/LeonLow97/inventory-management-system-golang-react-postgresql/utils"
)

type SignUpJson struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	IsActive int    `json:"isActive"`
}
func SignUp(w http.ResponseWriter, req *http.Request) {
	// Set Headers
	w.Header().Set("Content-Type", "application/json");
	var newUser SignUpJson

	// Reading the request body and UnMarshal the body to the LoginJson struct
	bs, _ := io.ReadAll(req.Body);
	if err := json.Unmarshal(bs, &newUser); err != nil {
		utils.InternalServerError(w, "Internal Server Error in UnMarshal JSON body in SignUp route", err)
		return;
	}

	// New User sign up details
	username := newUser.Username
	password := newUser.Password
	email := newUser.Email
	isActive := newUser.IsActive

	isValidUsername := CheckUsernameFormat(w, username)
	if (!isValidUsername) {return}

	// Check if username already exists in database (duplicates not allowed)
	isDuplicateUsername := database.UsernameExists(username)
	if (isDuplicateUsername) {
		utils.ResponseJson(w, http.StatusBadRequest, "Username already exists. Please try again.")
		return
	}

	isValidPassword := CheckPasswordFormat(w, password);
	if (!isValidPassword) {return}
	hashedPassword := utils.GenerateHash(password)

	isValidEmail := CheckEmailFormat(w, email);
	if (!isValidEmail) {return}

	// Check is email already exists in database (duplicates not allowed)
	isDuplicatedEmail := database.EmailExists(email)
	if (isDuplicatedEmail) {
		utils.ResponseJson(w, http.StatusBadRequest, "Email address already exists. Please try again.")
		return
	}

	err := database.InsertNewUser(username, hashedPassword, email, isActive)
	utils.CheckError(err)

	utils.ResponseJson(w, http.StatusOK, "Successfully Created User!");
}

func CheckUsernameFormat(w http.ResponseWriter, username string) bool {
	// Check if username has a length >=8 or <=50
	isValidUsernameLength := utils.CheckLength(username, 8, 50)
	if (!isValidUsernameLength) {
		utils.ResponseJson(w, http.StatusBadRequest, "Username must have a length between 8 and 50 characters.")
		return false
	}

	// Check if username contains white spaces
	isValidUsernameSpaces := utils.CheckWhiteSpaces(username)
	if (!isValidUsernameSpaces) {
		utils.ResponseJson(w, http.StatusBadRequest, "Username cannot contain white spaces.")
		return false
	}

	// Check if username contains special characters (other than underscore _)
	isValidUsernameCharacters := utils.CheckSpecialChar(username)
	if (!isValidUsernameCharacters) {
		utils.ResponseJson(w, http.StatusBadRequest, "Username cannot contain special characters other than underscore ('_').")
		return false
	}
	return true
}

func CheckPasswordFormat(w http.ResponseWriter, password string) bool {
	// Check if password has a length >=8 and <=20
	isValidPasswordLength := utils.CheckLength(password, 8, 20)
	if (!isValidPasswordLength) {
		utils.ResponseJson(w, http.StatusBadRequest, "Password must have a length between 8 and 20 characters.")
		return false;
	}
	return true
}

func CheckEmailFormat(w http.ResponseWriter, email string) bool {
	// Check if email address is in the correct format
	isValidEmailAddress := utils.CheckEmailAddress(email)
	if (!isValidEmailAddress) {
		utils.ResponseJson(w, http.StatusBadRequest, "Email address is not in the correct format. Please try again.")
		return false
	}
	return true
}