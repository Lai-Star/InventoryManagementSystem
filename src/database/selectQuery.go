package database

import (
	"database/sql"
)

var (
	SelectFromUser = "SELECT $1 FROM users WHERE username = $2;"
)

func CheckUsernameDuplicates(username string) bool {
	row := db.QueryRow(SelectFromUser, "username", username)
	
	if row.Scan() != sql.ErrNoRows {
		return true
	} 
	return false
}