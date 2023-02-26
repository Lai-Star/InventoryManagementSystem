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

	SQL_SELECT_ALL_FROM_PRODUCTS_BY_USERNAME = `SELECT p.product_id, p.product_name, p.product_description, p.product_sku, uc.colour_name, uctg.category_name,
									ub.brand_name, p.product_cost, us.size_name, upsm.size_quantity FROM products p
									INNER JOIN product_user_mapping pum ON p.product_id = pum.product_id
									LEFT JOIN user_product_sizes_mapping upsm ON p.product_id = upsm.product_id
									LEFT JOIN user_brands ub ON pum.brand_id = ub.brand_id
									LEFT JOIN user_colours uc ON pum.colour_id = uc.colour_id
									LEFT JOIN user_categories uctg ON pum.category_id = uctg.category_id
									LEFT JOIN user_sizes us ON upsm.size_id = us.size_id
									WHERE pum.user_id = $1
									ORDER BY p.added_date ASC`

	SQL_SELECT_ALL_FROM_PRODUCTS_BY_ORGANISATION = `SELECT p.product_id, p.product_name, p.product_description, p.product_sku, oc.colour_name, octg.category_name,
										ob.brand_name, p.product_cost, os.size_name, opsm.size_quantity FROM products p
										INNER JOIN product_organisation_mapping pom ON p.product_id = pom.product_id
										LEFT JOIN organisation_product_sizes_mapping opsm ON p.product_id = opsm.product_id
										LEFT JOIN organisation_brands ob ON pom.brand_id = ob.brand_id
										LEFT JOIN organisation_colours oc ON pom.colour_id = oc.colour_id
										LEFT JOIN organisation_categories octg ON pom.category_id = octg.category_id
										LEFT JOIN organisation_sizes os ON opsm.size_id = os.size_id
										WHERE pom.organisation_id = (SELECT organisation_id from organisations WHERE organisation_name = $1)
										ORDER BY p.added_date ASC;`

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

	SQL_SELECT_COUNT_BY_USERID_AND_PRODUCTID = `SELECT COUNT(*), p.product_sku FROM product_user_mapping pum
												JOIN products p ON pum.product_id = p.product_id
												WHERE pum.user_id = $1 AND pum.product_id = $2
												GROUP BY p.product_id;`
	SQL_SELECT_COUNT_BY_ORGANISATION_AND_PRODUCTID = `SELECT COUNT(*), p.product_sku FROM product_organisation_mapping pom
													JOIN products p ON pom.product_id = p.product_id
													WHERE pom.organisation_id = (SELECT organisation_id FROM organisations WHERE organisation_name = $1)
													AND pom.product_id = $2 GROUP BY p.product_id;`
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

func GetProductsByUsername(userId int) (*sql.Rows, error) {
	rows, err := db.Query(SQL_SELECT_ALL_FROM_PRODUCTS_BY_USERNAME, userId)
	return rows, err
}

func GetProductsByOrganisation(organisationName string) (*sql.Rows, error) {
	rows, err := db.Query(SQL_SELECT_ALL_FROM_PRODUCTS_BY_ORGANISATION, organisationName)
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

func GetCountByUserIdAndProductId(userId, productId int) (int, string, error) {
	var count int
	var currentProductSku string
	err := db.QueryRow(SQL_SELECT_COUNT_BY_USERID_AND_PRODUCTID, userId, productId).Scan(&count, &currentProductSku)
	return count, currentProductSku, err
}

func GetCountByOrganisationAndProductId(organisationName string, productId int) (int, string, error) {
	var count int
	var currentProductSku string
	err := db.QueryRow(SQL_SELECT_COUNT_BY_ORGANISATION_AND_PRODUCTID, organisationName, productId).Scan(&count, &currentProductSku)
	return count, currentProductSku, err
}

