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

func (app application) CreateUser(w http.ResponseWriter, req *http.Request) error {

	if req.Method != http.MethodPost {
		return utils.ApiError{Err: "Invalid Method", Status: http.StatusMethodNotAllowed}
	}

	var createUser types.AdminUserJSON

	if err := createUser.ReadJSON(req.Body); err != nil {
		log.Println("createUser.ReadJSON:", err)
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

	createUser.UserFieldsTrimSpaces()
	if err := createUser.CreateUserFormValidation(w); err != nil {
		return err
	}

	hashedPassword := utils.GenerateHash(createUser.Password)

	// Check for duplicates (username, email) and existing fields (organisation, user group)
	if err := app.DB.CheckDuplicatesAndExistingFieldsForCreateUser(ctx, createUser.Username, createUser.Email, createUser.OrganisationName, createUser.UserGroup...); err != nil {
		return err
	}

	if err := app.DB.CreateUserTransaction(ctx, createUser.Username, hashedPassword, createUser.Email, createUser.OrganisationName, createUser.IsActive, createUser.UserGroup...); err != nil {
		return err
	}

	return utils.WriteJSON(w, http.StatusCreated, utils.ApiSuccess{Success: "[Admin] Successfully created '" + createUser.Username + "' user!", Status: http.StatusCreated})

}
