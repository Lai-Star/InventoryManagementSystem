package dbrepo

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"sort"

	"github.com/LeonLow97/inventory-management-system-golang-react-postgresql/api/handlers"
	"github.com/LeonLow97/inventory-management-system-golang-react-postgresql/utils"
)

var (
	SQL_GET_COUNT_BY_USERNAME   = "SELECT COUNT(*) FROM users WHERE username = $1;"
	SQL_GET_COUNT_BY_EMAIL      = "SELECT COUNT(*) FROM users WHERE email = $1;"
	SQL_GET_COUNT_BY_ORG_NAME   = "SELECT COUNT(*) FROM organisations WHERE organisation_name = $1;"
	SQL_GET_COUNT_BY_USER_GROUP = "SELECT COUNT(*) FROM user_groups WHERE user_group = $1;"

	SQL_GET_PASSWORD_BY_USERNAME = "SELECT password FROM users WHERE username = $1;"
	SQL_GET_EMAIL_BY_USERNAME    = "SELECT email FROM users WHERE username = $1;"
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

	SQL_GET_ALL_USERS = `SELECT u.user_id, u.username, u.email, u.is_active, o.organisation_name, ug.user_group, u.added_date, u.updated_date FROM users u 
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
	fmt.Println("Email: ", email, "Count:", count)
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

func (m *PostgresDBRepo) GetEmailByUsername(ctx context.Context, username string) (string, error) {
	var email string
	if err := m.DB.QueryRow(ctx, SQL_GET_EMAIL_BY_USERNAME, username).Scan(&email); err != nil {
		log.Println("QueryRow failed at GetEmailByUsername:", err)
		return "", err
	}
	return email, nil
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

// Checking for duplicates for CreateUser
func (m *PostgresDBRepo) CheckDuplicatesAndExistingFieldsForCreateUser(ctx context.Context, username, email, organisationName string, userGroups ...string) error {

	usernameCount, err := m.GetCountByUsername(ctx, username)
	if err != nil {
		log.Println("Error in m.GetCountByUsername:", err)
		return utils.ApiError{Err: "Internal Server Error", Status: http.StatusInternalServerError}
	}
	if usernameCount == 1 {
		return utils.ApiError{Err: "Username " + username + " has already been taken. Please try again", Status: http.StatusBadRequest}
	}

	emailCount, err := m.GetCountByEmail(ctx, email)
	if err != nil {
		log.Println("Error in m.GetCountByEmail:", err)
		return utils.ApiError{Err: "Internal Server Error", Status: http.StatusInternalServerError}
	}
	if emailCount == 1 {
		return utils.ApiError{Err: "Email address " + email + " has already been taken. Please try again", Status: http.StatusBadRequest}
	}

	organisationCount, err := m.GetCountByOrganisationName(ctx, organisationName)
	if err != nil {
		log.Println("Error in m.GetCountByOrganisationName:", err)
		return utils.ApiError{Err: "Internal Server Error", Status: http.StatusInternalServerError}
	}
	if organisationCount != 1 {
		return utils.ApiError{Err: "Organisation " + organisationName + " does not exist. Please try again", Status: http.StatusBadRequest}
	}

	for _, userGroup := range userGroups {
		userGroupCount, err := m.GetCountByUserGroup(ctx, userGroup)
		if err != nil {
			log.Println("Error in m.GetCountByUserGroup:", err)
			return utils.ApiError{Err: "Internal Server Error", Status: http.StatusInternalServerError}
		}
		if userGroupCount != 1 {
			return utils.ApiError{Err: "User Group " + userGroup + " does not exist. Please try again", Status: http.StatusBadRequest}
		}
	}

	return nil

}

// Checking for duplicates and existing fields for UpdateUser
func (m *PostgresDBRepo) CheckDuplicatesAndExistingFieldsForUpdateUser(ctx context.Context, username, email, organisationName string, userGroups ...string) error {
	usernameCount, err := m.GetCountByUsername(ctx, username)
	if err != nil {
		log.Println("Error in m.GetCountByUsername:", err)
		return utils.ApiError{Err: "Internal Server Error", Status: http.StatusInternalServerError}
	}
	if usernameCount == 0 {
		return utils.ApiError{Err: "Username " + username + " does not exist. Please try again", Status: http.StatusBadRequest}
	}

	if len(email) > 0 {
		dbEmail, err := m.GetEmailByUsername(ctx, username)
		if err != nil {
			log.Println("Error in m.GetEmailByUsername:", err)
			return utils.ApiError{Err: "Internal Server Error", Status: http.StatusInternalServerError}
		}
		if dbEmail != email {
			return utils.ApiError{Err: "Email address " + email + " has already been taken. Please try again", Status: http.StatusBadRequest}
		}
	}

	if len(organisationName) > 0 {
		organisationCount, err := m.GetCountByOrganisationName(ctx, organisationName)
		if err != nil {
			log.Println("Error in m.GetCountByOrganisationName:", err)
			return utils.ApiError{Err: "Internal Server Error", Status: http.StatusInternalServerError}
		}
		if organisationCount != 1 {
			return utils.ApiError{Err: "Organisation " + organisationName + " does not exist. Please try again", Status: http.StatusBadRequest}
		}
	}

	for _, userGroup := range userGroups {
		userGroupCount, err := m.GetCountByUserGroup(ctx, userGroup)
		if err != nil {
			log.Println("Error in m.GetCountByUserGroup:", err)
			return utils.ApiError{Err: "Internal Server Error", Status: http.StatusInternalServerError}
		}
		if userGroupCount != 1 {
			return utils.ApiError{Err: "User Group " + userGroup + " does not exist. Please try again", Status: http.StatusBadRequest}
		}
	}

	return nil

}

func (m *PostgresDBRepo) GetAllUsers(ctx context.Context, data []handlers.User, users map[int]handlers.User) ([]handlers.User, error) {

	// To handle nullable columns in a database table
	var username, email, organisationName, userGroup, addedDate, updatedDate sql.NullString
	var userId, isActive sql.NullInt16

	rows, err := m.DB.Query(context.Background(), SQL_GET_ALL_USERS)
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&userId, &username, &email, &isActive, &organisationName, &userGroup, &addedDate, &updatedDate)
		if err != nil {
			log.Println("rows.Scan failed in GetAllUsers:", err)
			return nil, utils.ApiError{Err: "Internal Server Error", Status: http.StatusInternalServerError}
		}

		// Check if user already exists in map
		if user, ok := users[int(userId.Int16)]; ok {
			// User already exists, append userGroup to UserGroup array
			user.UserGroup = append(user.UserGroup, userGroup.String)
			users[int(userId.Int16)] = user
		} else {
			// User does not exist in map, create a new User object
			user := handlers.User{
				UserId:           int(userId.Int16),
				Username:         username.String,
				Email:            email.String,
				IsActive:         int(isActive.Int16),
				OrganisationName: organisationName.String,
				UserGroup:        []string{userGroup.String},
				AddedDate:        addedDate.String,
				UpdatedDate:      updatedDate.String,
			}
			users[int(userId.Int16)] = user
		}
	}

	// Convert map to slice
	for _, user := range users {
		data = append(data, user)
	}

	// Sort users by UserId in ascending order
	sort.Slice(data, func(i, j int) bool {
		return data[i].UserId < data[j].UserId
	})

	return data, err
}

// func (m *PostgresDBRepo) GetOrganisationName(organisationName string) bool {
// 	row := m.DB.QueryRow(context.Background(), fmt.Sprintf(SQL_SELECT_FROM_ORGANISATIONS, "organisation_name", "organisation_name"), organisationName)
// 	return row.Scan() != pgx.ErrNoRows
// }

func (m *PostgresDBRepo) GetOrganisationNameAndUserIdByUsername(username string) (string, int, error) {
	var organisationName string
	var userId int
	if err := m.DB.QueryRow(context.Background(), SQL_SELECT_ORGANISATION_NAME_BY_USERNAME, username).Scan(&organisationName, &userId); err != nil {
		return "", 0, fmt.Errorf("m.DB.QueryRow in GetOrganisationNameAndUserIdByUsername: %w", err)
	}
	return organisationName, userId, nil
}
