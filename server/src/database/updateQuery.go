package database

import (
	"fmt"
	"log"

	"github.com/lib/pq"
)

var (
	SQL_UPDATE_USERS = "UPDATE users SET " + 
					   "password = COALESCE(NULLIF($2, ''), password), " + 
					   "email = COALESCE(NULLIF($3, ''), email), " + 
					   "is_active = $4 " + 
					   "WHERE username = $1 RETURNING user_id;"
	SQL_UPDATE_USER_ORGANISATION_MAPPING = "UPDATE user_organisation_mapping SET " + 
										   "organisation_id = ( " + 
										   "SELECT organisation_id FROM organisations WHERE organisation_name = $2) " + 
										   "WHERE user_id = $1;"
	SQL_UPDATE_USER_GROUP_MAPPING = "WITH new_groups AS (SELECT user_group_id FROM user_groups WHERE user_group = ANY($2::text[])), " +
									"deleted_groups AS (DELETE FROM user_group_mapping WHERE user_id = $1 AND user_group_id NOT IN " +
									"(SELECT user_group_id FROM new_groups)) " +
									"INSERT INTO user_group_mapping (user_id, user_group_id) SELECT $1, ng.user_group_id " +
									"FROM new_groups ng WHERE NOT EXISTS ( " +
									"SELECT user_group_id FROM user_group_mapping ugm WHERE ugm.user_id = $1 AND ugm.user_group_id = ng.user_group_id);"
)

var (
	SQL_UPDATE_PRODUCTS = "UPDATE products SET %s = $1, %s = $2, %s = $3, %s = $4, %s = $5, %s = $6, %s = $7, updated_date = now() WHERE %s = $8;"
)

func UpdateUsers(username, password, email string, isActive int) (int, error) {
	var userId int
	err := db.QueryRow(SQL_UPDATE_USERS, username, password, email, isActive).Scan(&userId)
	return userId, err
}

func UpdateUserOrganisationMapping(userId int, organisationName string) error {
	_, err := db.Exec(SQL_UPDATE_USER_ORGANISATION_MAPPING, userId, organisationName)
	return err
}

func UpdateUserGroupMapping(userId int, userGroups []string) error {
	args := []interface{}{userId, pq.Array(userGroups)}
	_, err := db.Exec(SQL_UPDATE_USER_GROUP_MAPPING, args...)
	return err
}

func UpdateProduct(product_id int, product_name, product_description, product_sku, product_colour, product_category, product_brand string, product_cost float32) error  {
	query := fmt.Sprintf(SQL_UPDATE_PRODUCTS, "product_name", "product_description", "product_sku", "product_colour", "product_category", "product_brand", "product_cost", "product_id")
	_, err := db.Exec(query, product_name, product_description, product_sku, product_colour, product_category, product_brand, product_cost, product_id)
	if err != nil {
		log.Println("Error in Updating product: ", err)
	}
	return err
}

