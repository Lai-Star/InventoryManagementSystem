package utils

import (
	"net/http"
	"strings"

	"github.com/LeonLow97/inventory-management-system-golang-react-postgresql/database"
)

// Determines if a user has been assigned that usergroup
func CheckUserGroup(w http.ResponseWriter, username string, userGroups ...string) bool {

	rows, err := database.GetUserGroupsByUsername(username)

	if err != nil {
		InternalServerError(w, "Internal server error in check user group: ", err)
		return false
	}

	var userGroup string

	for rows.Next() {
		err = rows.Scan(&userGroup)
		if err != nil {
			InternalServerError(w, "Internal server error in scanning user groups in check user group function: ", err)
			return false
		}

		if Contains(userGroups, userGroup) {
			return true
		}
	}

	ResponseJson(w, http.StatusForbidden, "Access Denied: You do not have permission to access this resource.")
	return false
}

func Contains(s []string, e string) bool {
	for _, a := range s {
		if strings.EqualFold(a, e) {
			return true
		}
	}
	return false
}
