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

func (m *TestDBRepo) CheckUserGroup(ctx context.Context, username string, userGroups ...string) error {
	return nil
}

func (m *TestDBRepo) GetCountByUsername(ctx context.Context, username string) (int, error) {
	if username == "lowjiewei" {
		return 1, nil
	}
	return 0, nil
}

func (m *TestDBRepo) GetCountByEmail(ctx context.Context, email string) (int, error) {
	return 0, nil
}

func (m *TestDBRepo) GetCountByOrganisationName(ctx context.Context, organisationName string) (int, error) {
	return 0, nil
}

func (m *TestDBRepo) GetPasswordByUsername(ctx context.Context, username string) (string, error) {
	if username == "lowjiewei" {
		return "$2a$10$dxMeJmuR2p2EhxhuZC8DIezEpjpzFG6tWI6IKzJczHSfwkbsYQaDm", nil
	}
	return "", nil
}

func (m *TestDBRepo) GetIsActiveByUsername(ctx context.Context, username string) (int, error) {
	if username == "lowjiewei" {
		return 1, nil
	}
	return 0, nil
}

func (m *TestDBRepo) GetUserGroupsByUsername(ctx context.Context, username string) (pgx.Rows, error) {
	return nil, nil
}

func (m *TestDBRepo) SignUpTransaction(ctx context.Context, username, password, email, organisationName, userGroup string, isActive int) error {
	return nil
}

func (m *TestDBRepo) InsertIntoOrganisations(ctx context.Context, organisationName string) error {
	return nil
}
