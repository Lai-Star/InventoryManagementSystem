package user

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/LeonLow97/inventory-management-system-golang-react-postgresql/types"
	"github.com/LeonLow97/inventory-management-system-golang-react-postgresql/utils"
)

var (
	userGroup        = "InvenNexus User"
	organisationName = "InvenNexus"
	isActive         = 1
)

func (app application) SignUp(w http.ResponseWriter, req *http.Request) error {

	if req.Method != http.MethodPost {
		return utils.ApiError{Err: "Invalid Method", Status: http.StatusMethodNotAllowed}
	}

	var newUser types.SignUpJSON

	if err := newUser.ReadJSON(req.Body); err != nil {
		return utils.ApiError{Err: "Failed to decode JSON", Status: http.StatusInternalServerError}
	}

	// Setting timeout to follow SLA
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Minute*2)
	defer cancel()

	newUser.UserFieldsTrimSpaces()

	if err := newUser.SignUpFormValidation(w); err != nil {
		return err
	}

	newUser.Password = utils.GenerateHash(newUser.Password)

	// // Check if username already exists in database (duplicates not allowed)
	userCount, _ := app.DB.GetCountByUsername(newUser.Username)
	if userCount == 1 {
		return utils.ApiError{Err: "Username has already been taken. Please try again.", Status: http.StatusBadRequest}
	}

	// Check if email already exists in database (duplicates not allowed)
	emailCount, _ := app.DB.GetCountByEmail(newUser.Email)
	if emailCount == 1 {
		return utils.ApiError{Err: "Email has already been taken. Please try again.", Status: http.StatusBadRequest}
	}

	if err := app.DB.SignUpTransaction(ctx, newUser.Username, newUser.Password, newUser.Email, organisationName, userGroup, isActive); err != nil {
		log.Println("Internal Server Error in SignUpTransaction:", err)
		return utils.ApiError{Err: "Internal Server Error", Status: http.StatusInternalServerError}
	}

	return utils.WriteJSON(w, http.StatusCreated, utils.ApiSuccess{Success: "Successfully created user " + newUser.Username + "!", Status: http.StatusCreated})
}
