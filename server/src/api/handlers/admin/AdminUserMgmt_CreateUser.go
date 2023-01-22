package handlers_admin

import (
	"encoding/json"
	"io"
	"net/http"

	handlers_user "github.com/LeonLow97/inventory-management-system-golang-react-postgresql/api/handlers/user"
	"github.com/LeonLow97/inventory-management-system-golang-react-postgresql/database"
	"github.com/LeonLow97/inventory-management-system-golang-react-postgresql/utils"
)

func AdminCreateUser(w http.ResponseWriter, req *http.Request) {
	// Set Headers
	w.Header().Set("Content-Type", "application/json");
	var adminNewUser AdminUserMgmtJson

	// Reading the request body and UnMarshal the body to the AdminUserMgmt struct
	bs, _ := io.ReadAll(req.Body);
	if err := json.Unmarshal(bs, &adminNewUser); err != nil {
		utils.InternalServerError(w, "Internal Server Error in UnMarshal JSON body in AdminCreateUser route: ", err)
		return;
	}

	// Check User Group Admin
	if !CheckUserGroupAdmin(w, req) {return}

	// Check User Group
	handlers_user.RetrieveIssuer(w, req)
	checkUserGroup := utils.CheckUserGroup(w.Header().Get("username"), "Admin")
	if !checkUserGroup {
		utils.ResponseJson(w, http.StatusForbidden, "Access Denied: You do not have admin permissions to access this resource.")
		return
	}

	// Validate form inputs
	if !UserValidationForm(w, adminNewUser, "INSERT") {
		return 
	}

	hashedPassword := utils.GenerateHash(adminNewUser.Password)

	err := database.AdminInsertNewUser(adminNewUser.Username, hashedPassword, adminNewUser.Email, adminNewUser.UserGroup, adminNewUser.CompanyName, adminNewUser.IsActive)
	if err != nil {
		utils.InternalServerError(w, "Internal Server Error in AdminInsertNewUser: ", err)
		return
	}

	utils.ResponseJson(w, http.StatusOK, "Admin Successfully Created User!");
}