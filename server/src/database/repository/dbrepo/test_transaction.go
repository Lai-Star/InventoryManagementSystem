package dbrepo

import "context"

func (m *TestDBRepo) SignUpTransaction(ctx context.Context, username, password, email, organisationName, userGroup string, isActive int) error {
	return nil
}
