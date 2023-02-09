package handlers_admin

import (
	"encoding/json"
	"io"
	"net/http"

	handlers_user_management "github.com/LeonLow97/inventory-management-system-golang-react-postgresql/api/handlers/user-management"
	"github.com/LeonLow97/inventory-management-system-golang-react-postgresql/database"
	"github.com/LeonLow97/inventory-management-system-golang-react-postgresql/utils"
)

func AdminCreateUser(w http.ResponseWriter, req *http.Request) {
	// Set Headers
	w.Header().Set("Content-Type", "application/json");
	var adminNewUser handlers_user_management.AdminUserMgmtJson

	// Reading the request body and UnMarshal the body to the AdminUserMgmt struct
	bs, _ := io.ReadAll(req.Body);
	if err := json.Unmarshal(bs, &adminNewUser); err != nil {
		utils.InternalServerError(w, "Internal Server Error in UnMarshal JSON body in AdminCreateUser route: ", err)
		return;
	}

	// Check User Group Admin
	if !CheckUserGroupAdmin(w, req) {return}

	// Trim white spaces (username, password, email, company name)
	adminNewUser = adminNewUser.AdminUserMgmtFieldsTrimSpaces()

	// Validate form inputs
	if !handlers_user_management.AdminUserMgmtFormValidation(w, adminNewUser) {return}

	hashedPassword := utils.GenerateHash(adminNewUser.Password)

	err := database.AdminInsertNewUser(adminNewUser.Username, hashedPassword, adminNewUser.Email, adminNewUser.UserGroup, adminNewUser.OrganisationName, adminNewUser.IsActive)
	if err != nil {
		utils.InternalServerError(w, "Internal Server Error in AdminInsertNewUser: ", err)
		return
	}

	utils.ResponseJson(w, http.StatusOK, "Admin Successfully Created User!");
}