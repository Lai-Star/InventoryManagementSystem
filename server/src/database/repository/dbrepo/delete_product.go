package dbrepo

import (
	"context"
	"fmt"
)

var (
	SQL_DELETE_FROM_PRODUCTS_BY_ID = `DELETE FROM products WHERE product_id = $1;`
)

func (m *PostgresDBRepo) DeleteProductByID(productId int) error {
	if _, err := m.DB.Exec(context.Background(), SQL_DELETE_FROM_PRODUCTS_BY_ID, productId); err != nil {
		return fmt.Errorf("m.DB.Exec in DeleteProductByID: %w", err)
	}
	return nil
}
