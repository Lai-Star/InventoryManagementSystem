package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/LeonLow97/inventory-management-system-golang-react-postgresql/utils"
)

type AdminUserMgmt struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	UserGroup string `json:"user_group"`
	CompanyName string `json:"company_name"`
	IsActive int `json:"is_active"`
}

func AdminCreateUser(w http.ResponseWriter, req *http.Request) {
	// Set Headers
	w.Header().Set("Content-Type", "application/json");
	var adminNewUser AdminUserMgmt

	// Reading the request body and UnMarshal the body to the LoginJson struct
	bs, _ := io.ReadAll(req.Body);
	if err := json.Unmarshal(bs, &adminNewUser); err != nil {
		utils.InternalServerError(w, "Internal Server Error in UnMarshal JSON body in AdminCreateUser route", err)
		return;
	}

	// Create new user details
	username := adminNewUser.Username
	password := adminNewUser.Password
	email := adminNewUser.Email
	userGroup := adminNewUser.UserGroup
	companyName := adminNewUser.CompanyName
	isActive := adminNewUser.IsActive

	fmt.Println(username, password, email, userGroup, companyName, isActive)
}