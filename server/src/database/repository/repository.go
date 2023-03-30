package repository

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

type DatabaseRepo interface {
	Connection() *pgxpool.Pool

	GetCountByUsername(ctx context.Context, username string) (int, error)
	GetCountByEmail(ctx context.Context, email string) (int, error)
	GetCountByOrganisationName(ctx context.Context, organisationName string) (int, error)
	GetCountByUserGroup(ctx context.Context, userGroup string) (int, error)

	GetPasswordByUsername(ctx context.Context, username string) (string, error)
	GetIsActiveByUsername(ctx context.Context, username string) (int, error)
	GetUserGroupsByUsername(ctx context.Context, username string, userGroups ...string) (bool, error)

	SignUpTransaction(ctx context.Context, username, password, email, organisationName, userGroup string, isActive int) error

	InsertIntoOrganisations(ctx context.Context, organisationName string) error
	InsertIntoUserGroups(ctx context.Context, userGroup, description string) error
}
