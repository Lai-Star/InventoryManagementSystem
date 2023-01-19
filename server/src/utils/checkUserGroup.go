package utils

import (
	"database/sql"
	"strings"

	"github.com/LeonLow97/inventory-management-system-golang-react-postgresql/database"
)

// Determines if a user has been assigned that usergroup
func CheckUserGroup(username, userGroup string) bool {
	var user_groups []string
	
	result := false

	userGroupDB, err := database.GetUserGroupFromDB(username)

	user_groups = strings.Split(userGroupDB, ",")

	if err == sql.ErrNoRows {
		result = false
	} else if err == nil {
		if Contains(user_groups, userGroup) {
			result = true
		} else {
			result = false
		}
	}

	return result
}