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

func (app application) UpdateUser(w http.ResponseWriter, req *http.Request) error {

	if req.Method != http.MethodPatch {
		return utils.ApiError{Err: "Invalid Method", Status: http.StatusMethodNotAllowed}
	}

	var updateUser types.AdminUserJSON

	if err := updateUser.ReadJSON(req.Body); err != nil {
		log.Println("updateUser.ReadJSON:", err)
		return utils.ApiError{Err: "Internal Server Error", Status: http.StatusInternalServerError}
	}

	// Setting timeout to follow SLA
	ctx := req.Context()
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
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

	// updateUser.UserFieldsTrimSpaces()
	// if err := updateUser.UpdateUserFormValidation(w); err != nil {
	// 	return err
	// }

	// // Call CheckDuplicatesAndExistingFieldsForUpdateUser with a new context
	// checkDuplicatesCtx, checkDuplicatesCancel := context.WithTimeout(context.Background(), 5*time.Second)
	// defer checkDuplicatesCancel()

	// if err := app.DB.CheckDuplicatesAndExistingFieldsForUpdateUser(checkDuplicatesCtx, updateUser.Username, updateUser.Email, updateUser.OrganisationName, updateUser.UserGroup...); err != nil {
	// 	return err
	// }

	// // Only generate hash if password is not empty
	// if len(updateUser.Password) > 0 {
	// 	updateUser.Password = utils.GenerateHash(updateUser.Password)
	// }

	// // Call CheckDuplicatesAndExistingFieldsForUpdateUser with a new context
	// test, testCancel := context.WithTimeout(context.Background(), 5*time.Second)
	// defer testCancel()

	// if err := app.DB.UpdateUserTransaction(test, updateUser.Username, updateUser.Password, updateUser.Email, updateUser.OrganisationName, updateUser.IsActive, updateUser.UserGroup); err != nil {
	// 	return err
	// }

	return utils.WriteJSON(w, http.StatusOK, utils.ApiSuccess{Success: "[Admin] Successfully updated '" + updateUser.Username + "' user!", Status: http.StatusOK})

}
