package handlers_admin

import (
	"database/sql"
	"encoding/json"
	"io"
	"net/http"

	"github.com/LeonLow97/inventory-management-system-golang-react-postgresql/api/handlers"
	"github.com/LeonLow97/inventory-management-system-golang-react-postgresql/database"
	"github.com/LeonLow97/inventory-management-system-golang-react-postgresql/utils"
)

func AdminGetUsers(w http.ResponseWriter, req *http.Request) {

	// Check User Group Admin
	w.Header().Set("Content-Type", "application/json");

	// Check User Group (Admin)
	if !CheckUserGroupAdmin(w, req) {return}

	var data []handlers.User
	// To handle nullable columns in a database table
	var username, password, email, userGroup, organisationName, addedDate, updatedDate sql.NullString
	var isActive sql.NullInt16

	rows, err := database.GetUsers()
	if err != nil {
		utils.InternalServerError(w, "Internal Server Error in GetUsers: ", err)
		return 
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&username, &password, &email, &userGroup, &organisationName, &isActive, &addedDate, &updatedDate)
		if err != nil {
			utils.InternalServerError(w, "Internal Server Error in Scanning GetUsers: ", err)
			return
		}

		response := handlers.User {
			Username: username.String, 
			Password: password.String,
			Email: email.String,
			UserGroup: userGroup.String,
			OrganisationName: organisationName.String,
			IsActive: int(isActive.Int16),
			AddedDate: addedDate.String,
			UpdatedDate: updatedDate.String,
		}

		data = append(data, response)
	
	}
	jsonStatus := struct {
		Code int `json:"code"`
		Response []handlers.User `json:"response"`
	}{
		Response: data,
		Code: http.StatusOK,
	}

	bs, err := json.Marshal(jsonStatus);
	if err != nil {
		utils.InternalServerError(w, "Internal Server Error in Marshal JSON body in GetUsers: ", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	io.WriteString(w, string(bs));
}