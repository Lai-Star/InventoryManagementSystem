package dbrepo

import "context"

func (m *TestDBRepo) SignUpTransaction(ctx context.Context, username, password, email, organisationName, userGroup string, isActive int) error {
	return nil
}

func (m *TestDBRepo) CreateUserTransaction(ctx context.Context, username, password, email, organisationName string, isActive int, userGroups ...string) error {
	return nil
}
