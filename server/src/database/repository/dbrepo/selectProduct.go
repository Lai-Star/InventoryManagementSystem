package dbrepo

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
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

	SQL_SELECT_COUNT_FROM_USER_BRANDS         = `SELECT COUNT(*) FROM user_brands WHERE user_id = $1 AND brand_name = $2;`
	SQL_SELECT_COUNT_FROM_ORGANISATION_BRANDS = `SELECT COUNT(*) FROM organisation_brands ob
												INNER JOIN organisations o
												ON o.organisation_id = ob.organisation_id
												WHERE o.organisation_name = $1 AND ob.brand_name = $2;`

	SQL_SELECT_COUNT_FROM_USER_CATEGORIES         = `SELECT COUNT(*) FROM user_categories WHERE user_id = $1 AND category_name = $2;`
	SQL_SELECT_COUNT_FROM_ORGANISATION_CATEGORIES = `SELECT COUNT(*) FROM organisation_categories oc
												INNER JOIN organisations o
												ON o.organisation_id = oc.organisation_id
												WHERE o.organisation_name = $1 AND oc.category_name = $2;`

	SQL_SELECT_COUNT_FROM_USER_COLOURS         = `SELECT COUNT(*) FROM user_colours WHERE user_id = $1 AND colour_name = $2;`
	SQL_SELECT_COUNT_FROM_ORGANISATION_COLOURS = `SELECT COUNT(*) FROM organisation_colours oc
												INNER JOIN organisations o
												ON o.organisation_id = oc.organisation_id
												WHERE o.organisation_name = $1 AND oc.colour_name = $2;`

	SQL_SELECT_COUNT_FROM_USER_SIZES         = `SELECT COUNT(*) FROM user_sizes WHERE user_id = $1 AND size_name = $2;`
	SQL_SELECT_COUNT_FROM_ORGANISATION_SIZES = `SELECT COUNT(*) FROM organisation_sizes os
												INNER JOIN organisations o
												ON o.organisation_id = os.organisation_id
												WHERE o.organisation_name = $1 AND os.size_name = $2;`

	SQL_SELECT_COUNT_SIZES_BY_USERID_AND_PRODUCTID = `SELECT COUNT(*) FROM user_product_sizes_mapping upsm
										INNER JOIN user_sizes us ON upsm.size_id = us.size_id
										ON upsm.size_id = us.size_id
										WHERE product_id = $1 AND us.size_name = $2 AND us.user_ud = $3;`
	SQL_SELECT_COUNT_SIZES_BY_ORGANISATIONID_AND_PRODUCTID = `SELECT COUNT(*) FROM organisation_product_sizes_mapping opsm
														INNER JOIN organisation_sizes os ON opsm.size_id = os.size_id
														WHERE product_id = $1 AND os.size_name = $2
														AND os.organisation_id = (SELECT organisation_id FROM organisations WHERE organisation_name = $3);`

	SQL_SELECT_COUNT_PRODUCTSKU_BY_USERID_AND_PRODUCTID = `SELECT COUNT(*), p.product_sku FROM product_user_mapping pum
												JOIN products p ON pum.product_id = p.product_id
												WHERE pum.user_id = $1 AND pum.product_id = $2
												GROUP BY p.product_id;`
	SQL_SELECT_COUNT_PRODUCTSKU_BY_ORGANISATION_AND_PRODUCTID = `SELECT COUNT(*), p.product_sku FROM product_organisation_mapping pom
													JOIN products p ON pom.product_id = p.product_id
													WHERE pom.organisation_id = (SELECT organisation_id FROM organisations WHERE organisation_name = $1)
													AND pom.product_id = $2 GROUP BY p.product_id;`

	SQL_SELECT_COUNT_BY_USERID_AND_PRODUCTID           = `SELECT COUNT(*) FROM product_user_mapping WHERE user_id = $1 AND product_id = $2;`
	SQL_SELECT_COUNT_BY_ORGANISATIONNAME_AND_PRODUCTID = `SELECT COUNT(*) FROM product_organisation_mapping 
															WHERE organisation_id = (SELECT organisation_id FROM organisations WHERE organisation_name = $1)
															AND product_id = $2;`
)

func (m *PostgresDBRepo) GetProductSkuCountByUsername(username, productSku string) (int, error) {
	var count int
	if err := m.DB.QueryRow(context.Background(), fmt.Sprintf(SQL_SELECT_PRODUCT_SKU_BY_USERNAME, "COUNT(*)"), username, productSku).Scan(&count); err != nil {
		return 0, fmt.Errorf("m.DB.QueryRow in GetProductSkuCountByUsername: %w", err)
	}
	return count, nil
}

