package products

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
	if !handlers_user_management.RetrieveIssuer(w, req) {
		return
	}
	if !utils.CheckUserGroup(w, w.Header().Get("username"), "InvenNexus User", "Operations") {
		return
	}

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
	organisationName, userId, err := database.GetOrganisationNameAndUserIdByUsername(username)
	if err != nil {
		utils.InternalServerError(w, "Internal Server Error in getting company name from database: ", err)
		return
	}

	// Check if product id exists for the user or organisation
	if organisationName == "InvenNexus" {
		count, currentProductSku, _ = database.GetCountProductSkuByUserIdAndProductId(userId, updateProduct.ProductId)
	} else {
		count, currentProductSku, _ = database.GetCountProductSkuByOrganisationAndProductId(organisationName, updateProduct.ProductId)
	}

	// if err != nil {
	// 	utils.InternalServerError(w, "Internal server error in getting count by product id: ", err)
	// 	return
	// }
	if count == 0 {
		utils.ResponseJson(w, http.StatusNotFound, "This product does not exist. Please try again.")
		return
	}

	// Product Form Validation
	if !ProductFormValidation(w, updateProduct, "UPDATE") {
		return
	}

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

	// Check if brand, category or colour exists
	var brandCount, categoryCount, colourCount int
	var errBrand, errCategory, errColour error
	if organisationName == "InvenNexus" {
		brandCount, errBrand = database.GetBrandNameCountByUsername(userId, updateProduct.ProductBrand)
		colourCount, errCategory = database.GetColourNameCountByUsername(userId, updateProduct.ProductColour)
		categoryCount, errColour = database.GetCategoryNameCountByUsername(userId, updateProduct.ProductCategory)
	} else {
		brandCount, errBrand = database.GetBrandNameCountByOrganisation(organisationName, updateProduct.ProductBrand)
		colourCount, errCategory = database.GetColourNameCountByOrganisation(organisationName, updateProduct.ProductColour)
		categoryCount, errColour = database.GetCategoryNameCountByOrganisation(organisationName, updateProduct.ProductCategory)
	}

	if errBrand != nil {
		utils.InternalServerError(w, "Internal server error in getting brand name by count: ", err)
		return
	}
	if errCategory != nil {
		utils.InternalServerError(w, "Internal server error in getting category name by count: ", err)
		return
	}
	if errColour != nil {
		utils.InternalServerError(w, "Internal server error in getting colour name by count: ", err)
		return
	}

	if brandCount == 0 {
		utils.ResponseJson(w, http.StatusNotFound, "Brand name does not exist. Please try again.")
		return
	}
	if colourCount == 0 {
		utils.ResponseJson(w, http.StatusNotFound, "Colour does not exist. Please try again.")
		return
	}
	if categoryCount == 0 {
		utils.ResponseJson(w, http.StatusNotFound, "Category does not exist. Please try again.")
		return
	}

	// Check that the size name is exists/valid for the product
	if len(updateProduct.Sizes) >= 1 {
		for _, size := range updateProduct.Sizes {
			sizeName := size.SizeName

			var count int
			var err error

			if organisationName == "InvenNexus" {
				count, err = database.GetSizeNameCountByUsername(userId, sizeName)
			} else {
				count, err = database.GetSizeNameCountByOrganisation(organisationName, sizeName)
			}
			if err != nil {
				utils.InternalServerError(w, "Internal server error in getting size count: ", err)
				return
			}
			if count == 0 {
				utils.ResponseJson(w, http.StatusNotFound, sizeName+" size name does not exist. Please create this size or try again.")
				return
			}

			if organisationName == "InvenNexus" {
				count, err = database.GetSizeNameCountByUserIdAndProductId(updateProduct.ProductId, userId, sizeName)
			} else {
				count, err = database.GetSizeNameCountByOrganisationIdAndProductId(updateProduct.ProductId, organisationName, sizeName)
			}
			if err != nil {
				utils.InternalServerError(w, "Internal server error in getting size count by product id: ", err)
				return
			}
			if count == 0 {
				utils.ResponseJson(w, http.StatusNotFound, sizeName+" size name does not exist for this product. Please try again.")
				return
			}
		}
	}

	// Update products in the products table by product id
	err = database.UpdateProductsByProductID(updateProduct.ProductName, updateProduct.ProductDescription, updateProduct.ProductSku, updateProduct.ProductCost, updateProduct.ProductId)
	if err != nil {
		utils.InternalServerError(w, "Internal server error in updating products table: ", err)
		return
	}

	// Update product_user_mapping or product_organisation_mapping table
	if organisationName == "InvenNexus" {
		// regular user update
		err = database.UpdateProductUserMapping(userId, updateProduct.ProductId, updateProduct.ProductColour, updateProduct.ProductCategory, updateProduct.ProductBrand)
	} else {
		// organisation update
		err = database.UpdateProductOrganisationMapping(updateProduct.ProductId, organisationName, updateProduct.ProductColour, updateProduct.ProductCategory, updateProduct.ProductBrand)
	}

	if err != nil {
		utils.InternalServerError(w, "Internal server error in updating product user/organisation mapping table: ", err)
		return
	}

	// Update user_product_sizes_mapping or organisation_product_sizes_mapping
	for _, size := range updateProduct.Sizes {
		sizeName := size.SizeName
		sizeQuantity := size.SizeQuantity

		var updateFunc func(int, int, string) error
		if organisationName == "InvenNexus" {
			updateFunc = database.UpdateUserProductSizesMapping
		} else {
			updateFunc = database.UpdateOrganisationProductSizesMapping
		}
		err = updateFunc(sizeQuantity, updateProduct.ProductId, sizeName)
		if err != nil {
			utils.InternalServerError(w, "Internal server error in inserting new size: ", err)
			return
		}
	}

	utils.ResponseJson(w, http.StatusOK, "Successfully updated the product!")
}
