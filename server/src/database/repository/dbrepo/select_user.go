package dbrepo

import (
	"context"
	"fmt"
	"log"

	"github.com/LeonLow97/inventory-management-system-golang-react-postgresql/utils"
	"github.com/jackc/pgx/v4"
)

var (
	SQL_GET_COUNT_BY_USERNAME   = "SELECT COUNT(*) FROM users WHERE username = $1;"
	SQL_GET_COUNT_BY_EMAIL      = "SELECT COUNT(*) FROM users WHERE email = $1;"
	SQL_GET_COUNT_BY_ORG_NAME   = "SELECT COUNT(*) FROM organisations WHERE organisation_name = $1;"
	SQL_GET_COUNT_BY_USER_GROUP = "SELECT COUNT(*) FROM user_groups WHERE user_group = $1;"

	SQL_GET_PASSWORD_BY_USERNAME = "SELECT password FROM users WHERE username = $1;"
	SQL_GET_ISACTIVE_BY_USERNAME = "SELECT is_active FROM users WHERE username = $1;"

	SQL_GET_USERGROUPS_BY_USERNAME = `SELECT ug.user_group FROM user_groups ug
	LEFT JOIN user_group_mapping ugm 
	ON ugm.user_group_id = ug.user_group_id 
	WHERE ugm.user_id = (SELECT user_id FROM users WHERE username = $1);`

	SQL_SELECT_FROM_USERS = "SELECT %s FROM users WHERE %s = $1;"

	SQL_SELECT_ORGANISATION_NAME_BY_USERNAME = `SELECT o.organisation_name, u.user_id FROM organisations o
												INNER JOIN user_organisation_mapping uom
												ON o.organisation_id = uom.organisation_id
												INNER JOIN users u
												ON uom.user_id = u.user_id
												WHERE u.username = $1;`

	SQL_SELECT_ALL_USERS = `SELECT u.user_id, u.username, u.email, u.is_active, o.organisation_name, ug.user_group, u.added_date, u.updated_date FROM users u 
							LEFT JOIN user_organisation_mapping uom ON u.user_id = uom.user_id
							LEFT JOIN organisations o ON uom.organisation_id = o.organisation_id
							LEFT JOIN user_group_mapping ugm ON u.user_id = ugm.user_id
							LEFT JOIN user_groups ug ON ugm.user_group_id = ug.user_group_id
							ORDER BY user_id ASC;`
)

func (m *PostgresDBRepo) GetCountByUsername(ctx context.Context, username string) (int, error) {
	var count int
	err := m.DB.QueryRow(ctx, SQL_GET_COUNT_BY_USERNAME, username).Scan(&count)
	if err != nil {
		log.Println("QueryRow failed at GetCountByUsername:", err)
		return 0, err
	}
	return count, nil
}

func (m *PostgresDBRepo) GetCountByEmail(ctx context.Context, email string) (int, error) {
	var count int
	err := m.DB.QueryRow(ctx, SQL_GET_COUNT_BY_EMAIL, email).Scan(&count)
	if err != nil {
		log.Println("QueryRow failed at GetCountByEmail:", err)
		return 0, err
	}
	return count, nil
}

func (m *PostgresDBRepo) GetCountByOrganisationName(ctx context.Context, organisationName string) (int, error) {
	var count int
	if err := m.DB.QueryRow(ctx, SQL_GET_COUNT_BY_ORG_NAME, organisationName).Scan(&count); err != nil {
		log.Println("QueryRow failed at GetCountByOrganisationName:", err)
		return 0, err
	}
	return count, nil
}

func (m *PostgresDBRepo) GetPasswordByUsername(ctx context.Context, username string) (string, error) {
	var password string
	if err := m.DB.QueryRow(ctx, SQL_GET_PASSWORD_BY_USERNAME, username).Scan(&password); err != nil {
		log.Println("QueryRow failed at GetPasswordByUsername:", err)
		return "", err
	}
	return password, nil
}

func (m *PostgresDBRepo) GetIsActiveByUsername(ctx context.Context, username string) (int, error) {
	var isActive int
	if err := m.DB.QueryRow(ctx, SQL_GET_ISACTIVE_BY_USERNAME, username).Scan(&isActive); err != nil {
		return 0, err
	}
	return isActive, nil
}

func (m *PostgresDBRepo) GetUserGroupsByUsername(ctx context.Context, username string, userGroups ...string) (bool, error) {
	rows, err := m.DB.Query(ctx, SQL_GET_USERGROUPS_BY_USERNAME, username)
	if err != nil {
		log.Println("Query failed at GetUserGroupsByUsername:", err)
		return false, err
	}

	var userGroup string
	if rows != nil {
		for rows.Next() {
			if err = rows.Scan(&userGroup); err != nil {
				return false, err
			}

			if utils.Contains(userGroups, userGroup) {
				return true, nil
			}
		}
	}
	return false, nil
}

func (m *PostgresDBRepo) GetCountByUserGroup(ctx context.Context, userGroup string) (int, error) {
	var count int
	if err := m.DB.QueryRow(ctx, SQL_GET_COUNT_BY_USER_GROUP, userGroup).Scan(&count); err != nil {
		log.Println("Query failed at GetUserGroupCount:", err)
		return 0, err
	}
	return count, nil
}

// func (m *PostgresDBRepo) GetOrganisationName(organisationName string) bool {
// 	row := m.DB.QueryRow(context.Background(), fmt.Sprintf(SQL_SELECT_FROM_ORGANISATIONS, "organisation_name", "organisation_name"), organisationName)
// 	return row.Scan() != pgx.ErrNoRows
// }

func (m *PostgresDBRepo) GetEmailByUsername(username string) (string, error) {
	var email string
	if err := m.DB.QueryRow(context.Background(), fmt.Sprintf(SQL_SELECT_FROM_USERS, "email", "username"), username).Scan(&email); err != nil {
		return "", fmt.Errorf("m.DB.QueryRow in GetEmailByUsername: %w", err)
	}
	return email, nil
}

func (m *PostgresDBRepo) GetOrganisationNameAndUserIdByUsername(username string) (string, int, error) {
	var organisationName string
	var userId int
	if err := m.DB.QueryRow(context.Background(), SQL_SELECT_ORGANISATION_NAME_BY_USERNAME, username).Scan(&organisationName, &userId); err != nil {
		return "", 0, fmt.Errorf("m.DB.QueryRow in GetOrganisationNameAndUserIdByUsername: %w", err)
	}
	return organisationName, userId, nil
}

func (m *PostgresDBRepo) GetUsers() (pgx.Rows, error) {
	rows, err := m.DB.Query(context.Background(), SQL_SELECT_ALL_USERS)
	return rows, err
}
