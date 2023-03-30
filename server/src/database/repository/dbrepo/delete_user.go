package dbrepo

import (
	"context"
	"fmt"
)

var (
	SQL_DELETE_FROM_USERS_BY_ID = `DELETE FROM users WHERE user_id = 
																(SELECT user_id FROM users WHERE username = $1);`
)

func (m *PostgresDBRepo) DeleteUserByID(username string) error {
	if _, err := m.DB.Exec(context.Background(), SQL_DELETE_FROM_USERS_BY_ID, username); err != nil {
		return fmt.Errorf("m.DB.Exec in DeleteUserByID: %w", err)
	}
	return nil
}
