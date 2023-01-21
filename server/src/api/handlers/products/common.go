package handlers_products

import (
	"net/http"

	handlers_user "github.com/LeonLow97/inventory-management-system-golang-react-postgresql/api/handlers/user"
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

type DeleteProductJson struct {
	ProductSku string `json:"product_sku"`
}

func ProductsFormValdiation(w http.ResponseWriter, product ProductJson) bool {

	isValidProductSku := CheckProductSkuFormat(w, product.ProductSku)
	if !isValidProductSku {return false}
	
	return true
}

func CheckProductsUserGroup(w http.ResponseWriter, req *http.Request) bool {
	// CheckUserGroup: IMS User and Operations
	handlers_user.RetrieveIssuer(w, req)
	checkUserGroupIMSUser := utils.CheckUserGroup(w.Header().Get("username"), "IMS User")
	checkUserGroupOperations := utils.CheckUserGroup(w.Header().Get("username"), "Operations")
	if !checkUserGroupIMSUser || !checkUserGroupOperations {
		utils.ResponseJson(w, http.StatusForbidden, "Access Denied: You do not have permission to access this resource.")
		return false
	}
	return true
}

func CheckProductSkuFormat(w http.ResponseWriter, productSku string) bool {

	// Check Product Sku length to see if the length is less than 50
	isValidLengthProductSku := utils.CheckMaxLength(productSku, 50)
	if !isValidLengthProductSku {
		utils.ResponseJson(w, http.StatusBadRequest, "Product Sku should have less than or equal to 50 characters.")
		return false
	}

	return true
}



