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
											INNER JOIN organisations o ON pom.organisation_id = o.organisation_id
											WHERE o.organisation_name = $1 AND p.product_sku = $2;`

	SQL_SELECT_ALL_FROM_PRODUCTS = "SELECT p.product_name, p.product_description, p.product_sku, p.product_colour, p.product_category, p.product_brand, p.product_cost, s.size_name, s.size_quantity " +
									 "FROM products p " + 
									 "LEFT JOIN product_sizes ps ON p.product_id = ps.product_id " +
									 "LEFT JOIN %s s ON s.size_id = ps.size_id;"

	SQL_SELECT_ALL_FROM_PRODUCTS_BY_PRODUCTID = "SELECT product_name, product_description, product_sku, product_colour, product_category, product_brand, product_cost" + 
													"FROM products WHERE product_id = $1;"

	SQL_SELECT_COUNT_FROM_USER_BRANDS = `SELECT COUNT(*) FROM user_brands WHERE user_id = $1 AND brand_name = $2;`
	SQL_SELECT_COUNT_FROM_ORGANISATION_BRANDS = `SELECT COUNT(*) FROM organisation_brands ob
												INNER JOIN organisations o
												ON o.organisation_id = ob.organisation_id
												WHERE o.organisation_name = $1 AND ob.brand_name = $2;`

	SQL_SELECT_COUNT_FROM_USER_CATEGORIES = `SELECT COUNT(*) FROM user_categories WHERE user_id = $1 AND category_name = $2;`
	SQL_SELECT_COUNT_FROM_ORGANISATION_CATEGORIES = `SELECT COUNT(*) FROM organisation_categories oc
												INNER JOIN organisations o
												ON o.organisation_id = oc.organisation_id
												WHERE o.organisation_name = $1 AND oc.category_name = $2;`

	SQL_SELECT_COUNT_FROM_USER_COLOURS = `SELECT COUNT(*) FROM user_colours WHERE user_id = $1 AND colour_name = $2;`
	SQL_SELECT_COUNT_FROM_ORGANISATION_COLOURS = `SELECT COUNT(*) FROM organisation_colours oc
												INNER JOIN organisations o
												ON o.organisation_id = oc.organisation_id
												WHERE o.organisation_name = $1 AND oc.colour_name = $2;`

	SQL_SELECT_COUNT_FROM_USER_SIZES = `SELECT COUNT(*) FROM user_sizes WHERE user_id = $1 AND size_name = $2;`
	SQL_SELECT_COUNT_FROM_ORGANISATION_SIZES = `SELECT COUNT(*) FROM organisation_sizes os
												INNER JOIN organisations o
												ON o.organisation_id = os.organisation_id
												WHERE o.organisation_name = $1 AND os.size_name = $2;`
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

func GetProducts(tableName string) (*sql.Rows, error) {
	query := fmt.Sprintf(SQL_SELECT_ALL_FROM_PRODUCTS, tableName)
	rows, err := db.Query(query)
	return rows, err
}

func GetProductByProductId(product_id int) *sql.Row {
	row := db.QueryRow(SQL_SELECT_ALL_FROM_PRODUCTS_BY_PRODUCTID, product_id)
	return row
}

func GetBrandNameCountByUsername(userId int, brandName string) (int, error) {
	var count int
	err := db.QueryRow(SQL_SELECT_COUNT_FROM_USER_BRANDS, userId, brandName).Scan(&count)
	return count, err
}

func GetBrandNameCountByOrganisation(organisationName, brandName string) (int, error) {
	var count int
	err := db.QueryRow(SQL_SELECT_COUNT_FROM_ORGANISATION_BRANDS, organisationName, brandName).Scan(&count)
	return count, err
}

func GetCategoryNameCountByUsername(userId int, categoryName string) (int, error) {
	var count int
	err := db.QueryRow(SQL_SELECT_COUNT_FROM_USER_CATEGORIES, userId, categoryName).Scan(&count)
	return count, err
}

func GetCategoryNameCountByOrganisation(organisationName, categoryName string) (int, error) {
	var count int
	err := db.QueryRow(SQL_SELECT_COUNT_FROM_ORGANISATION_CATEGORIES, organisationName, categoryName).Scan(&count)
	return count, err
}

func GetColourNameCountByUsername(userId int, colourName string) (int, error) {
	var count int
	err := db.QueryRow(SQL_SELECT_COUNT_FROM_USER_COLOURS, userId, colourName).Scan(&count)
	return count, err
}

func GetColourNameCountByOrganisation(organisationName, colourName string) (int, error) {
	var count int
	err := db.QueryRow(SQL_SELECT_COUNT_FROM_ORGANISATION_COLOURS, organisationName, colourName).Scan(&count)
	return count, err
}

func GetSizeNameCountByUsername(userId int, sizeName string) (int, error) {
	var count int
	err := db.QueryRow(SQL_SELECT_COUNT_FROM_USER_SIZES, userId, sizeName).Scan(&count)
	return count, err
}

func GetSizeNameCountByOrganisation(organisationName, sizeName string) (int, error) {
	var count int
	err := db.QueryRow(SQL_SELECT_COUNT_FROM_ORGANISATION_SIZES, organisationName, sizeName).Scan(&count)
	return count, err
}

