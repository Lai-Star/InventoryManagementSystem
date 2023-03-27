package dbrepo

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type TestDBRepo struct{}

// Mock Database (not an actual connection)
func (m *TestDBRepo) Connection() *pgx.Conn {
	return nil
}

func (m *TestDBRepo) GetCountByUsername(ctx context.Context, username string) (int, error) {
	return 0, nil
}

func (m *TestDBRepo) GetCountByEmail(ctx context.Context, email string) (int, error) {
	return 0, nil
}

func (m *TestDBRepo) GetPasswordByUsername(ctx context.Context, username string) (string, error) {
	return "", nil
}

func (m *TestDBRepo) GetIsActiveByUsername(ctx context.Context, username string) (int, error) {
	return 0, nil
}

func (m *TestDBRepo) SignUpTransaction(ctx context.Context, username, password, email, organisationName, userGroup string, isActive int) error {
	return nil
}
