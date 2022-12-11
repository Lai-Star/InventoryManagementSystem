package database

import "github.com/LeonLow97/inventory-management-system-golang-react-postgresql/utils"

var (
	InsertIntoUser = "INSERT INTO users (username, password, email, isActive, added_date, updated_date) VALUES ($1, $2, $3, $4, now(), now());"
)

func InsertNewUser(username string, password string, email string, isActive int) error {
	_, err := db.Exec(InsertIntoUser, username, password, email, isActive)
	utils.CheckErrorDatabase(err)
	return err
}