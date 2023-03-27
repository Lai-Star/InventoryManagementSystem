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
		log.Println("newUser.ReadJSON:", err)
		return utils.ApiError{Err: "Internal Server Error", Status: http.StatusInternalServerError}
	}

	// Setting timeout to follow SLA
	ctx := req.Context()
	ctx, cancel := context.WithTimeout(ctx, 2*time.Minute)
	defer cancel()

	newUser.UserFieldsTrimSpaces()

	if err := newUser.SignUpFormValidation(w); err != nil {
		return err
	}

	newUser.Password = utils.GenerateHash(newUser.Password)

	// Check if username already exists in database (duplicates not allowed)
	userCount, err := app.DB.GetCountByUsername(ctx, newUser.Username)
	if err != nil {
		log.Println("app.DB.GetCountByUsername:", err)
		return utils.ApiError{Err: "Internal Server Error", Status: http.StatusInternalServerError}
	}
	if userCount == 1 {
		return utils.ApiError{Err: "Username has already been taken. Please try again.", Status: http.StatusBadRequest}
	}

	// Check if email already exists in database (duplicates not allowed)
	emailCount, err := app.DB.GetCountByEmail(ctx, newUser.Email)
	if err != nil {
		log.Println("app.DB.GetCountByEmail:", err)
		return utils.ApiError{Err: "Internal Server Error", Status: http.StatusInternalServerError}
	}
	if emailCount == 1 {
		return utils.ApiError{Err: "Email has already been taken. Please try again.", Status: http.StatusBadRequest}
	}

	if err := app.DB.SignUpTransaction(ctx, newUser.Username, newUser.Password, newUser.Email, organisationName, userGroup, isActive); err != nil {
		log.Println("Internal Server Error in SignUpTransaction:", err)
		return utils.ApiError{Err: "Internal Server Error", Status: http.StatusInternalServerError}
	}

	return utils.WriteJSON(w, http.StatusCreated, utils.ApiSuccess{Success: "User " + newUser.Username + " was successfully created!", Status: http.StatusCreated})
}
