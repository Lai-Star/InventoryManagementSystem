package user

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"

	auth_management "github.com/LeonLow97/inventory-management-system-golang-react-postgresql/api/handlers/user-management"
	"github.com/LeonLow97/inventory-management-system-golang-react-postgresql/database"
	"github.com/LeonLow97/inventory-management-system-golang-react-postgresql/utils"
)

func SignUp(w http.ResponseWriter, req *http.Request) {
	// Set Headers
	w.Header().Set("Content-Type", "application/json")
	var newUser auth_management.SignUpJson

	// Reading the request body and UnMarshal the body to the LoginJson struct
	bs, _ := io.ReadAll(req.Body)
	if err := json.Unmarshal(bs, &newUser); err != nil {
		utils.ResponseJson(w, http.StatusInternalServerError, "Internal Server Error")
		log.Println("Internal Server Error in UnMarshal JSON body in SignUp route:", err)
		return
	}

	ctx := context.Background()

	// Default company name (only admin can create a different organisation name)
	userGroup := "InvenNexus User"
	organisationName := "InvenNexus"
	isActive := 1

	// Trim whitespaces (username and email)
	newUser = newUser.UserFieldsTrimSpaces()

	// SignUp form validation
	if !auth_management.SignUpFormValidation(w, newUser) {
		return
	}

	// Generate password hash
	hashedPassword := utils.GenerateHash(newUser.Password)

	// Check if username already exists in database (duplicates not allowed)
	isExistingUsername := database.GetUsername(newUser.Username)
	if isExistingUsername {
		utils.ResponseJson(w, http.StatusBadRequest, "Username has already been taken. Please try again.")
		return
	}

	// Check if email already exists in database (duplicates not allowed)
	isDuplicatedEmail := database.GetEmail(newUser.Email)
	if isDuplicatedEmail {
		utils.ResponseJson(w, http.StatusBadRequest, "Email address has already been taken. Please try again.")
		return
	}

	if err := database.SignUpTransaction(ctx, newUser.Username, hashedPassword, newUser.Email, organisationName, userGroup, isActive); err != nil {
		utils.ResponseJson(w, http.StatusInternalServerError, "Internal Server Error")
		log.Println("Error in SignUp Transaction", err)
		return
	}

	utils.ResponseJson(w, http.StatusOK, "Successfully Created User!")
}
