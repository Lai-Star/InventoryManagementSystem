package dbrepo

import (
	"context"
	"fmt"
	"time"
)

var (
	SQL_INSERT_INTO_PRODUCTS = `INSERT INTO products (product_name, product_description, product_sku, product_cost, added_date, updated_date)
									 						VALUES($1, $2, $3, $4, $5, $6) RETURNING product_id;`

	SQL_INSERT_INTO_USER_PRODUCT_SIZES_MAPPING = `INSERT into user_product_sizes_mapping (product_id, size_quantity, size_id)
							 																	VALUES($1, $2, (SELECT size_id FROM user_sizes WHERE size_name = $3));`

	SQL_INSERT_INTO_ORGANISATION_PRODUCT_SIZES_MAPPING = `INSERT into organisation_product_sizes_mapping (product_id, size_quantity, size_id)
	VALUES($1, $2, (SELECT size_id FROM organisation_sizes WHERE size_name = $3));`

	SQL_INSERT_INTO_PRODUCT_USER_MAPPING = `INSERT INTO product_user_mapping
											(product_id, user_id, colour_id, category_id, brand_id, added_date, updated_date)
											VALUES ($1, $2, 
											(SELECT colour_id FROM user_colours WHERE colour_name = $3),
											(SELECT category_id FROM user_categories WHERE category_name = $4),
											(SELECT brand_id FROM user_brands WHERE brand_name = $5), now(), now());`

	SQL_INSERT_INTO_PRODUCT_ORGANISATION_MAPPING = `INSERT INTO product_organisation_mapping
													(product_id, organisation_id, colour_id, category_id, brand_id, added_date, updated_date)
													VALUES ($1,
													(SELECT organisation_id from organisations WHERE organisation_name = $2), 
													(SELECT colour_id FROM organisation_colours WHERE colour_name = $3),
													(SELECT category_id FROM organisation_categories WHERE category_name = $4),
													(SELECT brand_id FROM organisation_brands WHERE brand_name = $5), now(), now());`

	SQL_INSERT_INTO_USER_BRANDS         = `INSERT INTO user_brands (user_id, brand_name) VALUES ($1, $2);`
	SQL_INSERT_INTO_ORGANISATION_BRANDS = `INSERT INTO organisation_brands (organisation_id, brand_name) VALUES 
											((SELECT organisation_id from organisations WHERE organisation_name = $1), $2);`

	SQL_INSERT_INTO_USER_CATEGORIES         = `INSERT INTO user_categories (user_id, category_name) VALUES ($1, $2);`
	SQL_INSERT_INTO_ORGANISATION_CATEGORIES = `INSERT INTO organisation_categories (organisation_id, category_name) VALUES 
											((SELECT organisation_id from organisations WHERE organisation_name = $1), $2);`

	SQL_INSERT_INTO_USER_COLOURS         = `INSERT INTO user_colours (user_id, colour_name) VALUES ($1, $2);`
	SQL_INSERT_INTO_ORGANISATION_COLOURS = `INSERT INTO organisation_colours (organisation_id, colour_name) VALUES 
											((SELECT organisation_id from organisations WHERE organisation_name = $1), $2);`

	SQL_INSERT_INTO_USER_SIZES         = `INSERT INTO user_sizes (user_id, size_name) VALUES ($1, $2);`
	SQL_INSERT_INTO_ORGANISATION_SIZES = `INSERT INTO organisation_sizes (organisation_id, size_name) VALUES 
											((SELECT organisation_id from organisations WHERE organisation_name = $1), $2);`
)

func InsertNewProduct(productName, productDescription, productSku string, productCost float32) (int, error) {
	var productId int
	if err := conn.QueryRow(context.Background(), SQL_INSERT_INTO_PRODUCTS, productName, productDescription, productSku, productCost, time.Now(), time.Now()).Scan(&productId); err != nil {
		return 0, fmt.Errorf("conn.QueryRow in InsertNewProduct: %w", err)
	}
	return productId, nil
}

