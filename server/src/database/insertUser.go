package database

import (
	"context"
	"fmt"
)

var (
	SQL_INSERT_INTO_USERS = "INSERT INTO users (username, password, email, is_active, added_date, updated_date)" +
		"VALUES ($1, $2, $3, $4, now(), now()) RETURNING user_id;"

	SQL_INSERT_INTO_USER_ORGANISATION_MAPPING = "INSERT INTO user_organisation_mapping (user_id, organisation_id) " +
		"SELECT $1, organisation_id " +
		"FROM organisations " +
		"WHERE organisation_name = $2;"

	SQL_INSERT_INTO_USER_GROUP_MAPPING = "INSERT INTO user_group_mapping (user_id, user_group_id) " +
		"SELECT $1, user_group_id " +
		"FROM user_groups " +
		"WHERE user_group = $2;"
		
	SQL_INSERT_INTO_ORGANISATIONS = "INSERT INTO organisations (organisation_name, added_date, updated_date) VALUES ($1, now(), now());"
	SQL_INSERT_INTO_USER_GROUPS   = "INSERT INTO user_groups (user_group, description, added_date, updated_date) VALUES ($1, $2, now(), now());"
)

func InsertNewUser(username, password, email string, isActive int) (int, error) {
	var userId int
	if err := conn.QueryRow(context.Background(), username, password, email, isActive).Scan(&userId); err != nil {
		return 0, fmt.Errorf("conn.QueryRow in InsertNewUser: %w", err)
	}
	return userId, nil
}

func InsertIntoUserOrganisationMapping(userId int, organisationName string) error {
	if _, err := conn.Exec(context.Background(), SQL_INSERT_INTO_USER_ORGANISATION_MAPPING, userId, organisationName); err != nil {
		return fmt.Errorf("conn.Exec in InsertIntoUserOrganisationMapping: %w", err)
	}
	return nil
}

func InsertIntoUserGroupMapping(userId int, userGroup string) error {
	if _, err := conn.Exec(context.Background(), SQL_INSERT_INTO_USER_GROUP_MAPPING, userId, userGroup); err != nil {
		return fmt.Errorf("conn.Exec in InsertIntoUserGroupMapping: %w", err)
	}
	return nil
}

func InsertIntoOrganisations(organisationName string) error {
	if _, err := conn.Exec(context.Background(), SQL_INSERT_INTO_ORGANISATIONS, organisationName); err != nil {
		return fmt.Errorf("conn.Exec in InsertIntoOrganisations: %w", err)
	}
	return nil
}

func InsertIntoUserGroups(userGroup, description string) error {
	if _, err := conn.Exec(context.Background(), SQL_INSERT_INTO_USER_GROUPS, userGroup, description); err != nil {
		return fmt.Errorf("conn.Exec in InsertIntoUserGroups: %w", err)
	}
	return nil
}
