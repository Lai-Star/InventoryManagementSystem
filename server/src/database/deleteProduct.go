package database

import (
	"context"
	"fmt"
)

var (
	SQL_DELETE_FROM_PRODUCTS_BY_ID = `DELETE FROM products WHERE product_id = $1;`
)

func DeleteProductByID(productId int) error {
	if _, err := conn.Exec(context.Background(), SQL_DELETE_FROM_PRODUCTS_BY_ID, productId); err != nil {
		return fmt.Errorf("conn.Exec in DeleteProductByID: %w", err)
	}
	return nil
}
