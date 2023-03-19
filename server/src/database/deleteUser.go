package database

import (
	"context"
	"fmt"
)

var (
	SQL_DELETE_FROM_USERS_BY_ID = `DELETE FROM users WHERE user_id = 
																(SELECT user_id FROM users WHERE username = $1);`
)

func DeleteUserByID(username string) error {
	if _, err := conn.Exec(context.Background(), SQL_DELETE_FROM_USERS_BY_ID, username); err != nil {
		return fmt.Errorf("conn.Exec in DeleteUserByID: %w", err)
	}
	return nil
}
