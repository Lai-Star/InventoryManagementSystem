package database

import "github.com/LeonLow97/inventory-management-system-golang-react-postgresql/utils"

var (
	InsertIntoAccounts = "INSERT INTO accounts (username, password, email, user_group, is_active, added_date, updated_date) VALUES ($1, $2, $3, $4, $5, now(), now());"
)

func InsertNewUser(username string, password string, email string, user_group string, isActive int) error {
	_, err := db.Exec(InsertIntoAccounts, username, password, email, user_group, isActive)
	utils.CheckErrorDatabase(err)
	return err
}