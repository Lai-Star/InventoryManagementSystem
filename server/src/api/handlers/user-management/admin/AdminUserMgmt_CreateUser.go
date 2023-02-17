package handlers_admin

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	handlers_user_management "github.com/LeonLow97/inventory-management-system-golang-react-postgresql/api/handlers/user-management"
	"github.com/LeonLow97/inventory-management-system-golang-react-postgresql/database"
	"github.com/LeonLow97/inventory-management-system-golang-react-postgresql/utils"
)

func AdminCreateUser(w http.ResponseWriter, req *http.Request) {
	// Set Headers
	w.Header().Set("Content-Type", "application/json");
	var adminNewUser handlers_user_management.AdminUserMgmtJson

	// Reading the request body and UnMarshal the body to the AdminUserMgmt struct
	bs, _ := io.ReadAll(req.Body);
	if err := json.Unmarshal(bs, &adminNewUser); err != nil {
		utils.InternalServerError(w, "Internal Server Error in UnMarshal JSON body in AdminCreateUser route: ", err)
		return;
	}

	// Check User Group Admin
	// if !CheckUserGroupAdmin(w, req) {return}

	// Trim white spaces (username, password, email, company name)
	adminNewUser = adminNewUser.AdminUserMgmtFieldsTrimSpaces()

	// Validate form inputs
	if !handlers_user_management.AdminUserMgmtFormValidation(w, adminNewUser) {return}

	hashedPassword := utils.GenerateHash(adminNewUser.Password)

	// Check if username already exists in database (duplicates not allowed)
	isExistingUsername := database.GetUsername(adminNewUser.Username)
	if (isExistingUsername) {
		utils.ResponseJson(w, http.StatusBadRequest, "Username has already been taken. Please try again.")
		return
	}

	// Check if email already exists in database (duplicates not allowed)
	isExistingEmail := database.GetEmail(adminNewUser.Email)
	if (isExistingEmail) {
		utils.ResponseJson(w, http.StatusBadRequest, "Email address has already been taken. Please try again.")
		return
	}

	// Check if organisation already exists in database
	isExistingOrganisation := database.GetOrganisationName(adminNewUser.OrganisationName)
	if !isExistingOrganisation {
		utils.ResponseJson(w, http.StatusNotFound, "Organisation name cannot be found. Please try again.")
		return
	}

	// Check if user group is valid and trim user group
	isValidUserGroup := handlers_user_management.UserGroupFormValidation(w, adminNewUser.UserGroup)
	if !isValidUserGroup {return}

	// Insert users table
	fmt.Println(hashedPassword) // Insert hashed password
	// Insert user_organisation_mapping table
	// Insert user_group_mapping table

	// err := database.AdminInsertNewUser(adminNewUser.Username, hashedPassword, adminNewUser.Email, adminNewUser.UserGroup, adminNewUser.OrganisationName, adminNewUser.IsActive)
	// if err != nil {
	// 	utils.InternalServerError(w, "Internal Server Error in AdminInsertNewUser: ", err)
	// 	return
	// }

	utils.ResponseJson(w, http.StatusOK, "Admin Successfully Created User!");
}