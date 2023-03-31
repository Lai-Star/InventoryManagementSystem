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

type AdminDeleteUserMgmt struct {
	Username string `json:"username"`
}

func (app application) DeleteUser(w http.ResponseWriter, req *http.Request) error {

	if req.Method != http.MethodDelete {
		return utils.ApiError{Err: "Invalid Method", Status: http.StatusMethodNotAllowed}
	}

	var deleteUser types.AdminDeleteUserJSON

	if err := deleteUser.ReadJSON(req.Body); err != nil {
		log.Println("deleteUser.ReadJSON:", err)
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

	deleteUser.UserFieldsTrimSpaces()
	if err := deleteUser.DeleteUserFormValidation(w); err != nil {
		return err
	}

	// Check if username exists in the database
	usernameCount, err := app.DB.GetCountByUsername(ctx, deleteUser.Username)
	if err != nil {
		return utils.ApiError{Err: "Internal Server Error", Status: http.StatusInternalServerError}
	}
	if usernameCount != 1 {
		return utils.ApiError{Err: "User " + deleteUser.Username + " does not exist. Please try again!", Status: http.StatusNotFound}
	}

	// Delete the user
	if err := app.DB.DeleteUserByID(ctx, deleteUser.Username); err != nil {
		return utils.ApiError{Err: "Internal Server Error", Status: http.StatusInternalServerError}
	}

	return utils.WriteJSON(w, http.StatusOK, utils.ApiSuccess{Success: "Successfully deleted user " + deleteUser.Username + "!", Status: http.StatusOK})

}
