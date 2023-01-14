package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/LeonLow97/inventory-management-system-golang-react-postgresql/database"
	"github.com/LeonLow97/inventory-management-system-golang-react-postgresql/utils"
)

type AdminUserMgmt struct {
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

type UserMgmtModel struct {
	Username string
	Password string
	Email string
	UserGroup string
	CompanyName string
	IsActive int
	AddedDate string
	UpdatedDate string
}

func AdminCreateUser(w http.ResponseWriter, req *http.Request) {
	// Set Headers
	w.Header().Set("Content-Type", "application/json");
	var adminNewUser AdminUserMgmt

	// Reading the request body and UnMarshal the body to the AdminUserMgmt struct
	bs, _ := io.ReadAll(req.Body);
	if err := json.Unmarshal(bs, &adminNewUser); err != nil {
		utils.InternalServerError(w, "Internal Server Error in UnMarshal JSON body in AdminCreateUser route", err)
		return;
	}

	// Create new user details
	username := adminNewUser.Username
	password := adminNewUser.Password
	email := adminNewUser.Email
	userGroup := adminNewUser.UserGroup
	companyName := adminNewUser.CompanyName
	isActive := adminNewUser.IsActive

	isValidUsername := utils.CheckUsernameFormat(w, username)
	if (!isValidUsername) {return}

	// Check if username already exists in database (duplicates not allowed)
	isDuplicateUsername := database.UsernameExists(username)
	if (isDuplicateUsername) {
		utils.ResponseJson(w, http.StatusBadRequest, "Username already exists. Please try again.")
		return
	}

	isValidPassword := utils.CheckPasswordFormat(w, password);
	if (!isValidPassword) {return}
	hashedPassword := utils.GenerateHash(password)

	isValidEmail := utils.CheckEmailFormat(w, email);
	if (!isValidEmail) {return}

	// Check if email already exists in database (duplicates not allowed)
	isDuplicatedEmail := database.EmailExists(email)
	if (isDuplicatedEmail) {
		utils.ResponseJson(w, http.StatusBadRequest, "Email address already exists. Please try again.")
		return
	}

	// Check if user group belongs to the following (Admin, IMS User, Operations, Financial Analyst)
	isValidUserGroup := utils.CheckUserGroupFormat(w, userGroup)
	if (!isValidUserGroup) {return}

	// Check if company name is between 5 and 250 characters and if blank company name provided (default to IMS)
	isValidCompanyName := utils.CheckCompanyNameFormat(w, companyName)
	if (!isValidCompanyName) {return}

	err := database.AdminInsertNewUser(username, hashedPassword, email, userGroup, companyName, isActive)
	utils.CheckErrorDatabase(err)
	if err != nil {
		utils.DatabaseServerError(w, "Internal Server Error: ", err)
	}

	utils.ResponseJson(w, http.StatusOK, "Admin Successfully Created User!");
}

func AdminGetUsers(w http.ResponseWriter, req *http.Request) {
	var data []UserMgmtModel
	var username, password, email, userGroup, companyName, addedDate, updatedDate sql.NullString
	var isActive sql.NullInt16

	rows, err := database.GetUsers()
	if err != nil {
		utils.DatabaseServerError(w, "Database Server Error", err)
		return 
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&username, &password, &email, &userGroup, &companyName, &isActive, &addedDate, &updatedDate)
		if err != nil {
			utils.DatabaseServerError(w, "Database Server Error", err)
			return
		}

		response := UserMgmtModel {
			Username: username.String, 
			Password: password.String,
			Email: email.String,
			UserGroup: userGroup.String,
			CompanyName: companyName.String,
			IsActive: int(isActive.Int16),
			AddedDate: addedDate.String,
			UpdatedDate: updatedDate.String,
		}

		data = append(data, response)
	
	}
	jsonStatus := struct {
		Code int `json:"code"`
		Response []UserMgmtModel `json:"response"`
	}{
		Response: data,
		Code: http.StatusOK,
	}

	bs, err := json.Marshal(jsonStatus);
	utils.CheckError(err)

	io.WriteString(w, string(bs));
}

func AdminUpdateUser(w http.ResponseWriter, req *http.Request) {
	// Set Headers
	w.Header().Set("Content-Type", "application/json");
	var adminUpdateUser AdminUserMgmt

	// Reading the request body and UnMarshal the body to the AdminUserMgmt struct
	bs, _ := io.ReadAll(req.Body);
	if err := json.Unmarshal(bs, &adminUpdateUser); err != nil {
		utils.CheckError(err)
		utils.InternalServerError(w, "Internal Server Error: ", err)
		return;
	}

	// Update user details
	username := adminUpdateUser.Username
	password := adminUpdateUser.Password
	email := adminUpdateUser.Email
	userGroup := adminUpdateUser.UserGroup
	companyName := adminUpdateUser.CompanyName
	isActive := adminUpdateUser.IsActive

	

	fmt.Println(username, password, email, userGroup, companyName, isActive)
}

func AdminDeleteUser(w http.ResponseWriter, req *http.Request) {
	// Set Headers
	w.Header().Set("Content-Type", "application/json");
	var adminDeleteUser AdminDeleteUserMgmt

	// Reading the request body and UnMarshal the body to the AdminUserMgmt struct
	bs, _ := io.ReadAll(req.Body);
	if err := json.Unmarshal(bs, &adminDeleteUser); err != nil {
		utils.CheckError(err)
		utils.InternalServerError(w, "Internal Server Error: ", err)
		return;
	}
}

