package repository

import (
	"context"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type DatabaseRepo interface {
	Connection() *pgxpool.Pool
	CheckUserGroup(ctx context.Context, username string, userGroups ...string) error

	GetCountByUsername(ctx context.Context, username string) (int, error)
	GetCountByEmail(ctx context.Context, email string) (int, error)
	GetCountByOrganisationName(ctx context.Context, organisationName string) (int, error)

	GetPasswordByUsername(ctx context.Context, username string) (string, error)
	GetIsActiveByUsername(ctx context.Context, username string) (int, error)
	GetUserGroupsByUsername(ctx context.Context, username string) (pgx.Rows, error)

	SignUpTransaction(ctx context.Context, username, password, email, organisationName, userGroup string, isActive int) error

	InsertIntoOrganisations(ctx context.Context, organisationName string) error
}
