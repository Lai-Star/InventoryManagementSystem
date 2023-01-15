package handlers_user

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
	CompanyName string `json:"company_name"`
}

func SignUp(w http.ResponseWriter, req *http.Request) {
	// Set Headers
	w.Header().Set("Content-Type", "application/json");
	var newUser SignUpJson

	// Reading the request body and UnMarshal the body to the LoginJson struct
	bs, _ := io.ReadAll(req.Body);
	if err := json.Unmarshal(bs, &newUser); err != nil {
		utils.InternalServerError(w, "Internal Server Error in UnMarshal JSON body in SignUp route: ", err)
		return;
	}

	// New User sign up details
	username := newUser.Username
	password := newUser.Password
	email := newUser.Email
	
	company_name := "IMSer"
	isActive := 1

	isValidUsername := utils.CheckUsernameFormat(w, username)
	if (!isValidUsername) {return}

	// Check if username already exists in database (duplicates not allowed)
	isDuplicateUsername := database.UsernameExists(username)
	if (isDuplicateUsername) {
		utils.ResponseJson(w, http.StatusBadRequest, "Username already exists. Please try again.")
		return
	}

	isValidPassword := utils.CheckPasswordFormat(w, password);
	if (!isValidPassword) {return}
	hashedPassword := utils.GenerateHash(password)

	isValidEmail := utils.CheckEmailFormat(w, email);
	if (!isValidEmail) {return}

	// Check if email already exists in database (duplicates not allowed)
	isDuplicatedEmail := database.EmailExists(email)
	if (isDuplicatedEmail) {
		utils.ResponseJson(w, http.StatusBadRequest, "Email address already exists. Please try again.")
		return
	}

	// For new users, userGroup is default to "Retail Business Owner"
	userGroup := "IMS User"

	err := database.InsertNewUser(username, hashedPassword, email, userGroup, company_name, isActive)
	if err != nil {
		utils.InternalServerError(w, "Internal Server Error in InsertNewUser: ", err)
	}

	utils.ResponseJson(w, http.StatusOK, "Successfully Created User!");
}

