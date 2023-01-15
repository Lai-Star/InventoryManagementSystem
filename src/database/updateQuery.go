package database

import (
	"fmt"
	"log"
)

var (
	queryUpdateUsers = "UPDATE accounts SET %s = $1, updated_date = now() WHERE %s = $2;"
)

func AdminUpdateUserPassword(username, password string) error {
	_, err := db.Exec(fmt.Sprintf(queryUpdateUsers, "password", "username"), password, username)
	if err != nil {
		log.Println("Error in Admin updating user password: ", err)
	}
	return err
}