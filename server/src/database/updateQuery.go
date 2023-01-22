package database

import (
	"fmt"
	"log"
)

var (
	SQL_UPDATE_ACCOUNTS = "UPDATE accounts SET %s = $1, %s = $2, %s = $3, %s = $4, %s = $5, updated_date = now() WHERE %s = $6;"
)

var (
	SQL_UPDATE_PRODUCTS = "UPDATE products SET %s = $1, %s = $2, %s = $3, %s = $4, %s = $5, %s = $6, %s = $7, updated_date = now() WHERE %s = $8;"
)

func AdminUpdateUser(username, password, email, userGroup, companyName string, isActive int) error {
	query := fmt.Sprintf(SQL_UPDATE_ACCOUNTS, "password", "email", "user_group", "company_name", "is_active" , "username")
	_, err := db.Exec(query, password, email, userGroup, companyName, isActive, username)
	if err != nil {
		log.Println("Error in Admin updating user password: ", err)
	}
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

