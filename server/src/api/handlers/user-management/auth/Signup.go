package handlers_user

import (
	"encoding/json"
	"io"
	"net/http"

	handlers_user_management "github.com/LeonLow97/inventory-management-system-golang-react-postgresql/api/handlers/user-management"
	"github.com/LeonLow97/inventory-management-system-golang-react-postgresql/database"
	"github.com/LeonLow97/inventory-management-system-golang-react-postgresql/utils"
)

func SignUp(w http.ResponseWriter, req *http.Request) {
	// Set Headers
	w.Header().Set("Content-Type", "application/json");
	var newUser handlers_user_management.SignUpJson

	// Reading the request body and UnMarshal the body to the LoginJson struct
	bs, _ := io.ReadAll(req.Body);
	if err := json.Unmarshal(bs, &newUser); err != nil {
		utils.InternalServerError(w, "Internal Server Error in UnMarshal JSON body in SignUp route: ", err)
		return;
	}
	
	// Default company name (only admin can create a different organisation name)
	userGroup := "InvenNexus User"
	organisationName := "InvenNexus" 
	isActive := 1

	// Trim whitespaces (username and email)
	newUser = newUser.UserFieldsTrimSpaces()

	// SignUp form validation
	if !handlers_user_management.SignUpFormValidation(w, newUser) {return}

	// Generate password hash
	hashedPassword := utils.GenerateHash(newUser.Password)

	// Check if username already exists in database (duplicates not allowed)
	isExistingUsername := database.GetUsername(newUser.Username)
	if (isExistingUsername) {
		utils.ResponseJson(w, http.StatusBadRequest, "Username has already been taken. Please try again.")
		return
	}

	// Check if email already exists in database (duplicates not allowed)
	isDuplicatedEmail := database.GetEmail(newUser.Email)
	if (isDuplicatedEmail) {
		utils.ResponseJson(w, http.StatusBadRequest, "Email address has already been taken. Please try again.")
		return
	}

	// Insert new user to `users` table
	userId, err := database.InsertNewUser(newUser.Username, hashedPassword, newUser.Email, isActive)
	if err != nil {
		utils.InternalServerError(w, "Internal Server Error in InsertNewUser: ", err)
		return
	}

	// Insert `user_organisation_mapping` table
	err = database.InsertIntoUserOrganisationMapping(userId, organisationName)
	if err != nil {
		utils.InternalServerError(w, "Internal Server Error in inserting into user_organisation_mapping: ", err)
		return
	}

	// Insert user_id and user_group_id into `user_group_mapping` table
	err = database.InsertIntoUserGroupMapping(userId, userGroup)
	if err != nil {
		utils.InternalServerError(w, "Internal Server Error in inserting into user_group_mapping: ", err)
		return
	}

	utils.ResponseJson(w, http.StatusOK, "Successfully Created User!");
}

