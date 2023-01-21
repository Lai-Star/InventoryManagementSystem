package handlers_products

import (
	"fmt"
	"net/http"

	"github.com/LeonLow97/inventory-management-system-golang-react-postgresql/database"
	"github.com/LeonLow97/inventory-management-system-golang-react-postgresql/utils"
)

type ProductJson struct {
	ProductName        string `json:"product_name"`
	ProductDescription string `json:"product_description"`
	ProductSku         string `json:"product_sku"`
	ProductColour      string `json:"product_colour"`
	ProductCategory    string `json:"product_category"`
	ProductBrand       string `json:"product_brand"`
	ProductCost        float32 `json:"product_cost"`
}

func ProductsFormValdiation(w http.ResponseWriter, product ProductJson) bool {

	productSku := product.ProductSku

	// Check Product Sku to see if it exists in database
	isExistProductSku := database.ProductSkuExists(productSku)
	fmt.Println(isExistProductSku)
	if isExistProductSku {
		utils.ResponseJson(w, http.StatusBadRequest, "Product Sku already exists. Please try again.")
		return false
	}

	// Check Product Sku length to see if the length is less than 50
	isValidLengthProductSku := utils.CheckMaxLength(productSku, 50)
	if !isValidLengthProductSku {
		utils.ResponseJson(w, http.StatusBadRequest, "Product Sku should have less than or equal to 50 characters.")
		return false
	}

	return true
}



