package database

import "github.com/LeonLow97/inventory-management-system-golang-react-postgresql/utils"

var (
	InsertIntoAccounts = "INSERT INTO accounts (username, password, email, user_group, company_name, is_active, added_date, updated_date) VALUES ($1, $2, $3, $4, $5, $6, now(), now());"
)

func InsertNewUser(username, password, email, user_group, company_name string, isActive int) error {
	_, err := db.Exec(InsertIntoAccounts, username, password, email, user_group, company_name, isActive)
	utils.CheckErrorDatabase(err)
	return err
}

func AdminInsertNewUser(username, password, email, user_group, company_name string, isActive int) error {
	_, err := db.Exec(InsertIntoAccounts, username, password, email, user_group, company_name, isActive)
	utils.CheckErrorDatabase(err)
	return err
}