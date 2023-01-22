package handlers_products

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/LeonLow97/inventory-management-system-golang-react-postgresql/database"
	"github.com/LeonLow97/inventory-management-system-golang-react-postgresql/utils"
)

func UpdateProduct(w http.ResponseWriter, req *http.Request) {
	// Set Headers
	w.Header().Set("Content-Type", "application/json")
	var updateProduct ProductJson

	// Reading the request body and Unmarshal the body to the ProductJson struct
	bs, _ := io.ReadAll(req.Body)
	if err := json.Unmarshal(bs, &updateProduct); err != nil {
		utils.InternalServerError(w, "Internal Server Error in UnMarshal JSON body in CreateProduct route: ", err)
		return;
	}

	// Check user group for products
	if !CheckProductsUserGroup(w, req) {return}

	// Product Form Validation
	if !ProductsFormValdiation(w, updateProduct) {return}

	// Get productid from url params
	productId := req.URL.Query().Get("product_id")
	x, _ := strconv.Atoi(productId)

	// Update products table
	err := database.UpdateProduct(x, updateProduct.ProductName, updateProduct.ProductDescription, updateProduct.ProductSku, updateProduct.ProductColour, updateProduct.ProductCategory, updateProduct.ProductBrand, updateProduct.ProductCost)
	if err != nil {
		utils.InternalServerError(w, "Internal Server Error in UpdateProduct: ", err)
		return
	}

	utils.ResponseJson(w, http.StatusOK, "Successfully updated product!")
	
}