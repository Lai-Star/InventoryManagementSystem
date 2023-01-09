package database

import (
	"database/sql"
	"fmt"

	"github.com/LeonLow97/inventory-management-system-golang-react-postgresql/utils"
)

var (
	querySelectFromUsers = "SELECT %s FROM users WHERE %s = $1;"
)

func UsernameExists(username string) bool {
	row := db.QueryRow(fmt.Sprintf(querySelectFromUsers, "username", "username"), username)
	return row.Scan() != sql.ErrNoRows
}

func EmailExists(email string) bool {
	row := db.QueryRow(fmt.Sprintf(querySelectFromUsers, "email", "email"), email)
	return row.Scan() != sql.ErrNoRows
}

func GetPasswordFromDB(username string) (string, error) {
	var password string
	row := db.QueryRow(fmt.Sprintf(querySelectFromUsers, "password", "username"), username)
	err := row.Scan(&password)
	utils.CheckErrorDatabase(err)
	return password, nil
}

