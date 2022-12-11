package database

import (
	"database/sql"
)

var (
	SelectFromUser = "SELECT username FROM users WHERE username = $1"
	SelectEmailFromUser = "SELECT email FROM users WHERE email = $2;"
)

func CheckUsernameDuplicates(username string) bool {
	row := db.QueryRow(SelectFromUser, username)
	return row.Scan() != sql.ErrNoRows
}

func CheckEmailDuplicates(email string) bool {
	row := db.QueryRow(SelectEmailFromUser, email)
	return row.Scan() != sql.ErrNoRows
}

