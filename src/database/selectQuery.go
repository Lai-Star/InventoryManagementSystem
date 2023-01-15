package database

import (
	"database/sql"
	"fmt"
	"log"
)

var (
	querySelectFromAccounts = "SELECT %s FROM accounts WHERE %s = $1;"
	querySelectAllFromAccounts = "SELECT username, password, email, user_group, company_name, is_active, added_date, updated_date FROM accounts;"
)

func UsernameExists(username string) bool {
	row := db.QueryRow(fmt.Sprintf(querySelectFromAccounts, "username", "username"), username)
	return row.Scan() != sql.ErrNoRows
}

func EmailExists(email string) bool {
	row := db.QueryRow(fmt.Sprintf(querySelectFromAccounts, "email", "email"), email)
	return row.Scan() != sql.ErrNoRows
}

func GetPasswordFromDB(username string) (string, error) {
	var password string
	row := db.QueryRow(fmt.Sprintf(querySelectFromAccounts, "password", "username"), username)
	err := row.Scan(&password)
	if err != nil {
		log.Println("Error scanning when getting password from database: ", err)
	}
	return password, nil
}

func GetEmailFromDB(username string) (string, error) {
    var email string
    row := db.QueryRow(fmt.Sprintf(querySelectFromAccounts, "email", "username"), username)
    err := row.Scan(&email)
    if err != nil {
		log.Println("Error scanning when getting email from database: ", err)
	}
    return email, nil
}

func GetActiveStatusFromDB(username string) (int, error) {
	var isActive int
	row := db.QueryRow(fmt.Sprintf(querySelectFromAccounts, "is_active", "username"), username)
	err := row.Scan(&isActive)
	if err != nil {
		log.Println("Error scanning when getting isActive status from database: ", err)
	}
	return isActive, nil
}

func GetUserGroupFromDB(username string) (string, error) {
	var userGroup string
	row := db.QueryRow(fmt.Sprintf(querySelectFromAccounts, "user_group", "username"), username)
	err := row.Scan(&userGroup)
	if err != nil {
		log.Println("Error scanning when getting user group from database: ", err)
	}
	return userGroup, nil
}

func GetUsers() (*sql.Rows, error) {
	result, err := db.Query(querySelectAllFromAccounts)
	return result, err
}

