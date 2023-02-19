package database

import (
	"fmt"
	"log"
)

var (
	SQL_DELETE_FROM_USERS = "DELETE FROM users WHERE %s = $1;"
)

var (
	SQL_DELETE_FROM_PRODUCTS = "DELETE FROM products WHERE %s = $1;"
)

func DeleteUserFromUsers(username string) error {
	_, err := db.Exec(fmt.Sprintf(SQL_DELETE_FROM_USERS, "username"), username)
	return err
}

func DeleteProductFromProducts(product_id int) error {
	_, err := db.Exec(fmt.Sprintf(SQL_DELETE_FROM_PRODUCTS, "product_id"), product_id)
	if err != nil {
		log.Println("Internal Server Error deleting product from products: ", err)
	}
	return err
}