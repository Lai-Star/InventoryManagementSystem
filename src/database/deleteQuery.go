package database

import (
	"fmt"
	"log"
)

var (
	queryDeleteFromAccounts = "DELETE FROM accounts WHERE %s = $1;"
)

func DeleteUserFromAccounts(username string) error {
	_, err := db.Exec(fmt.Sprintf(queryDeleteFromAccounts, "username"), username)
	if err != nil {
		log.Println("Internal Server Error deleting user from accounts: ", err)
	}
	return err
}