func InsertIntoUserProductSizesMapping(sizeName string, sizeQuantity, productId int) error {
	if _, err := conn.Exec(context.Background(), SQL_INSERT_INTO_USER_PRODUCT_SIZES_MAPPING, productId, sizeQuantity, sizeName); err != nil {
		return fmt.Errorf("conn.Exec in InsertIntoUserProductSizesMapping: %w", err)
	}
	return nil
}

func InsertIntoOrganisationProductSizesMapping(sizeName string, sizeQuantity, productId int) error {
	if _, err := conn.Exec(context.Background(), SQL_INSERT_INTO_ORGANISATION_PRODUCT_SIZES_MAPPING, productId, sizeQuantity, sizeName); err != nil {
		return fmt.Errorf("conn.Exec in InsertIntoOrganisationProductSizesMapping: %w", err)
	}
	return nil
}

func InsertIntoProductUserMapping(productId, userId int, productColour, productCategory, productBrand string) error {
	if _, err := conn.Exec(context.Background(), SQL_INSERT_INTO_PRODUCT_USER_MAPPING, productId, userId, productColour, productCategory, productBrand); err != nil {
		return fmt.Errorf("conn.Exec in InsertIntoProductUserMapping: %w", err)
	}
	return nil
}

func InsertIntoProductOrganisationMapping(productId int, organisationName string, productColour, productCategory, productBrand string) error {
	if _, err := conn.Exec(context.Background(), SQL_INSERT_INTO_PRODUCT_ORGANISATION_MAPPING, productId, organisationName, productColour, productCategory, productBrand); err != nil {
		return fmt.Errorf("conn.Exec in InsertIntoProductOrganisationMapping: %w", err)
	}
	return nil
}

func InsertIntoUserBrands(userId int, brandName string) error {
	if _, err := conn.Exec(context.Background(), SQL_INSERT_INTO_USER_BRANDS, userId, brandName); err != nil {
		return fmt.Errorf("conn.Exec in InsertIntoUserBrands: %w", err)
	}
	return nil
}

func InsertIntoOrganisationBrands(organisationName, brandName string) error {
	if _, err := conn.Exec(context.Background(), SQL_INSERT_INTO_ORGANISATION_BRANDS, organisationName, brandName); err != nil {
		return fmt.Errorf("conn.Exec in InsertIntoOrganisationBrands: %w", err)
	}
	return nil
}

func InsertIntoUserCategories(userId int, categoryName string) error {
	if _, err := conn.Exec(context.Background(), SQL_INSERT_INTO_USER_CATEGORIES, userId, categoryName); err != nil {
		return fmt.Errorf("conn.Exec in InsertIntoUserCategories: %w", err)
	}
	return nil
}

func InsertIntoOrganisationCategories(organisationName, categoryName string) error {
	if _, err := conn.Exec(context.Background(), SQL_INSERT_INTO_ORGANISATION_CATEGORIES, organisationName, categoryName); err != nil {
		return fmt.Errorf("conn.Exec in InsertIntoOrganisationCategories: %w", err)
	}
	return nil
}

func InsertIntoUserColours(userId int, colourName string) error {
	if _, err := conn.Exec(context.Background(), SQL_INSERT_INTO_USER_COLOURS, userId, colourName); err != nil {
		return fmt.Errorf("conn.Exec in InsertIntoUserColours: %w", err)
	}
	return nil
}

func InsertIntoOrganisationColours(organisationName, colourName string) error {
	if _, err := conn.Exec(context.Background(), SQL_INSERT_INTO_ORGANISATION_COLOURS, organisationName, colourName); err != nil {
		return fmt.Errorf("conn.Exec in InsertIntoOrganisationColours: %w", err)
	}
	return nil
}

func InsertIntoUserSizes(userId int, sizeName string) error {
	if _, err := conn.Exec(context.Background(), SQL_INSERT_INTO_USER_SIZES, userId, sizeName); err != nil {
		return fmt.Errorf("conn.Exec in InsertIntoUserSizes: %w", err)
	}
	return nil
}

func InsertIntoOrganisationSizes(organisationName, sizeName string) error {
	if _, err := conn.Exec(context.Background(), SQL_INSERT_INTO_ORGANISATION_SIZES, organisationName, sizeName); err != nil {
		return fmt.Errorf("conn.Exec in InsertIntoOrganisationSizes: %w", err)
	}
	return nil
}
