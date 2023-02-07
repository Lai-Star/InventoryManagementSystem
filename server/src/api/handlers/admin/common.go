package handlers_admin

import (
	"net/http"

	handlers_user "github.com/LeonLow97/inventory-management-system-golang-react-postgresql/api/handlers/user"
	"github.com/LeonLow97/inventory-management-system-golang-react-postgresql/database"
	"github.com/LeonLow97/inventory-management-system-golang-react-postgresql/utils"
)

type AdminUserMgmtJson struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	UserGroup string `json:"user_group"`
	CompanyName string `json:"company_name"`
	IsActive int `json:"is_active"`
}

type AdminDeleteUserMgmt struct {
	Username string `json:"username"`
}

func UserValidationForm(w http.ResponseWriter, adminUser AdminUserMgmtJson, action string) bool {
	username := adminUser.Username
	password := adminUser.Password
	email := adminUser.Email
	userGroup := adminUser.UserGroup
	companyName := adminUser.CompanyName

	isValidUsername := utils.CheckUsernameFormat(w, username)
	if (!isValidUsername) {return false}

	// Check if username already exists in database
	isUsernameExists := database.GetUsername(username)
	if action == "INSERT" && isUsernameExists {
		utils.ResponseJson(w, http.StatusBadRequest, "Username already exists. Please try again.")
		return false
	} else if action == "UPDATE" && !isUsernameExists {
		utils.ResponseJson(w, http.StatusBadRequest, "Username does not exist in database. Please try again.")
		return false
	}

	if action == "INSERT" {
		isValidPassword := utils.CheckPasswordFormat(w, password);
		if (!isValidPassword) {return false}
	} else if action == "UPDATE" && !(len(password) > 20) {
		isValidPassword := utils.CheckPasswordFormat(w, password);
		if (!isValidPassword) {return false}
	}

	isValidEmail := utils.CheckEmailFormat(w, email);
	if (!isValidEmail) {return false}

	// Check if email already exists in database 
	isEmailExists := database.EmailExists(email)
	if action == "UPDATE" {
		userEmail, _ := database.GetEmailFromDB(username)
		if isEmailExists && userEmail != email {
			utils.ResponseJson(w, http.StatusBadRequest, "Email address already exists. Please try again.")
			return false
		}
	} else if action == "INSERT" && isEmailExists {
		utils.ResponseJson(w, http.StatusBadRequest, "Email address already exists. Please try again.")
		return false
	} 

	// Check if user group belongs to the following (Admin, IMS User, Operations, Financial Analyst)
	isValidUserGroup := utils.CheckUserGroupFormat(w, userGroup)
	if (!isValidUserGroup) {return false}
	
	// Check if company name is between 5 and 250 characters and if blank company name provided (default to IMS)
	isValidCompanyName := utils.CheckCompanyNameFormat(w, companyName)
	if !isValidCompanyName {return false}

	return true
}

func CheckUserGroupAdmin(w http.ResponseWriter, req *http.Request) bool {
	// CheckUserGroup: IMS User and Operations
	if !handlers_user.RetrieveIssuer(w, req) {return false}
	
	checkUserGroupIMSUser := utils.CheckUserGroup(w.Header().Get("username"), "Admin")
	if !checkUserGroupIMSUser {
		utils.ResponseJson(w, http.StatusForbidden, "Access Denied: You do not have permission to access this resource.")
		return false
	}
	return true
}

