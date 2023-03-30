package dbrepo

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/LeonLow97/inventory-management-system-golang-react-postgresql/utils"
	"github.com/jackc/pgx/v4"
)

var (
	SQL_INSERT_INTO_USERS = "INSERT INTO users (username, password, email, is_active, added_date, updated_date)" +
		"VALUES ($1, $2, $3, $4, $5, $6) RETURNING user_id;"

	SQL_INSERT_INTO_USER_ORGANISATION_MAPPING = "INSERT INTO user_organisation_mapping (user_id, organisation_id) " +
		"SELECT $1, organisation_id " +
		"FROM organisations " +
		"WHERE organisation_name = $2;"

	SQL_INSERT_INTO_USER_GROUP_MAPPING = "INSERT INTO user_group_mapping (user_id, user_group_id) " +
		"SELECT $1, user_group_id " +
		"FROM user_groups " +
		"WHERE user_group = $2;"

	SQL_INSERT_INTO_ORGANISATIONS = "INSERT INTO organisations (organisation_name, added_date, updated_date) VALUES ($1, $2, $3);"
	SQL_INSERT_INTO_USER_GROUPS   = "INSERT INTO user_groups (user_group, description, added_date, updated_date) VALUES ($1, $2, now(), now());"
)

func (m *PostgresDBRepo) InsertIntoOrganisations(ctx context.Context, organisationName string) error {
	if _, err := m.DB.Exec(ctx, SQL_INSERT_INTO_ORGANISATIONS, organisationName, time.Now(), time.Now()); err != nil {
		log.Println("Exec failed in InsertIntoOrganisations:", err)
		return utils.ApiError{Err: "Internal Server Error", Status: http.StatusInternalServerError}
	}
	return nil
}

func (m *PostgresDBRepo) InsertIntoUserGroups(ctx context.Context, userGroup, description string) error {
	if _, err := m.DB.Exec(ctx, SQL_INSERT_INTO_USER_GROUPS, userGroup, description); err != nil {
		log.Println("Exec failed in InsertIntoUserGroups:", err)
		return utils.ApiError{Err: "Internal Server Error", Status: http.StatusInternalServerError}
	}
	return nil
}

func (m *PostgresDBRepo) InsertIntoUser(ctx context.Context, username, password, email string, isActive int) (int, error) {
	var userId int
	if err := m.DB.QueryRow(context.Background(), SQL_INSERT_INTO_USERS, username, password, email, isActive).Scan(&userId); err != nil {
		return 0, fmt.Errorf("m.DB.QueryRow in InsertNewUser: %w", err)
	}
	return userId, nil
}

func (m *PostgresDBRepo) InsertIntoUserOrganisationMapping(userId int, organisationName string) error {
	if _, err := m.DB.Exec(context.Background(), SQL_INSERT_INTO_USER_ORGANISATION_MAPPING, userId, organisationName); err != nil {
		return fmt.Errorf("m.DB.Exec in InsertIntoUserOrganisationMapping: %w", err)
	}
	return nil
}

func (m *PostgresDBRepo) InsertIntoUserGroupMapping(userId int, userGroup string) error {
	if _, err := m.DB.Exec(context.Background(), SQL_INSERT_INTO_USER_GROUP_MAPPING, userId, userGroup); err != nil {
		return fmt.Errorf("m.DB.Exec in InsertIntoUserGroupMapping: %w", err)
	}
	return nil
}

func (m *PostgresDBRepo) CreateUserTransaction(ctx context.Context, username, password, email, organisationName string, isActive int, userGroups ...string) error {
	// Setting timeout context of 1 minutes
	ctx, cancel := context.WithTimeout(ctx, 1*time.Minute)
	defer cancel()

	tx, err := m.DB.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		log.Println("BeginTx failed in CreateUserTransaction:", err)
		return utils.ApiError{Err: "Internal Server Error", Status: http.StatusInternalServerError}
	}

	defer func() {
		_ = tx.Rollback(ctx)
	}()

	var userId int

	if err := tx.QueryRow(ctx, SQL_INSERT_INTO_USERS, username, password, email, isActive, time.Now(), time.Now()).Scan(&userId); err != nil {
		log.Println("QueryRow failed in CreateUserTransaction SQL_INSERT_INTO_USERS:", err)
		return utils.ApiError{Err: "Internal Server Error", Status: http.StatusInternalServerError}
	}

	if _, err := tx.Exec(ctx, SQL_INSERT_INTO_USER_ORGANISATION_MAPPING, userId, organisationName); err != nil {
		log.Println("Exec failed in CreateUserTransaction SQL_INSERT_INTO_USER_ORGANISATION_MAPPING:", err)
		return utils.ApiError{Err: "Internal Server Error", Status: http.StatusInternalServerError}
	}

	for _, userGroup := range userGroups {
		if _, err := tx.Exec(ctx, SQL_INSERT_INTO_USER_GROUP_MAPPING, userId, userGroup); err != nil {
			log.Println("Exec failed in CreateUserTransaction SQL_INSERT_INTO_USER_GROUP_MAPPING:", err)
			return utils.ApiError{Err: "Internal Server Error", Status: http.StatusInternalServerError}
		}
	}

	if err := tx.Commit(ctx); err != nil {
		log.Println("Commit failed in CreateUserTransaction:", err)
		return utils.ApiError{Err: "Internal Server Error", Status: http.StatusInternalServerError}
	}

	return nil
}

func (m *PostgresDBRepo) SignUpTransaction(ctx context.Context, username, password, email, organisationName, userGroup string, isActive int) error {
	// Setting timeout context of 1 minutes
	ctx, cancel := context.WithTimeout(ctx, 1*time.Minute)
	defer cancel()

	tx, err := m.DB.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		log.Println("BeginTx failed in SignUpTransaction:", err)
		return utils.ApiError{Err: "Internal Server Error", Status: http.StatusInternalServerError}
	}

	defer func() {
		_ = tx.Rollback(ctx)
	}()

	var userId int

	if err := tx.QueryRow(ctx, SQL_INSERT_INTO_USERS, username, password, email, isActive, time.Now(), time.Now()).Scan(&userId); err != nil {
		log.Println("QueryRow failed in SignUpTransaction SQL_INSERT_INTO_USERS:", err)
		return utils.ApiError{Err: "Internal Server Error", Status: http.StatusInternalServerError}
	}

	if _, err := tx.Exec(ctx, SQL_INSERT_INTO_USER_ORGANISATION_MAPPING, userId, organisationName); err != nil {
		log.Println("Exec failed in SignUpTransaction SQL_INSERT_INTO_USER_ORGANISATION_MAPPING:", err)
		return utils.ApiError{Err: "Internal Server Error", Status: http.StatusInternalServerError}
	}

	if _, err := tx.Exec(ctx, SQL_INSERT_INTO_USER_GROUP_MAPPING, userId, userGroup); err != nil {
		log.Println("Exec failed in SignUpTransaction SQL_INSERT_INTO_USER_GROUP_MAPPING:", err)
		return utils.ApiError{Err: "Internal Server Error", Status: http.StatusInternalServerError}
	}

	if err := tx.Commit(ctx); err != nil {
		log.Println("Commit failed in SignUpTransaction:", err)
		return utils.ApiError{Err: "Internal Server Error", Status: http.StatusInternalServerError}
	}

	return nil
}
