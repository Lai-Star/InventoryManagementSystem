package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type DatabaseRepo interface {
	Connection() *pgx.Conn

	GetCountByUsername(ctx context.Context, username string) (int, error)
	GetCountByEmail(ctx context.Context, email string) (int, error)
	GetPasswordByUsername(ctx context.Context, username string) (string, error)
	GetIsActiveByUsername(ctx context.Context, username string) (int, error)

	SignUpTransaction(ctx context.Context, username, password, email, organisationName, userGroup string, isActive int) error
}
