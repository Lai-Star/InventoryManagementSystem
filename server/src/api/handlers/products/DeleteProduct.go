package handlers_products

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/LeonLow97/inventory-management-system-golang-react-postgresql/utils"
)

func DeleteProduct(w http.ResponseWriter, req *http.Request) {
	
	// Set Headers
	w.Header().Set("Content-Type", "application/json")
	var deleteProduct DeleteProductJson

	// Reading the request body and UnMarshal the body to the DeleteProductJson struct
	bs, _ := io.ReadAll(req.Body);
	if err := json.Unmarshal(bs, &deleteProduct); err != nil {
		utils.InternalServerError(w, "Internal Server Error in Unmarshal JSON body in DeleteProduct: ", err)
		return;
	}

	// CheckUserGroup: IMS User and Operations
	if !CheckProductsUserGroup(w, req) {return}

	productSku := deleteProduct.ProductSku

	// Check Product Sku Format
	if !CheckProductSkuFormat(w, productSku) {return}

	

}



