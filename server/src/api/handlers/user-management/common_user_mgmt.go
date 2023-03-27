package auth_management

import (
	"net/http"
	"strings"

	"github.com/LeonLow97/inventory-management-system-golang-react-postgresql/utils"
)

type AdminUserMgmtJson struct {
	Username         string   `json:"username"`
	Password         string   `json:"password"`
	Email            string   `json:"email"`
	UserGroup        []string `json:"user_groups"`
	OrganisationName string   `json:"organisation_name"`
	IsActive         int      `json:"is_active"`
}

type AdminCreateUserGroupJson struct {
	UserGroup []string `json:"user_group"`
}

// Form Validation: Username
func UsernameFormValidation(w http.ResponseWriter, username string) bool {
	// Ensure username is not blank
	if utils.IsBlankField(username) {
		utils.WriteJSON(w, http.StatusBadRequest, "Username cannot be blank. Please try again.")
		return false
	}

	// Ensure username has a length of 5 - 50 characters
	if !utils.CheckLengthRange(username, 5, 50) {
		utils.WriteJSON(w, http.StatusBadRequest, "Username must have a length of 5 - 50 characters. Please try again.")
		return false
	}

	// Ensure username does not have special characters (only underscore is allowed)
	if !utils.CheckUsernameSpecialChar(username) {
		utils.WriteJSON(w, http.StatusBadRequest, "Username cannot contain special characters other than underscore (_). Please try again.")
		return false
	}
	return true
}

// Form Validation: Password
func PasswordFormValidation(w http.ResponseWriter, password, action string) bool {
	// Ensure password is not blank
	if action == "CREATE_USER" {
		if utils.IsBlankField(password) {
			utils.WriteJSON(w, http.StatusBadRequest, "Password cannot be blank. Please try again.")
			return false
		}
	}

	if action == "CREATE_USER" || (action == "UPDATE_USER" && len(password) > 0) {
		// Ensure password has a length of 8 - 20 characters
		if !utils.CheckLengthRange(password, 8, 20) {
			utils.WriteJSON(w, http.StatusBadRequest, "Password must have a length of 8 - 20 characters. Please try again.")
			return false
		}

		// Check if password contains the correct format
		isValidPasswordCharacters := utils.CheckPasswordSpecialChar(password)
		if !isValidPasswordCharacters {
			utils.WriteJSON(w, http.StatusBadRequest, "Password must contain at least one lowercase, uppercase, numbers and special character.")
			return false
		}
	}

	return true
}

// Form Validation: Email
func EmailFormValidation(w http.ResponseWriter, email, action string) bool {
	// Email cannot be blank
	if action == "CREATE_USER" {
		if utils.IsBlankField(email) {
			utils.WriteJSON(w, http.StatusBadRequest, "Email address cannot be blank. Please try again.")
			return false
		}
	}

	if action == "CREATE_USER" || (action == "UPDATE_USER" && len(email) > 0) {
		// Ensure email has a maximum length of 255 characters.
		if !utils.CheckLengthRange(email, 1, 255) {
			utils.WriteJSON(w, http.StatusBadRequest, "Email address has a maximum length of 255 characters. Please try again.")
			return false
		}

		// Ensure email address is in the correct format
		if !utils.CheckEmailAddress(email) {
			utils.WriteJSON(w, http.StatusBadRequest, "Email address is not in the correct format. Please try again.")
			return false
		}
	}

	return true
}

// Form Validation: Organisation
func OrganisationFormValidation(w http.ResponseWriter, organisationName, action string) bool {

	if (action == "CREATE_USER" || action == "CREATE_ORGANISATION") && utils.IsBlankField(organisationName) {
		if action == "CREATE_ORGANISATION" {
			utils.WriteJSON(w, http.StatusBadRequest, "Please provide an organisation name.")
			return false
		} else {
			// If organisation name is blank, default to 'InvenNexus'
			organisationName = "InvenNexus"
		}
	}

	// Ensure organisation name has a maximum length of 255 characters
	if action == "CREATE_USER" || action == "CREATE_ORGANISATION" || (action == "UPDATE_USER" && len(organisationName) > 0) {
		if !utils.CheckLengthRange(organisationName, 1, 255) {
			utils.WriteJSON(w, http.StatusBadRequest, "Organisation name has a maximum length of 255 characters. Please try again.")
			return false
		}
	}

	return true
}

// Trim Spaces for fields in the Admin User Management JSON
func (user AdminUserMgmtJson) AdminUserMgmtFieldsTrimSpaces() AdminUserMgmtJson {
	user.Username = strings.TrimSpace(user.Username)
	user.Password = strings.TrimSpace(user.Password)
	user.Email = strings.TrimSpace(user.Email)
	user.OrganisationName = strings.TrimSpace(user.OrganisationName)

	return user
}

func AdminUserMgmtFormValidation(w http.ResponseWriter, user AdminUserMgmtJson, action string) bool {

	// Username form validation
	if !UsernameFormValidation(w, user.Username) {
		return false
	}

	// Password form validation
	if !PasswordFormValidation(w, user.Password, action) {
		return false
	}

	// Email form validation
	if !EmailFormValidation(w, user.Email, action) {
		return false
	}

	// Organisation form validation
	if !OrganisationFormValidation(w, user.OrganisationName, action) {
		return false
	}

	// Ensure isActive has values of 0 or 1
	if user.IsActive != 0 && user.IsActive != 1 {
		utils.WriteJSON(w, http.StatusBadRequest, "IsActive field can only have values Active or Inactive. Please try again.")
		return false
	}

	return true
}

func UserGroupFormValidation(w http.ResponseWriter, userGroups []string) bool {

	// iterate over the user groups array
	for idx, ug := range userGroups {
		// trim user group
		userGroups[idx] = strings.TrimSpace(ug)

		// Check if user group has a length of more than 255 characters
		if len(ug) > 255 {
			utils.WriteJSON(w, http.StatusBadRequest, "User Group cannot have a length of more than 255 characters. Please try again.")
			return false
		}

		// check if user group belongs in the group of user groups
		// count, err := database.GetUserGroupCount(ug)
		// if err != nil {
		// 	utils.WriteJSON(w, http.StatusInternalServerError, "Internal Server Error")
		// 	log.Println("Internal server error in getting user groups:", err)
		// 	return false
		// }
		// if count == 0 {
		// 	utils.WriteJSON(w, http.StatusNotFound, ug+" does not exist. Please try again.")
		// 	return false
		// }
	}

	return true
}