func (m *PostgresDBRepo) GetProductSkuCountByOrganisation(organisationName, productSku string) (int, error) {
	var count int
	if err := m.DB.QueryRow(context.Background(), fmt.Sprintf(SQL_SELECT_PRODUCT_SKU_BY_ORGANISATION, "COUNT(*)"), organisationName, productSku).Scan(&count); err != nil {
		return 0, fmt.Errorf("m.DB.QueryRow in GetProductSkuCountByOrganisation: %w", err)
	}
	return count, nil
}

func (m *PostgresDBRepo) ProductSkuExistsByUsername(product_sku, username string) bool {
	row := m.DB.QueryRow(context.Background(), fmt.Sprintf(SQL_SELECT_FROM_PRODUCTS, "product_sku", "product_sku"), product_sku)
	return row.Scan() != pgx.ErrNoRows
}

func (m *PostgresDBRepo) GetProductsByUsername(userId int) (pgx.Rows, error) {
	rows, err := m.DB.Query(context.Background(), SQL_SELECT_ALL_FROM_PRODUCTS_BY_USERNAME, userId)
	if err != nil {
		return nil, fmt.Errorf("m.DB.Query in GetProductsByUsername: %w", err)
	}
	return rows, nil
}

func (m *PostgresDBRepo) GetProductsByOrganisation(organisationName string) (pgx.Rows, error) {
	rows, err := m.DB.Query(context.Background(), SQL_SELECT_ALL_FROM_PRODUCTS_BY_ORGANISATION, organisationName)
	if err != nil {
		return nil, fmt.Errorf("m.DB.Query in GetProductsByOrganisation: %w", err)
	}
	return rows, nil
}

func (m *PostgresDBRepo) GetProductByProductId(product_id int) pgx.Row {
	row := m.DB.QueryRow(context.Background(), SQL_SELECT_ALL_FROM_PRODUCTS_BY_PRODUCTID, product_id)
	return row
}

func (m *PostgresDBRepo) GetBrandNameCountByUsername(userId int, brandName string) (int, error) {
	var count int
	if err := m.DB.QueryRow(context.Background(), SQL_SELECT_COUNT_FROM_USER_BRANDS, userId, brandName).Scan(&count); err != nil {
		return 0, fmt.Errorf("m.DB.QueryRow in GetBrandNameCountByUsername: %w", err)
	}
	return count, nil
}

func (m *PostgresDBRepo) GetBrandNameCountByOrganisation(organisationName, brandName string) (int, error) {
	var count int
	if err := m.DB.QueryRow(context.Background(), SQL_SELECT_COUNT_FROM_ORGANISATION_BRANDS, organisationName, brandName).Scan(&count); err != nil {
		return 0, fmt.Errorf("m.DB.QueryRow in GetBrandNameCountByOrganisation: %w", err)
	}
	return count, nil
}

func (m *PostgresDBRepo) GetCategoryNameCountByUsername(userId int, categoryName string) (int, error) {
	var count int
	if err := m.DB.QueryRow(context.Background(), SQL_SELECT_COUNT_FROM_USER_CATEGORIES, userId, categoryName).Scan(&count); err != nil {
		return 0, fmt.Errorf("m.DB.QueryRow in GetCategoryNameCountByUsername: %w", err)
	}
	return count, nil
}

func (m *PostgresDBRepo) GetCategoryNameCountByOrganisation(organisationName, categoryName string) (int, error) {
	var count int
	if err := m.DB.QueryRow(context.Background(), SQL_SELECT_COUNT_FROM_ORGANISATION_CATEGORIES, organisationName, categoryName).Scan(&count); err != nil {
		return 0, fmt.Errorf("m.DB.QueryRow in GetCategoryNameCountByOrganisation: %w", err)
	}
	return count, nil
}

func (m *PostgresDBRepo) GetColourNameCountByUsername(userId int, colourName string) (int, error) {
	var count int
	if err := m.DB.QueryRow(context.Background(), SQL_SELECT_COUNT_FROM_USER_COLOURS, userId, colourName).Scan(&count); err != nil {
		return 0, fmt.Errorf("m.DB.QueryRow in GetColourNameCountByUsername: %w", err)
	}
	return count, nil
}

