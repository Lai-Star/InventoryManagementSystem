package database

import (
	"context"
	"fmt"
)

var (
	SQL_UPDATE_PRODUCTS_BY_PRODUCTID = `UPDATE products SET product_name = $1, product_description = $2, product_sku = $3, product_cost = $4, updated_date = NOW()
							WHERE product_id = $5;`

	SQL_UPDATE_PRODUCT_USER_MAPPING = `UPDATE product_user_mapping AS pum
										SET colour_id = uc.colour_id, category_id = ucat.category_id, brand_id = ub.brand_id, updated_date = NOW()
										FROM (SELECT user_id, product_id FROM product_user_mapping WHERE user_id = $1 AND product_id = $2) AS pum_sub
										INNER JOIN user_colours AS uc ON uc.user_id = pum_sub.user_id AND uc.colour_name = COALESCE(NULLIF($3,''), uc.colour_name)
										INNER JOIN user_categories AS ucat ON ucat.user_id = pum_sub.user_id AND ucat.category_name = COALESCE(NULLIF($4,''), ucat.category_name)
										INNER JOIN user_brands AS ub ON ub.user_id = pum_sub.user_id AND ub.brand_name = COALESCE(NULLIF($5,''), ub.brand_name)
										WHERE pum.product_id = pum_sub.product_id;`

	SQL_UPDATE_PRODUCT_ORGANISATION_MAPPING = `UPDATE product_organisation_mapping AS pom
												SET colour_id = oc.colour_id, category_id = ocat.category_id, brand_id = ob.brand_id, updated_date = NOW()
												FROM (SELECT organisation_id, product_id FROM product_organisation_mapping WHERE organisation_id = 
													(SELECT organisation_id FROM organisations WHERE organisation_name = $1) AND product_id = $2) AS pom_sub
												INNER JOIN organisation_colours oc ON oc.organisation_id = pom_sub.organisation_id AND oc.colour_name = COALESCE(NULLIF($3,''), oc.colour_name)
												INNER JOIN organisation_categories AS ocat ON ocat.organisation_id = pom_sub.organisation_id AND ocat.category_name = COALESCE(NULLIF($4,''), ocat.category_name)
												INNER JOIN organisation_brands AS ob ON ob.organisation_id = pom_sub.organisation_id AND ob.brand_name = COALESCE(NULLIF($5,''), ob.brand_name)
												WHERE pom.product_id = pom_sub.product_id;`

	SQL_UPDATE_USER_PRODUCT_SIZES_MAPPING = `UPDATE user_product_sizes_mapping SET
												size_quantity = $1 WHERE product_id = $2
												AND size_id = (SELECT size_id FROM user_sizes WHERE size_name = $3);`

	SQL_UPDATE_ORGANISATION_PRODUCT_SIZES_MAPPING = `UPDATE organisation_product_sizes_mapping SET
														size_quantity = $1 WHERE product_id = $2
														AND size_id = (SELECT size_id FROM organisation_sizes WHERE size_name = $3);`
)

func UpdateProductsByProductID(productName, productDesc, productSku string, productCost float32, productId int) error {
	if _, err := conn.Exec(context.Background(), SQL_UPDATE_PRODUCTS_BY_PRODUCTID, productName, productDesc, productSku, productCost, productId); err != nil {
		return fmt.Errorf("conn.Exec in UpdateProductsByProductID: %w", err)
	}
	return nil
}

func UpdateProductUserMapping(userId, productId int, colourName, categoryName, brandName string) error {
	if _, err := conn.Exec(context.Background(), SQL_UPDATE_PRODUCT_USER_MAPPING, userId, productId, colourName, categoryName, brandName); err != nil {
		return fmt.Errorf("conn.Exec in UpdateProductUserMapping: %w", err)
	}
	return nil
}

func UpdateProductOrganisationMapping(productId int, organisationName, colourName, categoryName, brandName string) error {
	if _, err := conn.Exec(context.Background(), SQL_UPDATE_PRODUCT_ORGANISATION_MAPPING, organisationName, productId, colourName, categoryName, brandName); err != nil {
		return fmt.Errorf("conn.Exec in UpdateProductOrganisationMapping: %w", err)
	}
	return nil
}

func UpdateUserProductSizesMapping(sizeQuantity, productId int, sizeName string) error {
	if _, err := conn.Exec(context.Background(), SQL_UPDATE_USER_PRODUCT_SIZES_MAPPING, sizeQuantity, productId, sizeName); err != nil {
		return fmt.Errorf("conn.Exec in UpdateUserProductSizesMapping: %w", err)
	}
	return nil
}

func UpdateOrganisationProductSizesMapping(sizeQuantity, productId int, sizeName string) error {
	if _, err := conn.Exec(context.Background(), SQL_UPDATE_ORGANISATION_PRODUCT_SIZES_MAPPING, sizeQuantity, productId, sizeName); err != nil {
		return fmt.Errorf("conn.Exec in UpdateOrganisationProductSizesMapping: %w", err)
	}
	return nil
}
