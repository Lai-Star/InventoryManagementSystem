package handlers_user_management

import (
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/LeonLow97/inventory-management-system-golang-react-postgresql/database"
	"github.com/LeonLow97/inventory-management-system-golang-react-postgresql/utils"
	"github.com/golang-jwt/jwt/v4"
)

type SignUpJson struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type AdminUserMgmtJson struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	UserGroup []string `json:"user_groups"`
	OrganisationName string `json:"organisation_name"`
	IsActive int `json:"is_active"`
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

// Form Validation: Username
func UsernameFormValidation(w http.ResponseWriter, username string) bool {
	// Ensure username is not blank
	if !utils.CheckMinLength(username, 0) {
		utils.ResponseJson(w, http.StatusBadRequest, "Username cannot be blank. Please try again.")
		return false
	}

	// Ensure username has a length of 5 - 50 characters
	if !utils.CheckLengthRange(username, 5, 50) {
		utils.ResponseJson(w, http.StatusBadRequest, "Username must have a length of 5 - 50 characters. Please try again.")
		return false
	}

	// Ensure username does not have special characters (only underscore is allowed)
	if !CheckUsernameSpecialCharacter(username) {
		utils.ResponseJson(w, http.StatusBadRequest, "Username cannot contain special characters other than underscore (_). Please try again.")
		return false
	}
	return true
}

// Form Validation: Password
func PasswordFormValidation(w http.ResponseWriter, password string) bool {
	// Ensure password is not blank
	if !utils.CheckMinLength(password, 0) {
		utils.ResponseJson(w, http.StatusBadRequest, "Password cannot be blank. Please try again.")
		return false
	}

	// Ensure password has a length of 8 - 20 characters 
	if !utils.CheckLengthRange(password, 8, 20) {
		utils.ResponseJson(w, http.StatusBadRequest, "Password must have a length of 8 - 20 characters. Please try again.")
		return false
	}

	// Check if password contains the correct format
	isValidPasswordCharacters := utils.CheckPasswordSpecialChar(password)
	if (!isValidPasswordCharacters) {
		utils.ResponseJson(w, http.StatusBadRequest, "Password must contain at least one lowercase, uppercase, numbers and special character.")
		return false
	}
	return true
}

// Form Validation: Email
func EmailFormValidation(w http.ResponseWriter, email string) bool {
	// Email cannot be blank
	if !utils.CheckMinLength(email, 0) {
		utils.ResponseJson(w, http.StatusBadRequest, "Email address cannot be blank. Please try again.")
		return false
	}

	// Ensure email has a maximum length of 255 characters.
	if !utils.CheckMaxLength(email, 255) {
		utils.ResponseJson(w, http.StatusBadRequest, "Email address has a maximum length of 255 characters. Please try again.")
		return false
	}

	// Ensure email address is in the correct format
	if !CheckEmailAddressFormat(email) {
		utils.ResponseJson(w, http.StatusBadRequest, "Email address is not in the correct format. Please try again.")
		return false
	}
	return true
}

func SignUpFormValidation(w http.ResponseWriter, user SignUpJson) bool {

	// Username form validation
	if !UsernameFormValidation(w, user.Username) {return false}

	// Password form validation
	if !PasswordFormValidation(w, user.Password) {return false}

	// Email form validation
	if !EmailFormValidation(w, user.Email) {return false}

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

// Trim Spaces for fields in the Admin User Management JSON
func (user AdminUserMgmtJson) AdminUserMgmtFieldsTrimSpaces() AdminUserMgmtJson {
	user.Username = strings.TrimSpace(user.Username)
	user.Password = strings.TrimSpace(user.Password)
	user.Email = strings.TrimSpace(user.Email)
	user.OrganisationName = strings.TrimSpace(user.OrganisationName)

	return user
}

func AdminUserMgmtFormValidation(w http.ResponseWriter, user AdminUserMgmtJson) bool {

	// Username form validation
	if !UsernameFormValidation(w, user.Username) {return false}

	// Password form validation
	if !PasswordFormValidation(w, user.Password) {return false}

	// Email form validation
	if !EmailFormValidation(w, user.Email) {return false}

	// If organisation name is blank, default to 'InvenNexus'
	if !utils.CheckMinLength(user.OrganisationName, 0) {
		user.OrganisationName = "InvenNexus"
	}

	// Ensure organisation name has a maximum length of 255 characters
	if !utils.CheckMaxLength(user.OrganisationName, 255) {
		utils.ResponseJson(w, http.StatusBadRequest, "Organisation name has a maximum length of 255 characters. Please try again.")
		return false
	}

	// Ensure isActive has values of 0 or 1
	if user.IsActive != 0 && user.IsActive != 1 {
		utils.ResponseJson(w, http.StatusBadRequest, "IsActive field can only have values Active or Inactive. Please try again.")
		return false
	}

	return true
}

func UserGroupFormValidation(w http.ResponseWriter, userGroups []string) bool {

	// iterate over the user groups array
	for idx, ug := range(userGroups) {
		// trim user group
		userGroups[idx] = strings.TrimSpace(ug)
		
		// check if user group belongs in the group of user groups
		count, err := database.GetUserGroupCount(ug)
		if err != nil {
			utils.InternalServerError(w, "Internal server error in getting user groups: ", err)
			return false
		}
		if count == 0 {
			utils.ResponseJson(w, http.StatusNotFound, ug + " does not exist. Please try again.")
			return false
		}
	}

	return true
}