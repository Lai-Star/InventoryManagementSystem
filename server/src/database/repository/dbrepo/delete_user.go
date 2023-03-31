package dbrepo

import (
	"context"
	"log"
)

var (
	SQL_DELETE_FROM_USERS_BY_ID = `DELETE FROM users WHERE user_id = 
																(SELECT user_id FROM users WHERE username = $1);`
)

func (m *PostgresDBRepo) DeleteUserByID(ctx context.Context, username string) error {
	if _, err := m.DB.Exec(ctx, SQL_DELETE_FROM_USERS_BY_ID, username); err != nil {
		log.Println("m.DB.Exec failed in DeleteUserByID:", err)
		return err
	}
	return nil
}
