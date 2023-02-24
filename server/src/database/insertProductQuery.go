package database

var (
	SQL_INSERT_INTO_PRODUCTS      = `INSERT INTO products (product_name, product_description, product_sku, product_cost, added_date, updated_date)
									 VALUES($1, $2, $3, $4, now(), now()) RETURNING product_id;`

	SQL_INSERT_INTO_PRODUCT_SIZES_MAPPING = `INSERT into product_sizes_mapping (product_id, size_quantity, size_id)
							 VALUES($1, $2, (SELECT size_id FROM sizes WHERE size_name = $3));`

	SQL_INSERT_INTO_PRODUCT_USER_MAPPING = `INSERT INTO product_user_mapping
											(product_id, user_id, colour_id, category_id, brand_id, added_date, updated_date)
											VALUES ($1, $2, 
											(SELECT colour_id FROM colours WHERE colour_name = $3),
											(SELECT category_id FROM categories WHERE category_name = $4),
											(SELECT brand_id FROM brands WHERE brand_name = $5), now(), now());`
)

func InsertNewProduct(productName, productDescription, productSku string, productCost float32) (int, error) {
	var productId int
	err := db.QueryRow(SQL_INSERT_INTO_PRODUCTS, productName, productDescription, productSku, productCost).Scan(&productId)
	return productId, err
}

func InsertIntoProductSizesMapping(sizeName string, sizeQuantity, productId int) error {
	_, err := db.Exec(SQL_INSERT_INTO_PRODUCT_SIZES_MAPPING, productId, sizeQuantity, sizeName)
	return err
}

func InsertIntoProductUserMapping(productId, userId int, colourName, categoryName, brandName string) error {
	_, err := db.Exec(SQL_INSERT_INTO_PRODUCT_USER_MAPPING, productId, userId, colourName, categoryName, brandName)
	return err
}