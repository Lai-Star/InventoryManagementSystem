package admin

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/LeonLow97/inventory-management-system-golang-react-postgresql/api/handlers/user-management/auth"
	"github.com/LeonLow97/inventory-management-system-golang-react-postgresql/types"
	"github.com/LeonLow97/inventory-management-system-golang-react-postgresql/utils"
)

func (app application) AdminCreateUserGroup(w http.ResponseWriter, req *http.Request) error {

	if req.Method != http.MethodPost {
		return utils.ApiError{Err: "Invalid Method", Status: http.StatusMethodNotAllowed}
	}

	var newUg types.AdminCreateUserGroupJSON

	if err := newUg.ReadJSON(req.Body); err != nil {
		log.Println("newUg.ReadJSON:", err)
		return utils.ApiError{Err: "Internal Server Error", Status: http.StatusInternalServerError}
	}

	// Setting timeout to follow SLA
	ctx := req.Context()
	ctx, cancel := context.WithTimeout(ctx, 2*time.Minute)
	defer cancel()

	// Check User Group Admin
	if err := auth.RetrieveIssuer(w, req); err != nil {
		return err
	}
	utilsApp := utils.Application{DB: app.DB}
	err := utils.InjectUG(utilsApp, ctx, w.Header().Get("username"), "Admin")
	if err != nil {
		return err
	}

	newUg.UGFieldsTrimSpaces()
	if err := newUg.UGFormValidation(w); err != nil {
		return err
	}

	// Check if user group already exists
	ugCount, err := app.DB.GetCountByUserGroup(ctx, newUg.UserGroup)
	if err != nil {
		return utils.ApiError{Err: "Internal Server Error", Status: http.StatusInternalServerError}
	}
	if ugCount == 1 {
		return utils.ApiError{Err: "User Group " + newUg.UserGroup + " already exist. Please try again.", Status: http.StatusBadRequest}
	}

	// Insert user group into user_groups table
	err = app.DB.InsertIntoUserGroups(ctx, newUg.UserGroup, newUg.Description)
	if err != nil {
		return utils.ApiError{Err: "Internal Server Error", Status: http.StatusInternalServerError}
	}

	return utils.WriteJSON(w, http.StatusCreated, utils.ApiSuccess{Success: "Successfully created user group!", Status: http.StatusCreated})
}
