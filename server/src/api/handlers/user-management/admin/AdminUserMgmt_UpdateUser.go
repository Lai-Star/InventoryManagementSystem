package handlers_admin

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	handlers_user_mgmt "github.com/LeonLow97/inventory-management-system-golang-react-postgresql/api/handlers/user-management"
	"github.com/LeonLow97/inventory-management-system-golang-react-postgresql/database"
	"github.com/LeonLow97/inventory-management-system-golang-react-postgresql/utils"
)

func AdminUpdateUser(w http.ResponseWriter, req *http.Request) {
	// Set Headers
	w.Header().Set("Content-Type", "application/json");
	var adminUpdateUser handlers_user_mgmt.AdminUserMgmtJson

	// Reading the request body and UnMarshal the body to the AdminUserMgmt struct
	bs, _ := io.ReadAll(req.Body);
	if err := json.Unmarshal(bs, &adminUpdateUser); err != nil {
		utils.InternalServerError(w, "Internal Server Error in Unmarshal JSON body in AdminUpdateUser: ", err)
		return;
	}
	
	// Check User Group Admin
	// if !CheckUserGroupAdmin(w, req) {return}

	// Trim whitespaces (username, password, email, organisation_name)
	adminUpdateUser = adminUpdateUser.AdminUserMgmtFieldsTrimSpaces()
	
	// Check if username field is empty (mandatory field to update user, others can be empty)
	if len(adminUpdateUser.Username) == 0 {
		utils.ResponseJson(w, http.StatusBadRequest, "Please enter a username.")
		return
	}

	// Check if username exists in database
	isExistingUsername := database.GetUsername(adminUpdateUser.Username)
	if (!isExistingUsername) {
		utils.ResponseJson(w, http.StatusBadRequest, "Username does not exist. Please try again.")
		return
	}

	// Admin User Form Validation
	isValidAdminUserForm := handlers_user_mgmt.AdminUserMgmtFormValidation(w, adminUpdateUser, "UPDATE_USER")
	if !isValidAdminUserForm {return}

	// Check if email exists
	isExistingEmail := database.GetEmail(adminUpdateUser.Email)
	userCurrentEmail, err := database.GetEmailByUsername(adminUpdateUser.Username)
	if err != nil {
		utils.InternalServerError(w, "Internal server error in getting email by username: ", err)
		return
	}
	if isExistingEmail {
		if userCurrentEmail != adminUpdateUser.Email {
			utils.ResponseJson(w, http.StatusBadRequest, adminUpdateUser.Email + " already exists. Please try again.")
			return
		}
	}

	// Check if organisation name exists
	isExistingOrganisation := database.GetOrganisationName(adminUpdateUser.OrganisationName)
	if !isExistingOrganisation && len(adminUpdateUser.OrganisationName) > 0 {
		utils.ResponseJson(w, http.StatusNotFound, "Organisation name cannot be found. Please try again.")
		return
	}

	// Perform user group validation and check if user group exists
	isValidUserGroup := handlers_user_mgmt.UserGroupFormValidation(w, adminUpdateUser.UserGroup)
	if !isValidUserGroup {return}
	
	// Only generate hash if password is not empty
	if len(adminUpdateUser.Password) > 0 {
		adminUpdateUser.Password = utils.GenerateHash(adminUpdateUser.Password)
	}

	// Update users table (get user_id)
	userId, err := database.UpdateUsers(adminUpdateUser.Username, adminUpdateUser.Password, adminUpdateUser.Email, adminUpdateUser.IsActive)
	if err != nil {
		utils.InternalServerError(w, "Internal server error in admin update users table: ", err)
		return
	}

	// Update user_organisation_mapping table
	if len(adminUpdateUser.OrganisationName) > 0 {
		err = database.UpdateUserOrganisationMapping(userId, adminUpdateUser.OrganisationName)
		if err != nil {
			utils.InternalServerError(w, "Internal server error in admin update user_organisation_mapping table: ", err)
			return
		}
	}

	// Update user_group_mapping table
	if len(adminUpdateUser.UserGroup) > 0 {
		fmt.Println(adminUpdateUser.UserGroup)
		err = database.UpdateUserGroupMapping(userId, adminUpdateUser.UserGroup)
		if err != nil {
			utils.InternalServerError(w, "Internal server error in admin update user_group_mapping table: ", err)
			return
		}
	}

	utils.ResponseJson(w, http.StatusOK, "Successfully updated user!")
}