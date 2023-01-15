package handlers_admin

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/LeonLow97/inventory-management-system-golang-react-postgresql/api/handlers"
	"github.com/LeonLow97/inventory-management-system-golang-react-postgresql/database"
	"github.com/LeonLow97/inventory-management-system-golang-react-postgresql/utils"
)

func AdminUpdateUser(w http.ResponseWriter, req *http.Request) {
	// Set Headers
	w.Header().Set("Content-Type", "application/json");
	var adminUpdateUser AdminUserMgmt

	// Reading the request body and UnMarshal the body to the AdminUserMgmt struct
	bs, _ := io.ReadAll(req.Body);
	if err := json.Unmarshal(bs, &adminUpdateUser); err != nil {
		utils.InternalServerError(w, "Internal Server Error in Unmarshal JSON body in AdminUpdateUser: ", err)
		return;
	}

	// Update user with current user data (if none provided)
	adminUpdateUser = UpdateCurrentData(w, adminUpdateUser)
	
	adminUpdateUser.Password = utils.GenerateHash(adminUpdateUser.Password)

	fmt.Println(adminUpdateUser.Password)
	
	// Validate form inputs
	if !UserValidationForm(w, adminUpdateUser, "UPDATE") {return}

	// Update accounts table
	err := database.AdminUpdateUser(adminUpdateUser.Username, adminUpdateUser.Password, adminUpdateUser.Email, adminUpdateUser.UserGroup, adminUpdateUser.CompanyName, adminUpdateUser.IsActive)
	if err != nil {
		utils.InternalServerError(w, "Internal Server Error in AdminUpdateUser: ", err)
		return
	}

	utils.ResponseJson(w, http.StatusOK, "Successfully updated user!")
}

func UpdateCurrentData(w http.ResponseWriter, adminUpdateUser AdminUserMgmt) AdminUserMgmt {
	currentUserData, err := GetCurrentUserData(w, adminUpdateUser.Username)
	if err != nil {
		utils.InternalServerError(w, "Internal Server Error when getting current user data: ", err)
	}

	// Fill empty password
	if adminUpdateUser.Password == "" {
		adminUpdateUser.Password = currentUserData.Password
	}

	// Fill empty email
	if adminUpdateUser.Email == "" {
		adminUpdateUser.Email = currentUserData.Email
	}

	// Fill empty user group
	if adminUpdateUser.UserGroup == "" {
		adminUpdateUser.UserGroup = currentUserData.UserGroup
	}
	
	// Fill empty company name
	if adminUpdateUser.CompanyName == "" {
		adminUpdateUser.CompanyName = currentUserData.CompanyName
	}

	// Fill empty isActive
	if adminUpdateUser.IsActive == -1 {
		adminUpdateUser.IsActive = currentUserData.IsActive
	}

	return adminUpdateUser
}

func GetCurrentUserData(w http.ResponseWriter, username string) (handlers.User, error) {
	var password, email, userGroup, companyName, addedDate, updatedDate sql.NullString
	var isActive sql.NullInt16
	result := database.GetUserByUsername(username)

	var currentUserData handlers.User

	err := result.Scan(&username, &password, &email, &userGroup, &companyName, &isActive, &addedDate, &updatedDate)
	if err != sql.ErrNoRows {
		currentUserData.Password = password.String
		currentUserData.Email = email.String
		currentUserData.UserGroup = userGroup.String
		currentUserData.CompanyName = companyName.String
		currentUserData.IsActive = int(isActive.Int16)
		currentUserData.AddedDate = addedDate.String
		currentUserData.UpdatedDate = updatedDate.String
	} else {
		utils.InternalServerError(w, "Internal Server Error at getCurrentUserData: ", err)
		return handlers.User{}, err
	}

	return currentUserData, nil
}