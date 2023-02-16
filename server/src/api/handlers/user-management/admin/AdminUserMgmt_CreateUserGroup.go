package handlers_admin

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"

	handlers_user_management "github.com/LeonLow97/inventory-management-system-golang-react-postgresql/api/handlers/user-management"
	"github.com/LeonLow97/inventory-management-system-golang-react-postgresql/database"
	"github.com/LeonLow97/inventory-management-system-golang-react-postgresql/utils"
)

type AdminCreateUserGroupJson struct {
	UserGroup string `json:"user_group"`
	Description string `json:"description"`
}

func AdminCreateUserGroup(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json");
	var newUserGroup AdminCreateUserGroupJson

	// Reading the request body and UnMarshal the body to the AdminCreateUserGroup struct
	bs, _ := io.ReadAll(req.Body);
	if err := json.Unmarshal(bs, &newUserGroup); err != nil {
		utils.InternalServerError(w, "Internal Server Error in UnMarshal JSON body in AdminCreateUser route: ", err)
		return;
	}

	// Check User Group Admin
	if !handlers_user_management.RetrieveIssuer(w, req) {return}
	if !utils.CheckUserGroup(w, w.Header().Get("username"), "Admin") {return}

	userGroup := newUserGroup.UserGroup

	// trim user group
	userGroup = strings.TrimSpace(userGroup)

	// UserGroup form validation 
	isValidUserGroupValidation := UserGroupValidation(w, userGroup)
	if !isValidUserGroupValidation {return}

	// Check if user group already exists
	count, err := database.GetUserGroupCount(userGroup)
	if err != nil {
		utils.InternalServerError(w, "Internal server error in getting user group count: ", err)
		return
	}
	if count == 1 {
		utils.ResponseJson(w, http.StatusBadRequest, userGroup + " already exists. Please try again.")
		return
	}

	// Insert user group into user_groups table
	err = database.InsertIntoUserGroups(userGroup, newUserGroup.Description)
	if err != nil {
		utils.InternalServerError(w, "Internal server error in inserting into organisations table", err)
		return
	}

	utils.ResponseJson(w, http.StatusOK, "Successfully created a new organisation.")

}

func UserGroupValidation(w http.ResponseWriter, userGroup string) bool {
		
	// Check if user group is empty.
	if len(userGroup) == 0 {
		utils.ResponseJson(w, http.StatusBadRequest, "User Group cannot be empty. Please try again.")
		return false
	}

	// Check if user group has a length of more than 255 characters
	if len(userGroup) > 255 {
		utils.ResponseJson(w, http.StatusBadRequest, "User Group cannot have a length of more than 255 characters. Please try again.")
		return false
	}

	return true
}