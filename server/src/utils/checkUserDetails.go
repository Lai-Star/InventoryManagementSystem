package utils

import (
	"net/http"
	"strings"
)

func CheckUsernameFormat(w http.ResponseWriter, username string) bool {
	username = strings.TrimSpace(username)

	// Check if username is blank
	if username == "" {
		ResponseJson(w, http.StatusBadRequest, "Please enter a username.")
		return false
	}

	// Check if username has a length >=5 or <=50
	isValidUsernameLength := CheckLength(username, 5, 50)
	if (!isValidUsernameLength) {
		ResponseJson(w, http.StatusBadRequest, "Username must have a length between 5 and 50 characters.")
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
	email = strings.TrimSpace(email)

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

func CheckUserGroupFormat(w http.ResponseWriter, userGroups string) bool {
	if userGroups == "" {
		userGroups = "IMS User"
	}

	userGroups = strings.TrimSpace(userGroups)

	groups := strings.Split(userGroups, ",")
	validGroups := []string{"Admin", "IMS User", "Operations", "Financial Analyst"}
	for _, group := range groups {
		group = strings.TrimSpace(group)
		if !Contains(validGroups, group) {
			ResponseJson(w, http.StatusBadRequest, group + " user group is not registered. Please try again.")
			return false
		}
	}
	return true
}

func CheckCompanyNameFormat(w http.ResponseWriter, companyName string) bool {
	if companyName == "" {
		companyName = "IMSer"
	}

	companyName = strings.TrimSpace(companyName)

	isValidCompanyNameLength := CheckLength(companyName, 5, 250)
	if (!isValidCompanyNameLength) {
		ResponseJson(w, http.StatusBadRequest, "Company Name must be between 5 and 250 characters. Please try again.")
		return false
	}

	return true
}

func Contains(s []string, e string) bool {
    for _, a := range s {
        if strings.EqualFold(a, e) {
            return true
        }
    }
    return false
}
