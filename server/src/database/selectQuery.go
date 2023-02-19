package database

import (
	"database/sql"
	"fmt"
)

var (
	SQL_SELECT_FROM_USERS = "SELECT %s FROM users WHERE %s = $1;"
	SQL_SELECT_ALL_FROM_USERS = "SELECT username, password, email, user_group, company_name, is_active, added_date, updated_date FROM users;"
	SQL_SELECT_ALL_FROM_USERS_BY_USERNAME = "SELECT username, password, email, user_group, company_name, is_active, added_date, updated_date FROM users WHERE username = $1;"
	SQL_SELECT_FROM_ORGANISATIONS = "SELECT %s FROM organisations WHERE %s = $1;"
	SQL_SELECT_FROM_USERGROUPS = "SELECT COUNT(*) FROM user_groups WHERE %s = $1;"
	SQL_SELECT_USERGROUPS_BY_USERNAME = `SELECT ug.user_group FROM user_groups ug
										 LEFT JOIN user_group_mapping ugm 
										 ON ugm.user_group_id = ug.user_group_id 
										 WHERE ugm.user_id = (SELECT user_id FROM users WHERE username = $1);`
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

func GetUsername(username string) bool {
	row := db.QueryRow(fmt.Sprintf(SQL_SELECT_FROM_USERS, "username", "username"), username)
	return row.Scan() != sql.ErrNoRows
}

func GetEmail(email string) bool {
	row := db.QueryRow(fmt.Sprintf(SQL_SELECT_FROM_USERS, "email", "email"), email)
	return row.Scan() != sql.ErrNoRows
}

func GetOrganisationName(organisationName string) bool {
	row := db.QueryRow(fmt.Sprintf(SQL_SELECT_FROM_ORGANISATIONS, "organisation_name", "organisation_name"), organisationName)
	return row.Scan() != sql.ErrNoRows
}

func GetPasswordByUsername(username string) (string, error) {
	var password string
	row := db.QueryRow(fmt.Sprintf(SQL_SELECT_FROM_USERS, "password", "username"), username)
	err := row.Scan(&password)
	return password, err
}

func GetEmailByUsername(username string) (string, error) {
    var email string
    row := db.QueryRow(fmt.Sprintf(SQL_SELECT_FROM_USERS, "email", "username"), username)
    err := row.Scan(&email)
    return email, err
}

func GetActiveStatusByUsername(username string) (int, error) {
	var isActive int
	row := db.QueryRow(fmt.Sprintf(SQL_SELECT_FROM_USERS, "is_active", "username"), username)
	err := row.Scan(&isActive)
	return isActive, err
}

func GetCompanyNameByUsername(username string) (string, error) {
	var companyName string
	row := db.QueryRow(fmt.Sprintf(SQL_SELECT_FROM_USERS, "company_name", "username"), username)
	err := row.Scan(&companyName)
	return companyName, err
}

func GetUserGroupsByUsername(username string) (*sql.Rows, error) {
	rows, err := db.Query(SQL_SELECT_USERGROUPS_BY_USERNAME, username)
	return rows, err
}

// func GetUserGroupByUsername(username string) (string, error) {
// 	var userGroup string
// 	row := db.QueryRow(fmt.Sprintf(SQL_SELECT_FROM_USERS, "user_group", "username"), username)
// 	err := row.Scan(&userGroup)
// 	return userGroup, err
// }

func GetUserGroupCount(usergroup string) (int, error) {
	var count int
	row := db.QueryRow(fmt.Sprintf(SQL_SELECT_FROM_USERGROUPS, "user_group"), usergroup)
	err := row.Scan(&count)
	return count, err
}

func GetOrganisationNameCount(organisationName string) (int, error) {
	var count int
	row := db.QueryRow(fmt.Sprintf(SQL_SELECT_FROM_ORGANISATIONS, "COUNT(*)", "organisation_name"), organisationName)
	err := row.Scan(&count)
	return count, err
}

func GetUsers() (*sql.Rows, error) {
	row, err := db.Query(SQL_SELECT_ALL_FROM_USERS)
	return row, err
}

func GetUserByUsername(username string) *sql.Row {
	row := db.QueryRow(SQL_SELECT_ALL_FROM_USERS_BY_USERNAME, username)
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
