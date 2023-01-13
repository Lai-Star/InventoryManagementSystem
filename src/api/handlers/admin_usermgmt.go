package handlers

import (
	"encoding/json"
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

func AdminCreateUser(w http.ResponseWriter, req *http.Request) {
	// Set Headers
	w.Header().Set("Content-Type", "application/json");
	var adminNewUser AdminUserMgmt

	// Reading the request body and UnMarshal the body to the LoginJson struct
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

	err := database.AdminInsertNewUser(username, hashedPassword, email, userGroup, companyName, isActive)
	utils.CheckErrorDatabase(err)
	if err != nil {
		utils.DatabaseServerError(w, "Internal Server Error: ", err)
	}

	utils.ResponseJson(w, http.StatusOK, "Successfully Created User!");
}