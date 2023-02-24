package database

import (
	"database/sql"
	"fmt"
)

var (
	SQL_SELECT_FROM_USERS = "SELECT %s FROM users WHERE %s = $1;"

	SQL_SELECT_FROM_ORGANISATIONS = "SELECT %s FROM organisations WHERE %s = $1;"

	SQL_SELECT_FROM_USERGROUPS = "SELECT COUNT(*) FROM user_groups WHERE %s = $1;"

	SQL_SELECT_ORGANISATION_NAME_BY_USERNAME = `SELECT o.organisation_name, u.user_id FROM organisations o
												INNER JOIN user_organisation_mapping uom
												ON o.organisation_id = uom.organisation_id
												INNER JOIN users u
												ON uom.user_id = u.user_id
												WHERE u.username = $1;`

	SQL_SELECT_USERGROUPS_BY_USERNAME = `SELECT ug.user_group FROM user_groups ug
										 LEFT JOIN user_group_mapping ugm 
										 ON ugm.user_group_id = ug.user_group_id 
										 WHERE ugm.user_id = (SELECT user_id FROM users WHERE username = $1);`
										 
	SQL_SELECT_ALL_USERS = `SELECT u.user_id, u.username, u.email, u.is_active, o.organisation_name, ug.user_group, u.added_date, u.updated_date FROM users u 
							LEFT JOIN user_organisation_mapping uom ON u.user_id = uom.user_id
							LEFT JOIN organisations o ON uom.organisation_id = o.organisation_id
							LEFT JOIN user_group_mapping ugm ON u.user_id = ugm.user_id
							LEFT JOIN user_groups ug ON ugm.user_group_id = ug.user_group_id
							ORDER BY user_id ASC;` 
)

func GetUsername(username string) bool {
	row := db.QueryRow(fmt.Sprintf(SQL_SELECT_FROM_USERS, "username", "username"), username)
	return row.Scan() != sql.ErrNoRows
}

func GetEmail(email string) bool {
	row := db.QueryRow(fmt.Sprintf(SQL_SELECT_FROM_USERS, "email", "email"), email)
	return row.Scan() != sql.ErrNoRows
}

func GetOrganisationName(organisationName string) bool {
	row := db.QueryRow(fmt.Sprintf(SQL_SELECT_FROM_ORGANISATIONS, "organisation_name", "organisation_name"), organisationName)
	return row.Scan() != sql.ErrNoRows
}

func GetPasswordByUsername(username string) (string, error) {
	var password string
	row := db.QueryRow(fmt.Sprintf(SQL_SELECT_FROM_USERS, "password", "username"), username)
	err := row.Scan(&password)
	return password, err
}

func GetEmailByUsername(username string) (string, error) {
    var email string
    row := db.QueryRow(fmt.Sprintf(SQL_SELECT_FROM_USERS, "email", "username"), username)
    err := row.Scan(&email)
    return email, err
}

func GetActiveStatusByUsername(username string) (int, error) {
	var isActive int
	row := db.QueryRow(fmt.Sprintf(SQL_SELECT_FROM_USERS, "is_active", "username"), username)
	err := row.Scan(&isActive)
	return isActive, err
}

func GetOrganisationNameByUsername(username string) (string, int, error) {
	var organisationName string
	var userId int
	err := db.QueryRow(SQL_SELECT_ORGANISATION_NAME_BY_USERNAME, username).Scan(&organisationName, &userId)
	return organisationName, userId, err
}

func GetUserGroupsByUsername(username string) (*sql.Rows, error) {
	rows, err := db.Query(SQL_SELECT_USERGROUPS_BY_USERNAME, username)
	return rows, err
}

func GetUserGroupCount(usergroup string) (int, error) {
	var count int
	row := db.QueryRow(fmt.Sprintf(SQL_SELECT_FROM_USERGROUPS, "user_group"), usergroup)
	err := row.Scan(&count)
	return count, err
}

func GetOrganisationNameCount(organisationName string) (int, error) {
	var count int
	row := db.QueryRow(fmt.Sprintf(SQL_SELECT_FROM_ORGANISATIONS, "COUNT(*)", "organisation_name"), organisationName)
	err := row.Scan(&count)
	return count, err
}

func GetUsers() (*sql.Rows, error) {
	row, err := db.Query(SQL_SELECT_ALL_USERS)
	return row, err
}
