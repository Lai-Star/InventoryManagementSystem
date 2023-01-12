package utils

import "net/http"

func CheckUsernameFormat(w http.ResponseWriter, username string) bool {
	// Check if username is blank
	if username == "" {
		ResponseJson(w, http.StatusBadRequest, "Username cannot be empty.")
		return false
	}

	// Check if username has a length >=8 or <=50
	isValidUsernameLength := CheckLength(username, 8, 50)
	if (!isValidUsernameLength) {
		ResponseJson(w, http.StatusBadRequest, "Username must have a length between 8 and 50 characters.")
		return false
	}

	// Check if username contains white spaces
	isValidUsernameSpaces := CheckWhiteSpaces(username)
	if (!isValidUsernameSpaces) {
		ResponseJson(w, http.StatusBadRequest, "Username cannot contain white spaces.")
		return false
	}

	// Check if username contains special characters (other than underscore _)
	isValidUsernameCharacters := CheckSpecialChar(username)
	if (!isValidUsernameCharacters) {
		ResponseJson(w, http.StatusBadRequest, "Username cannot contain special characters other than underscore ('_').")
		return false
	}
	return true
}

func CheckPasswordFormat(w http.ResponseWriter, password string) bool {
	// Check if password is blank
	if password == "" {
		ResponseJson(w, http.StatusBadRequest, "Password cannot be empty.")
		return false
	}

	// Check if password has a length >=8 and <=20
	isValidPasswordLength := CheckLength(password, 8, 20)
	if (!isValidPasswordLength) {
		ResponseJson(w, http.StatusBadRequest, "Password must have a length between 8 and 20 characters.")
		return false;
	}
	return true
}

func CheckEmailFormat(w http.ResponseWriter, email string) bool {
	// Check if email is blank
	if email == "" {
		ResponseJson(w, http.StatusBadRequest, "Email cannot be empty.")
		return false
	}

	// Check if email address is in the correct format
	isValidEmailAddress := CheckEmailAddress(email)
	if (!isValidEmailAddress) {
		ResponseJson(w, http.StatusBadRequest, "Email address is not in the correct format. Please try again.")
		return false
	}
	return true
}