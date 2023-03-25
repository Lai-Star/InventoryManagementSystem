package user

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	auth "github.com/LeonLow97/inventory-management-system-golang-react-postgresql/api/handlers/user-management"
	"github.com/LeonLow97/inventory-management-system-golang-react-postgresql/database"
	"github.com/LeonLow97/inventory-management-system-golang-react-postgresql/utils"
)

type SignUpJson struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func SignUp(w http.ResponseWriter, req *http.Request) {
	// Set Headers
	w.Header().Set("Content-Type", "application/json")
	var newUser auth.SignUpJson

	// Reading the request body and UnMarshal the body to the LoginJson struct
	bs, _ := io.ReadAll(req.Body)
	if err := json.Unmarshal(bs, &newUser); err != nil {
		utils.ResponseJson(w, http.StatusInternalServerError, "Internal Server Error")
		log.Println("Internal Server Error in UnMarshal JSON body in SignUp route:", err)
		return
	}

	// Setting timeout to follow SLA
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Minute*2)
	defer cancel()
	if ctx == nil {
		fmt.Println("I am the culprit")
	}

	// Default company name (only admin can create a different organisation name)
	// userGroup := "InvenNexus User"
	// organisationName := "InvenNexus"
	// isActive := 1

	// Trim whitespaces (username and email)
	newUser = newUser.UserFieldsTrimSpaces()

	// SignUp form validation
	if !auth.SignUpFormValidation(w, newUser) {
		return
	}

	// Generate password hash
	// hashedPassword := utils.GenerateHash(newUser.Password)

	// Check if username already exists in database (duplicates not allowed)
	isExistingUsername := database.GetUsername(newUser.Username)
	if isExistingUsername {
		utils.ResponseJson(w, http.StatusBadRequest, "Username has already been taken. Please try again.")
		return
	}

	// // Check if email already exists in database (duplicates not allowed)
	// isDuplicatedEmail := database.GetEmail(newUser.Email)
	// if isDuplicatedEmail {
	// 	utils.ResponseJson(w, http.StatusBadRequest, "Email address has already been taken. Please try again.")
	// 	return
	// }

	// if err := database.SignUpTransaction(ctx, newUser.Username, hashedPassword, newUser.Email, organisationName, userGroup, isActive); err != nil {
	// 	utils.ResponseJson(w, http.StatusInternalServerError, "Internal Server Error")
	// 	log.Println("Error in SignUp Transaction", err)
	// 	return
	// }

	utils.ResponseJson(w, http.StatusOK, "Successfully Created User!")
}
