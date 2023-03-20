package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
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
	row := conn.QueryRow(context.Background(), fmt.Sprintf(SQL_SELECT_FROM_USERS, "username", "username"), username)
	return row.Scan() != pgx.ErrNoRows
}

func GetEmail(email string) bool {
	row := conn.QueryRow(context.Background(), fmt.Sprintf(SQL_SELECT_FROM_USERS, "email", "email"), email)
	return row.Scan() != pgx.ErrNoRows
}

func GetOrganisationName(organisationName string) bool {
	row := conn.QueryRow(context.Background(), fmt.Sprintf(SQL_SELECT_FROM_ORGANISATIONS, "organisation_name", "organisation_name"), organisationName)
	return row.Scan() != pgx.ErrNoRows
}

func GetPasswordByUsername(username string) (string, error) {
	var password string
	if err := conn.QueryRow(context.Background(), fmt.Sprintf(SQL_SELECT_FROM_USERS, "password", "username"), username).Scan(&password); err != nil {
		return "", fmt.Errorf("conn.QueryRow in GetPasswordByUsername: %w", err)
	}
	return password, nil
}

func GetEmailByUsername(username string) (string, error) {
	var email string
	if err := conn.QueryRow(context.Background(), fmt.Sprintf(SQL_SELECT_FROM_USERS, "email", "username"), username).Scan(&email); err != nil {
		return "", fmt.Errorf("conn.QueryRow in GetEmailByUsername: %w", err)
	}
	return email, nil
}

func GetActiveStatusByUsername(username string) (int, error) {
	var isActive int
	if err := conn.QueryRow(context.Background(), fmt.Sprintf(SQL_SELECT_FROM_USERS, "is_active", "username"), username).Scan(&isActive); err != nil {
		return 0, fmt.Errorf("conn.QueryRow in GetActiveStatusByUsername: %w", err)
	}
	return isActive, nil
}

func GetOrganisationNameAndUserIdByUsername(username string) (string, int, error) {
	var organisationName string
	var userId int
	if err := conn.QueryRow(context.Background(), SQL_SELECT_ORGANISATION_NAME_BY_USERNAME, username).Scan(&organisationName, &userId); err != nil {
		return "", 0, fmt.Errorf("conn.QueryRow in GetOrganisationNameAndUserIdByUsername: %w", err)
	}
	return organisationName, userId, nil
}

func GetUserGroupsByUsername(username string) (pgx.Rows, error) {
	rows, err := conn.Query(context.Background(), SQL_SELECT_USERGROUPS_BY_USERNAME, username)
	return rows, err
}

func GetUserGroupCount(usergroup string) (int, error) {
	var count int
	if err := conn.QueryRow(context.Background(), fmt.Sprintf(SQL_SELECT_FROM_USERGROUPS, "user_group"), usergroup).Scan(&count); err != nil {
		return 0, fmt.Errorf("conn.QueryRow in GetUserGroupCount: %w", err)
	}
	return count, nil
}

func GetOrganisationNameCount(organisationName string) (int, error) {
	var count int
	if err := conn.QueryRow(context.Background(), fmt.Sprintf(SQL_SELECT_FROM_ORGANISATIONS, "COUNT(*)", "organisation_name"), organisationName).Scan(&count); err != nil {
		return 0, fmt.Errorf("conn.QueryRow in GetOrganisationNameCount: %w", err)
	}
	return count, nil
}

func GetUsers() (pgx.Rows, error) {
	rows, err := conn.Query(context.Background(), SQL_SELECT_ALL_USERS)
	return rows, err
}
