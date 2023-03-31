package admin

import (
	"context"
	"net/http"
	"time"

	"github.com/LeonLow97/inventory-management-system-golang-react-postgresql/api/handlers"
	"github.com/LeonLow97/inventory-management-system-golang-react-postgresql/api/handlers/user-management/auth"
	"github.com/LeonLow97/inventory-management-system-golang-react-postgresql/utils"
)

func (app application) GetUsers(w http.ResponseWriter, req *http.Request) error {

	if req.Method != http.MethodGet {
		return utils.ApiError{Err: "Invalid Method", Status: http.StatusMethodNotAllowed}
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

	var data []handlers.User
	var users = make(map[int]handlers.User) // A map to store unique user records by UserId

	data, err = app.DB.GetAllUsers(ctx, data, users)
	if err != nil {
		return err
	}

	return utils.WriteJSON(w, http.StatusOK, utils.ApiGetSuccess{Success: "Successfully get all users!", Status: http.StatusOK, Result: data})
}
