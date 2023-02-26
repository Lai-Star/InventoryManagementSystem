package handlers_products

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	handlers_user_management "github.com/LeonLow97/inventory-management-system-golang-react-postgresql/api/handlers/user-management"
	"github.com/LeonLow97/inventory-management-system-golang-react-postgresql/database"
	"github.com/LeonLow97/inventory-management-system-golang-react-postgresql/utils"
	"github.com/go-chi/chi"
)

func UpdateProduct(w http.ResponseWriter, req *http.Request) {
	// Set Headers
	w.Header().Set("Content-Type", "application/json")
	var updateProduct ProductJson

	// Reading the request body and Unmarshal the body to the ProductJson struct
	bs, _ := io.ReadAll(req.Body)
	if err := json.Unmarshal(bs, &updateProduct); err != nil {
		utils.InternalServerError(w, "Internal Server Error in Unmarshal JSON body in UpdateProduct route", err)
		return
	}

	// CheckUserGroup: IMS User and Operations
	if !handlers_user_management.RetrieveIssuer(w, req) {return}
	if !utils.CheckUserGroup(w, w.Header().Get("username"), "InvenNexus User", "Operations") {return}

	// Get product id from url params
	productIdStr := chi.URLParam(req, "product_id")
	updateProduct.ProductId, _ = strconv.Atoi(productIdStr)
	
	// Trim white spaces in product fields (except sizeName)
	updateProduct = updateProduct.ProductFieldsTrimSpaces()

	var count int
	var err error
	var currentProductSku string

	// Check User Organisation
	username := w.Header().Get("username")
	organisationName, userId, err := database.GetOrganisationNameByUsername(username)
	if err != nil {
		utils.InternalServerError(w, "Internal Server Error in getting company name from database: ", err)
		return
	}
	
	// Check if product id exists for the user or organisation
	if organisationName == "InvenNexus" {
		count, currentProductSku, err = database.GetCountByUserIdAndProductId(userId, updateProduct.ProductId)
	} else {
		count, currentProductSku, err = database.GetCountByOrganisationAndProductId(organisationName, updateProduct.ProductId)
	}

	if err != nil {
		utils.InternalServerError(w, "Internal server error in getting count by product id: ", err)
		return
	}
	if count == 0 {
		utils.ResponseJson(w, http.StatusNotFound, "This product does not exist. Please try again.")
		return
	}

	// Product Form Validation
	if !ProductFormValidation(w, updateProduct, "UPDATE") {return}

	// check if updated product sku is equal to the current product sku in db
	if currentProductSku != updateProduct.ProductSku {
		// Check product sku for duplicates (unless product sku is the same for the current product)
		if organisationName == "InvenNexus" {
			// check the product sku for duplicates based on username
			count, err = database.GetProductSkuCountByUsername(username, updateProduct.ProductSku)
		} else {
			// check the product sku for duplicates based on organisation name
			count, err = database.GetProductSkuCountByOrganisation(organisationName, updateProduct.ProductSku)
		}
		if err != nil {
			utils.InternalServerError(w, "Internal server error in getting product sku count: ", err)
			return
		}
		if count >= 1 {
			utils.ResponseJson(w, http.StatusBadRequest, "Product Sku already exists. Please try again.")
			return
		}
	}

	// Check that the size name is exists/valid
	
	

	utils.ResponseJson(w, http.StatusOK, "Successfully updated the product!")
}