package database

import (
	"fmt"
)

var (
	SQL_DELETE_FROM_USERS = "DELETE FROM users WHERE %s = $1;"
)

func DeleteUserFromUsers(username string) error {
	_, err := db.Exec(fmt.Sprintf(SQL_DELETE_FROM_USERS, "username"), username)
	return err
}