package database

import (
	"database/sql"
	"fmt"
	"log"
)

var (
	SQL_SELECT_FROM_ACCOUNTS = "SELECT %s FROM accounts WHERE %s = $1;"
	SQL_SELECT_ALL_FROM_ACCOUNTS = "SELECT username, password, email, user_group, company_name, is_active, added_date, updated_date FROM accounts;"
	SQL_SELECT_ALL_FROM_ACCOUNTS_BY_USERNAME = "SELECT username, password, email, user_group, company_name, is_active, added_date, updated_date FROM accounts WHERE username = $1;"
)

var (
	SQL_SELECT_FROM_PRODUCTS = "SELECT %s FROM products WHERE %s = $1;"
	SQL_SELECT_ALL_FROM_PRODUCTS = "SELECT p.product_name, p.product_description, p.product_sku, p.product_colour, p.product_category, p.product_brand, p.product_cost, s.size_name, s.size_quantity " +
									 "FROM products p " + 
									 "LEFT JOIN product_sizes ps ON p.product_id = ps.product_id " +
									 "LEFT JOIN sizes s ON s.size_id = ps.size_id;"
	SQL_SELECT_ALL_FROM_PRODUCTS_BY_PRODUCTID = "SELECT product_name, product_description, product_sku, product_colour, product_category, product_brand, product_cost" + 
													"FROM products WHERE product_id = $1;"
)

func UsernameExists(username string) bool {
	row := db.QueryRow(fmt.Sprintf(SQL_SELECT_FROM_ACCOUNTS, "username", "username"), username)
	return row.Scan() != sql.ErrNoRows
}

func EmailExists(email string) bool {
	row := db.QueryRow(fmt.Sprintf(SQL_SELECT_FROM_ACCOUNTS, "email", "email"), email)
	return row.Scan() != sql.ErrNoRows
}

func GetPasswordFromDB(username string) (string, error) {
	var password string
	row := db.QueryRow(fmt.Sprintf(SQL_SELECT_FROM_ACCOUNTS, "password", "username"), username)
	err := row.Scan(&password)
	if err != nil {
		log.Println("Error scanning when getting password from database: ", err)
	}
	return password, nil
}

func GetEmailFromDB(username string) (string, error) {
    var email string
    row := db.QueryRow(fmt.Sprintf(SQL_SELECT_FROM_ACCOUNTS, "email", "username"), username)
    err := row.Scan(&email)
    if err != nil {
		log.Println("Error scanning when getting email from database: ", err)
	}
    return email, nil
}

func GetActiveStatusFromDB(username string) (int, error) {
	var isActive int
	row := db.QueryRow(fmt.Sprintf(SQL_SELECT_FROM_ACCOUNTS, "is_active", "username"), username)
	err := row.Scan(&isActive)
	if err != nil {
		log.Println("Error scanning when getting isActive status from database: ", err)
	}
	return isActive, nil
}

func GetCompanyNameFromDB(username string) (string, error) {
	var companyName string
	row := db.QueryRow(fmt.Sprintf(SQL_SELECT_FROM_ACCOUNTS, "company_name", "username"), username)
	err := row.Scan(&companyName)
	if err != nil {
		log.Println("Error scanning when getting companyName from database: ", err)
	}
	return companyName, err
}

func GetUserGroupFromDB(username string) (string, error) {
	var userGroup string
	row := db.QueryRow(fmt.Sprintf(SQL_SELECT_FROM_ACCOUNTS, "user_group", "username"), username)
	err := row.Scan(&userGroup)
	if err != nil {
		log.Println("Error scanning when getting user group from database: ", err)
	}
	return userGroup, nil
}

func GetUsers() (*sql.Rows, error) {
	row, err := db.Query(SQL_SELECT_ALL_FROM_ACCOUNTS)
	return row, err
}

func GetUserByUsername(username string) *sql.Row {
	row := db.QueryRow(SQL_SELECT_ALL_FROM_ACCOUNTS_BY_USERNAME, username)
	return row
}

func ProductIdExists(product_id int) bool {
	row := db.QueryRow(fmt.Sprintf(SQL_SELECT_FROM_PRODUCTS, "product_id", "product_id"), product_id)
	return row.Scan() != sql.ErrNoRows
}

func ProductSkuExistsByUsername(product_sku, username string) bool {

	row := db.QueryRow(fmt.Sprintf(SQL_SELECT_FROM_PRODUCTS, "product_sku", "product_sku"), product_sku)
	return row.Scan() != sql.ErrNoRows
}

func GetProducts() (*sql.Rows, error) {
	rows, err := db.Query(SQL_SELECT_ALL_FROM_PRODUCTS)
	return rows, err
}

func GetProductByProductId(product_id int) *sql.Row {
	row := db.QueryRow(SQL_SELECT_ALL_FROM_PRODUCTS_BY_PRODUCTID, product_id)
	return row
}
