package dbrepo

import "context"

func (m *TestDBRepo) UpdateUserTransaction(ctx context.Context, username, password, email, organisationName string, isActive int, userGroups []string) error {
	return nil
}
