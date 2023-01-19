package handlers_products

type ProductJson struct {
	ProductName   string `json:"product_name"`
	ProductSku    string `json:"product_sku"`
	ProductColour string `json:"product_colour"`
	TotalQuantity int `json:"total_quantity"`
	Xxs           int `json:"XXS"`
	Xs            int `json:"XS"`
	S             int `json:"S"`
	M             int `json:"M"`
	L             int `json:"L"`
	Xl            int `json:"XL"`
	Xxl           int `json:"XXL"`
}




