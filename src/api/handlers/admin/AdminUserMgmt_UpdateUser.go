package handlers_admin

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/LeonLow97/inventory-management-system-golang-react-postgresql/utils"
)

func AdminUpdateUser(w http.ResponseWriter, req *http.Request) {
	// Set Headers
	w.Header().Set("Content-Type", "application/json");
	var adminUpdateUser AdminUserMgmt

	// Reading the request body and UnMarshal the body to the AdminUserMgmt struct
	bs, _ := io.ReadAll(req.Body);
	if err := json.Unmarshal(bs, &adminUpdateUser); err != nil {
		utils.InternalServerError(w, "Internal Server Error in Unmarshal JSON body in AdminUpdateUser: ", err)
		return;
	}

	// Update user details
	username := adminUpdateUser.Username
	password := adminUpdateUser.Password
	email := adminUpdateUser.Email
	userGroup := adminUpdateUser.UserGroup
	companyName := adminUpdateUser.CompanyName
	isActive := adminUpdateUser.IsActive

	// Validate form inputs
	if !UserValidationForm(w, adminUpdateUser) {
		return 
	}

	fmt.Println(username, password, email, userGroup, companyName, isActive)
}