package dbrepo

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/LeonLow97/inventory-management-system-golang-react-postgresql/utils"
	"github.com/jackc/pgx/v4"
	"github.com/lib/pq"
)

var (
	SQL_UPDATE_USERS = "UPDATE users SET " +
		"password = COALESCE(NULLIF($2, ''), password), " +
		"email = COALESCE(NULLIF($3, ''), email), " +
		"is_active = $4 " +
		"WHERE username = $1 RETURNING user_id;"

	SQL_UPDATE_USER_ORGANISATION_MAPPING = "UPDATE user_organisation_mapping SET " +
		"organisation_id = ( " +
		"SELECT organisation_id FROM organisations WHERE organisation_name = $2) " +
		"WHERE user_id = $1;"

	SQL_UPDATE_USER_GROUP_MAPPING = "WITH new_groups AS (SELECT user_group_id FROM user_groups WHERE user_group = ANY($2::text[])), " +
		"deleted_groups AS (DELETE FROM user_group_mapping WHERE user_id = $1 AND user_group_id NOT IN " +
		"(SELECT user_group_id FROM new_groups)) " +
		"INSERT INTO user_group_mapping (user_id, user_group_id) SELECT $1, ng.user_group_id " +
		"FROM new_groups ng WHERE NOT EXISTS ( " +
		"SELECT user_group_id FROM user_group_mapping ugm WHERE ugm.user_id = $1 AND ugm.user_group_id = ng.user_group_id);"
)

func (m *PostgresDBRepo) UpdateUsers(ctx context.Context, username, password, email string, isActive int) (int, error) {
	var userId int
	if err := m.DB.QueryRow(ctx, SQL_UPDATE_USERS, username, password, email, isActive).Scan(&userId); err != nil {
		log.Println("QueryRow failed in UpdateUsers:", err)
		return 0, err
	}
	return userId, nil
}

func (m *PostgresDBRepo) UpdateUserOrganisationMapping(ctx context.Context, userId int, organisationName string) error {
	if _, err := m.DB.Exec(ctx, SQL_UPDATE_USER_ORGANISATION_MAPPING, userId, organisationName); err != nil {
		log.Println("QueryRow failed in UpdateUserOrganisationMapping:", err)
		return err
	}
	return nil
}

func (m *PostgresDBRepo) UpdateUserGroupMapping(ctx context.Context, userId int, userGroups []string) error {
	args := []interface{}{userId, pq.Array(userGroups)}
	if _, err := m.DB.Exec(ctx, SQL_UPDATE_USER_GROUP_MAPPING, args...); err != nil {
		log.Println("QueryRow failed in UpdateUserGroupMapping:", err)
		return err
	}
	return nil
}

func (m *PostgresDBRepo) UpdateUserTransaction(ctx context.Context, username, password, email, organisationName string, isActive int, userGroups []string) error {

	// Setting timeout context of 1 minutes
	ctx, cancel := context.WithTimeout(ctx, 1*time.Minute)
	defer cancel()

	tx, err := m.DB.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		log.Println("BeginTx failed in UpdateUserTransaction:", err)
		return utils.ApiError{Err: "Internal Server Error", Status: http.StatusInternalServerError}
	}

	defer func() {
		_ = tx.Rollback(ctx)
	}()

	var userId int

	if err := tx.QueryRow(ctx, SQL_UPDATE_USERS, username, password, email, isActive).Scan(&userId); err != nil {
		log.Println("QueryRow failed in UpdateUserTransaction SQL_UPDATE_USERS:", err)
		return utils.ApiError{Err: "Internal Server Error", Status: http.StatusInternalServerError}
	}

	if _, err := tx.Exec(ctx, SQL_UPDATE_USER_ORGANISATION_MAPPING, userId, organisationName); err != nil {
		log.Println("Exec failed in UpdateUserTransaction SQL_UPDATE_USER_ORGANISATION_MAPPING:", err)
		return utils.ApiError{Err: "Internal Server Error", Status: http.StatusInternalServerError}
	}

	if len(userGroups) > 0 {
		args := []interface{}{userId, pq.Array(userGroups)}
		if _, err := tx.Exec(ctx, SQL_UPDATE_USER_GROUP_MAPPING, args...); err != nil {
			log.Println("Exec failed in UpdateUserTransaction SQL_UPDATE_USERS:", err)
			return utils.ApiError{Err: "Internal Server Error", Status: http.StatusInternalServerError}
		}
	}

	if err := tx.Commit(ctx); err != nil {
		log.Println("Commit failed in UpdateUserTransaction:", err)
		return utils.ApiError{Err: "Internal Server Error", Status: http.StatusInternalServerError}
	}

	return nil

}
