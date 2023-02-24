package database

import (
	"database/sql"
	"fmt"
)

var (
	SQL_SELECT_FROM_PRODUCTS = "SELECT %s FROM products WHERE %s = $1;"

	SQL_SELECT_PRODUCT_SKU_BY_USERNAME = `SELECT %s FROM products p
									  INNER JOIN product_user_mapping pum ON p.product_id = pum.product_id
									  INNER JOIN users u ON pum.user_id = u.user_id
									  WHERE u.username = $1 AND p.product_sku = $2;`

	SQL_SELECT_PRODUCT_SKU_BY_ORGANISATION = `SELECT %s FROM products p
											INNER JOIN product_organisation_mapping pom ON p.product_id = pom.product_id
											INNER JOIN organisation o ON pon.organisation_id = o.organisation_id
											WHERE o.organisation_name = $1 AND p.product_sku = $2;`

	SQL_SELECT_ALL_FROM_PRODUCTS = "SELECT p.product_name, p.product_description, p.product_sku, p.product_colour, p.product_category, p.product_brand, p.product_cost, s.size_name, s.size_quantity " +
									 "FROM products p " + 
									 "LEFT JOIN product_sizes ps ON p.product_id = ps.product_id " +
									 "LEFT JOIN sizes s ON s.size_id = ps.size_id;"

	SQL_SELECT_ALL_FROM_PRODUCTS_BY_PRODUCTID = "SELECT product_name, product_description, product_sku, product_colour, product_category, product_brand, product_cost" + 
													"FROM products WHERE product_id = $1;"
)

func GetProductSkuCountByUsername(username, productSku string) (int, error) {
	var count int
	row := db.QueryRow(fmt.Sprintf(SQL_SELECT_PRODUCT_SKU_BY_USERNAME, "COUNT(*)"), username, productSku)
	err := row.Scan(&count)
	return count, err
}

func GetProductSkuCountByOrganisation(organisationName, productSku string) (int, error) {
	var count int
	row := db.QueryRow(fmt.Sprintf(SQL_SELECT_PRODUCT_SKU_BY_ORGANISATION, "COUNT(*)"), organisationName, productSku)
	err := row.Scan(&count)
	return count, err
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