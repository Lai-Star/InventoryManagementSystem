package database

import (
	"fmt"
	"log"
)

var (
	SQL_DELETE_FROM_ACCOUNTS = "DELETE FROM accounts WHERE %s = $1;"
)

func DeleteUserFromAccounts(username string) error {
	_, err := db.Exec(fmt.Sprintf(SQL_DELETE_FROM_ACCOUNTS, "username"), username)
	if err != nil {
		log.Println("Internal Server Error deleting user from accounts: ", err)
	}
	return err
}