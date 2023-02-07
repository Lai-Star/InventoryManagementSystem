package handlers_admin

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/LeonLow97/inventory-management-system-golang-react-postgresql/database"
	"github.com/LeonLow97/inventory-management-system-golang-react-postgresql/utils"
)

func AdminDeleteUser(w http.ResponseWriter, req *http.Request) {
	// Set Headers
	w.Header().Set("Content-Type", "application/json");
	var adminDeleteUser AdminDeleteUserMgmt

	// Reading the request body and UnMarshal the body to the AdminDeleteUserMgmt struct
	bs, _ := io.ReadAll(req.Body);
	if err := json.Unmarshal(bs, &adminDeleteUser); err != nil {
		utils.InternalServerError(w, "Internal Server Error in Unmarshal JSON body in AdminDeleteUser: ", err)
		return;
	}

	// Check User Group Admin
	if !CheckUserGroupAdmin(w, req) {return}

	username := adminDeleteUser.Username

	// Check username format
	if !utils.CheckUsernameFormat(w, username) {return}

	// Check if username exists in the database
	if !database.GetUsername(username) {
		utils.ResponseJson(w, http.StatusNotFound, "Username does not exist in database. Please try again.")
		return
	}

	err := database.DeleteUserFromAccounts(username)
	if err != nil {
		utils.InternalServerError(w, "Internal Server Error in deleting user from accounts table: ", err)
		return
	}
	
	utils.ResponseJson(w, http.StatusOK, "Successfully Deleted User!")
}