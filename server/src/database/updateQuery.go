package database

import (
	"fmt"
	"log"
)

var (
	SQL_UPDATE_ACCOUNTS = "UPDATE accounts SET %s = $1, %s = $2, %s = $3, %s = $4, %s = $5, updated_date = now() WHERE %s = $6;"
)

var (
	SQL_UPDATE_PRODUCTS = "UPDATE products SET %s = $1"
)

func AdminUpdateUser(username, password, email, userGroup, companyName string, isActive int) error {
	_, err := db.Exec(fmt.Sprintf(SQL_UPDATE_ACCOUNTS, "password", "email", "user_group", "company_name", "is_active" , "username"), password, email, userGroup, companyName, isActive, username)
	if err != nil {
		log.Println("Error in Admin updating user password: ", err)
	}
	return err
}

