package database

import (
	"database/sql"
	"fmt"

	"github.com/LeonLow97/inventory-management-system-golang-react-postgresql/utils"
)

var (
	querySelectFromAccounts = "SELECT %s FROM accounts WHERE %s = $1;"
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
	utils.CheckErrorDatabase(err)
	return password, nil
}

func GetEmailFromDB(username string) (string, error) {
    var email string
    row := db.QueryRow(fmt.Sprintf(querySelectFromAccounts, "email", "username"), username)
    err := row.Scan(&email)
    utils.CheckErrorDatabase(err)
    return email, nil
}

func GetActiveStatusFromDB(username string) (int, error) {
	var isActive int
	row := db.QueryRow(fmt.Sprintf(querySelectFromAccounts, "is_active", "username"), username)
	err := row.Scan(&isActive)
	utils.CheckErrorDatabase(err)
	return isActive, nil
}

