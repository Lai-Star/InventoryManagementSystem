package database

import (
	"fmt"
	"log"
)

var (
	queryUpdateUsers = "UPDATE accounts SET %s = $1, %s = $2, %s = $3, %s = $4, %s = $5, updated_date = now() WHERE %s = $6;"
)

func AdminUpdateUser(username, password, email, userGroup, companyName string, isActive int) error {
	_, err := db.Exec(fmt.Sprintf(queryUpdateUsers, "password", "email", "user_group", "company_name", "is_active" , "username"), password, email, userGroup, companyName, isActive, username)
	if err != nil {
		log.Println("Error in Admin updating user password: ", err)
	}
	return err
}