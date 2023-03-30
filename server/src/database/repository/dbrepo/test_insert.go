package dbrepo

import "context"

func (m *TestDBRepo) InsertIntoOrganisations(ctx context.Context, organisationName string) error {
	return nil
}

func (m *TestDBRepo) InsertIntoUserGroups(ctx context.Context, userGroup, description string) error {
	return nil
}
