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
	hashedPassword := utils.GenerateHash(password);
	email := newUser.Email
	isActive := newUser.IsActive

	isValidUsername := CheckUsernameFormat(w, username)
	if (!isValidUsername) {return}

	// Check if username already exists in database (duplicates not allowed)
	isDuplicateUsername := database.CheckUsernameDuplicates(username)
	if (isDuplicateUsername) {
		utils.ResponseJson(w, http.StatusBadRequest, "Username already exists. Please try again.")
		return
	}

	err := database.InsertNewUser(username, hashedPassword, email, isActive)
	utils.CheckError(err)

	utils.ResponseJson(w, http.StatusOK, "Successfully Created User!");
}

func CheckUsernameFormat(w http.ResponseWriter, username string) bool {
	// Check if username has a length <8 or >50
	isValidUsernameLength := utils.CheckLength(username, 8, 50)
	if (!isValidUsernameLength) {
		utils.ResponseJson(w, http.StatusBadRequest, "Username must have a length between 8 and 50 characters")
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