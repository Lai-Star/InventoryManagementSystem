package database

import (
	"log"
)

var (
	SQL_INSERT_INTO_ACCOUNTS = "INSERT INTO accounts (username, password, email, user_group, company_name, is_active, added_date, updated_date) VALUES ($1, $2, $3, $4, $5, $6, now(), now());"
)

func InsertNewUser(username, password, email, user_group, company_name string, isActive int) error {
	_, err := db.Exec(SQL_INSERT_INTO_ACCOUNTS, username, password, email, user_group, company_name, isActive)
	if err != nil {
		log.Println("Error inserting new user to database: ", err)
	}
	return err
}

func AdminInsertNewUser(username, password, email, user_group, company_name string, isActive int) error {
	_, err := db.Exec(SQL_INSERT_INTO_ACCOUNTS, username, password, email, user_group, company_name, isActive)
	if err != nil {
		log.Println("Error admin inserting new user to database: ", err)
	}
	return err
}