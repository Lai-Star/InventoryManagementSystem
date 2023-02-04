package database

import (
	"log"
)

var (
	SQL_INSERT_INTO_ACCOUNTS = "INSERT INTO accounts (username, password, email, user_group, company_name, is_active, added_date, updated_date) VALUES ($1, $2, $3, $4, $5, $6, now(), now());"
)

var (
	SQL_INSERT_INTO_PRODUCTS = "INSERT INTO products (product_name, product_description, product_sku, product_colour, product_category, product_brand, product_cost, added_date, updated_date) VALUES ($1, $2, $3, $4, $5, $6, $7, now(), now()) RETURNING product_id"
	SQL_INSERT_INTO_SIZES = "INSERT INTO sizes (size_name, size_quantity, added_date, updated_date) VALUES ($1, $2, now(), now()) RETURNING size_id;"
	SQL_INSERT_INTO_PRODUCT_SIZES = "INSERT INTO product_sizes (product_id, size_id) VALUES ($1, $2);"
)

func InsertNewUser(username, password, email, user_group, company_name string, isActive int) error {
	_, err := db.Exec(SQL_INSERT_INTO_ACCOUNTS, username, password, email, user_group, company_name, isActive)
	if err != nil {
		log.Println("Error inserting new user to database: ", err)
	}
	return err
}

func AdminInsertNewUser(username, password, email, user_group, company_name string, isActive int) error {
	_, err := db.Exec(SQL_INSERT_INTO_ACCOUNTS, username, password, email, user_group, company_name, isActive)
	if err != nil {
		log.Println("Error admin inserting new user to database: ", err)
	}
	return err
}

func InsertNewProduct(product_name, product_description, product_sku, product_colour, product_category, product_brand string, product_cost float32) (error, int32) {
	var product_id int32
	err := db.QueryRow(SQL_INSERT_INTO_PRODUCTS, product_name, product_description, product_sku, product_colour, product_category, product_brand, product_cost).Scan(&product_id)
	if err != nil {
		log.Println("Error inserting new product: ", err)
	}
	return err, product_id
}

func InsertNewSize(size_name string, size_quantity int) (error, int32) {
	var size_id int32
	err := db.QueryRow(SQL_INSERT_INTO_SIZES, size_name, size_quantity).Scan(&size_id)
	if err != nil {
		log.Println("Error inserting new size: ", err)
	}
	return err, size_id
}

func InsertNewProductSizes(product_id, size_id int32) error {
	_, err := db.Exec(SQL_INSERT_INTO_PRODUCT_SIZES, product_id, size_id)
	if err != nil {
		log.Println("Error inserting productid and sizeid to product_sizes table: ", err)
	}
	return err
}