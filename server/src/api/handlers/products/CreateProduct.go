package handlers_products

import (
	"encoding/json"
	"fmt"
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

	var isValidFormValidation bool
	// Check if user provided a size
	if len(createProduct.Sizes) != 0 {
		// Trim White Spaces in each size, make all size name uppercase and check format provided.
		isValidFormValidation, createProduct.Sizes = SizesFormValidation(w, createProduct.Sizes)
		if !isValidFormValidation {return}
	}

	// Insert new product into PostgreSQL database
	err, productId := database.InsertNewProduct(createProduct.ProductName, createProduct.ProductDescription, createProduct.ProductSku, createProduct.ProductColour, createProduct.ProductCategory, createProduct.ProductBrand, createProduct.ProductCost)
	if err != nil {
		utils.InternalServerError(w, "Internal Server Error in InsertNewProduct: ", err)
		return
	}

	// Insert new size into PostgreSQL database

	// Insert new product_size into PostgreSQL database
	

	fmt.Println("Product Id: ", productId)

	utils.ResponseJson(w, http.StatusOK, "Successfully created a new product!")
}