func (m *PostgresDBRepo) GetColourNameCountByOrganisation(organisationName, colourName string) (int, error) {
	var count int
	if err := m.DB.QueryRow(context.Background(), SQL_SELECT_COUNT_FROM_ORGANISATION_COLOURS, organisationName, colourName).Scan(&count); err != nil {
		return 0, fmt.Errorf("m.DB.QueryRow in GetColourNameCountByOrganisation: %w", err)
	}
	return count, nil
}

func (m *PostgresDBRepo) GetSizeNameCountByUsername(userId int, sizeName string) (int, error) {
	var count int
	if err := m.DB.QueryRow(context.Background(), SQL_SELECT_COUNT_FROM_USER_SIZES, userId, sizeName).Scan(&count); err != nil {
		return 0, fmt.Errorf("m.DB.QueryRow in GetSizeNameCountByUsername: %w", err)
	}
	return count, nil
}

func (m *PostgresDBRepo) GetSizeNameCountByOrganisation(organisationName, sizeName string) (int, error) {
	var count int
	if err := m.DB.QueryRow(context.Background(), SQL_SELECT_COUNT_FROM_ORGANISATION_SIZES, organisationName, sizeName).Scan(&count); err != nil {
		return 0, fmt.Errorf("m.DB.QueryRow in GetSizeNameCountByOrganisation: %w", err)
	}
	return count, nil
}

func (m *PostgresDBRepo) GetSizeNameCountByUserIdAndProductId(productId, userId int, sizeName string) (int, error) {
	var count int
	if err := m.DB.QueryRow(context.Background(), SQL_SELECT_COUNT_SIZES_BY_USERID_AND_PRODUCTID, productId, sizeName, userId).Scan(&count); err != nil {
		return 0, fmt.Errorf("m.DB.QueryRow in GetSizeNameCountByUserIdAndProductId: %w", err)
	}
	return count, nil
}

func (m *PostgresDBRepo) GetSizeNameCountByOrganisationIdAndProductId(productId int, sizeName, organisationName string) (int, error) {
	var count int
	if err := m.DB.QueryRow(context.Background(), SQL_SELECT_COUNT_SIZES_BY_ORGANISATIONID_AND_PRODUCTID, productId, sizeName, organisationName).Scan(&count); err != nil {
		return 0, fmt.Errorf("m.DB.QueryRow in GetSizeNameCountByOrganisationIdAndProductId: %w", err)
	}
	return count, nil
}

func (m *PostgresDBRepo) GetCountProductSkuByUserIdAndProductId(userId, productId int) (int, string, error) {
	var count int
	var currentProductSku string
	if err := m.DB.QueryRow(context.Background(), SQL_SELECT_COUNT_PRODUCTSKU_BY_USERID_AND_PRODUCTID, userId, productId).Scan(&count, &currentProductSku); err != nil {
		return 0, "", fmt.Errorf("m.DB.QueryRow in GetCountProductSkuByUserIdAndProductId: %w", err)
	}
	return count, currentProductSku, nil
}

func (m *PostgresDBRepo) GetCountProductSkuByOrganisationAndProductId(organisationName string, productId int) (int, string, error) {
	var count int
	var currentProductSku string
	if err := m.DB.QueryRow(context.Background(), SQL_SELECT_COUNT_PRODUCTSKU_BY_ORGANISATION_AND_PRODUCTID, organisationName, productId).Scan(&count, &currentProductSku); err != nil {
		return 0, "", fmt.Errorf("m.DB.QueryRow in GetCountProductSkuByOrganisationAndProductId: %w", err)
	}
	return count, currentProductSku, nil
}

func (m *PostgresDBRepo) GetCountByUserIdAndProductId(userId, productId int) (int, error) {
	var count int
	if err := m.DB.QueryRow(context.Background(), SQL_SELECT_COUNT_BY_USERID_AND_PRODUCTID, userId, productId).Scan(&count); err != nil {
		return 0, fmt.Errorf("m.DB.QueryRow in GetCountByUserIdAndProductId: %w", err)
	}
	return count, nil
}

func (m *PostgresDBRepo) GetCountByOrganisationNameAndProductId(organisationName string, productId int) (int, error) {
	var count int
	if err := m.DB.QueryRow(context.Background(), SQL_SELECT_COUNT_BY_ORGANISATIONNAME_AND_PRODUCTID, organisationName, productId).Scan(&count); err != nil {
		return 0, fmt.Errorf("m.DB.QueryRow in GetCountByOrganisationNameAndProductId: %w", err)
	}
	return count, nil
}
