package database

import (
	"log"
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
	SQL_INSERT_INTO_USER_GROUPS = "INSERT INTO user_groups (user_group, description, added_date, updated_date) VALUES ($1, $2, now(), now());"
)

func InsertNewUser(username, password, email string, isActive int) (int, error) {
	var user_id int
	err := db.QueryRow(SQL_INSERT_INTO_USERS, username, password, email, isActive).Scan(&user_id)
	if err != nil {
		log.Println("Error inserting new user to database: ", err)
		return 0, err
	}
	return user_id, nil
}

func InsertIntoUserOrganisationMapping(userId int, organisationName string) error {
	_, err := db.Exec(SQL_INSERT_INTO_USER_ORGANISATION_MAPPING, userId, organisationName)
	if err != nil {
		log.Println("Error inserting into user_organisation_mapping table: ", err)
	}
	return err
}

func InsertIntoUserGroupMapping(userId int, userGroup string) error {
	_, err := db.Exec(SQL_INSERT_INTO_USER_GROUP_MAPPING, userId, userGroup)
	if err != nil {
		log.Println("Error inserting into user_group_mapping table: ", err)
	}
	return err
}

func InsertIntoOrganisations(organisationName string) error {
	_, err := db.Exec(SQL_INSERT_INTO_ORGANISATIONS, organisationName)
	return err
}

func InsertIntoUserGroups(userGroup, description string) error {
	_, err := db.Exec(SQL_INSERT_INTO_USER_GROUPS, userGroup, description)
	return err
}