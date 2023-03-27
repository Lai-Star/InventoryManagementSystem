package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type DatabaseRepo interface {
	Connection() *pgx.Conn
	GetCountByUsername(username string) (int, error)
	GetCountByEmail(email string) (int, error)

	SignUpTransaction(ctx context.Context, username, password, email, organisationName, userGroup string, isActive int) error
}
