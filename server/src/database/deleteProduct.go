package database

var (
	SQL_DELETE_FROM_PRODUCTS = `DELETE FROM products WHERE product_id = $1;`
)

func DeleteProduct(productId int) error {
	_, err := db.Exec(SQL_DELETE_FROM_PRODUCTS, productId)
	return err
}