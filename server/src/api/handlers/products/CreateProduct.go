package handlers_products

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/LeonLow97/inventory-management-system-golang-react-postgresql/database"
	"github.com/LeonLow97/inventory-management-system-golang-react-postgresql/utils"
)

func CreateProduct(w http.ResponseWriter, req *http.Request) {
	// Set Headers
	w.Header().Set("Content-Type", "application/json")
	var createProduct ProductJson

	// Reading the request body and Unmarshal the body to the ProductJson struct
	bs, _ := io.ReadAll(req.Body)
	if err := json.Unmarshal(bs, &createProduct); err != nil {
		utils.InternalServerError(w, "Internal Server Error in UnMarshal JSON body in CreateProduct route: ", err)
		return;
	}

	// CheckUserGroup: IMS User and Operations
	if !CheckProductsUserGroup(w, req) {return}

	// Trim White Spaces in product fields
	createProduct = createProduct.ProductFieldsTrimSpaces()

	// Product Form Validation
	if !ProductFormValidation(w, createProduct) {return}

	// Check Product Sku to see if it exists in database (cannot have duplicates)
	isExistProductSku := database.ProductSkuExists(createProduct.ProductSku)
	if isExistProductSku {
		utils.ResponseJson(w, http.StatusBadRequest, "Product Sku already exists. Please try again.")
		return
	}

	// Check if user provided a size
	


	// Insert new product into PostgreSQL database
	err := database.InsertNewProduct(createProduct.ProductName, createProduct.ProductDescription, createProduct.ProductSku, createProduct.ProductColour, createProduct.ProductCategory, createProduct.ProductBrand, createProduct.ProductCost)
	if err != nil {
		utils.InternalServerError(w, "Internal Server Error in InsertNewProduct: ", err)
		return
	}

	utils.ResponseJson(w, http.StatusOK, "Successfully created a new product!")
}

