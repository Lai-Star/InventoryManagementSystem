package controller

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/LeonLow97/inventory-management-system-golang-react-postgresql/database"
	"github.com/LeonLow97/inventory-management-system-golang-react-postgresql/utils"
)

type SignUpJson struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	IsActive int    `json:"isActive"`
}
func SignUp(w http.ResponseWriter, req *http.Request) {
	// Set Headers
	w.Header().Set("Content-Type", "application/json");
	var newUser SignUpJson

	// Reading the request body and UnMarshal the body to the LoginJson struct
	bs, _ := io.ReadAll(req.Body);
	if err := json.Unmarshal(bs, &newUser); err != nil {
		utils.InternalServerError(w, "Internal Server Error in UnMarshal JSON body in SignUp route", err)
		return;
	}

	// New User sign up details
	username := newUser.Username
	password := newUser.Password
	hashedPassword := utils.GenerateHash(password);
	email := newUser.Email
	isActive := newUser.IsActive

	err := database.CreateNewUser(username, hashedPassword, email, isActive)
	utils.CheckError(err)

	utils.ResponseJson(w, http.StatusOK, "Successfully Created User!");
}