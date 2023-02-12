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
)

var (
	SQL_INSERT_INTO_PRODUCTS = "INSERT INTO products (product_name, product_description, product_sku, product_colour, product_category, product_brand, product_cost, added_date, updated_date) VALUES ($1, $2, $3, $4, $5, $6, $7, now(), now()) RETURNING product_id"
	SQL_INSERT_INTO_SIZES = "INSERT INTO sizes (size_name, size_quantity, added_date, updated_date) VALUES ($1, $2, now(), now()) RETURNING size_id;"
	SQL_INSERT_INTO_PRODUCT_SIZES = "INSERT INTO product_sizes (product_id, size_id) VALUES ($1, $2);"
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

func AdminInsertNewUser(username, password, email, user_group, company_name string, isActive int) error {
	_, err := db.Exec(SQL_INSERT_INTO_USERS, username, password, email, user_group, company_name, isActive)
	if err != nil {
		log.Println("Error admin inserting new user to database: ", err)
	}
	return err
}

func InsertNewProduct(product_name, product_description, product_sku, product_colour, product_category, product_brand string, product_cost float32) (int32, error) {
	var productId int32
	err := db.QueryRow(SQL_INSERT_INTO_PRODUCTS, product_name, product_description, product_sku, product_colour, product_category, product_brand, product_cost).Scan(&productId)
	if err != nil {
		log.Println("Error inserting new product: ", err)
	}
	return productId, err
}

func InsertNewSize(size_name string, size_quantity int) (int32, error) {
	var sizeId int32
	err := db.QueryRow(SQL_INSERT_INTO_SIZES, size_name, size_quantity).Scan(&sizeId)
	if err != nil {
		log.Println("Error inserting new size: ", err)
	}
	return sizeId, err
}

func InsertNewProductSizes(product_id, size_id int32) error {
	_, err := db.Exec(SQL_INSERT_INTO_PRODUCT_SIZES, product_id, size_id)
	if err != nil {
		log.Println("Error inserting productid and sizeid to product_sizes table: ", err)
	}
	return err
}