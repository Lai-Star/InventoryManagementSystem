package handlers_products

import (
	"encoding/json"
	"io"
	"net/http"

	handlers_user_management "github.com/LeonLow97/inventory-management-system-golang-react-postgresql/api/handlers/user-management"
	"github.com/LeonLow97/inventory-management-system-golang-react-postgresql/database"
	"github.com/LeonLow97/inventory-management-system-golang-react-postgresql/utils"
)

func CreateProduct(w http.ResponseWriter, req *http.Request) {
	// Set Headers
	w.Header().Set("Content-Type", "application/json")
	var newProduct ProductJson

	// Reading the request body and Unmarshal the body to the ProductJson struct
	bs, _ := io.ReadAll(req.Body)
	if err := json.Unmarshal(bs, &newProduct); err != nil {
		utils.InternalServerError(w, "Internal Server Error in UnMarshal JSON body in CreateProduct route: ", err)
		return;
	}
	
	// CheckUserGroup: IMS User and Operations
	if !handlers_user_management.RetrieveIssuer(w, req) {return}
	if !utils.CheckUserGroup(w, w.Header().Get("username"), "InvenNexus User", "Operations") {return}
	
	// Trim White Spaces in product fields
	newProduct = newProduct.ProductFieldsTrimSpaces()
	
	// Product Form Validation
	if !ProductFormValidation(w, newProduct, "CREATE") {return}

	// Check User Organisation
	username := w.Header().Get("username")
	organisationName, userId, err := database.GetOrganisationNameByUsername(username)
	if err != nil {
		utils.InternalServerError(w, "Internal Server Error in getting company name from database: ", err)
		return
	}

	// Check Product Sku to see if it exists in database (cannot have duplicates within the same organisation)
	var count int
	if organisationName == "InvenNexus" {
		// check the product sku for duplicates based on username
		count, err = database.GetProductSkuCountByUsername(username, newProduct.ProductSku)
	} else {
		// check the product sku for duplicates based on organisation name
		count, err = database.GetProductSkuCountByOrganisation(organisationName, newProduct.ProductSku)
	}

	if err != nil {
		utils.InternalServerError(w, "Internal server error in getting product sku count: ", err)
		return
	}
	if count >= 1 {
		utils.ResponseJson(w, http.StatusBadRequest, "Product Sku already exists. Please try again.")
		return
	}

	// Check if brand, category or colour exists
	var brandCount, categoryCount, colourCount int
	var errBrand, errCategory, errColour error
	if organisationName == "InvenNexus" {
		brandCount, errBrand = database.GetBrandNameCountByUsername(userId, newProduct.ProductBrand)
		colourCount, errCategory = database.GetColourNameCountByUsername(userId, newProduct.ProductColour)
		categoryCount, errColour = database.GetCategoryNameCountByUsername(userId, newProduct.ProductCategory)
	} else {
		brandCount, errBrand = database.GetBrandNameCountByOrganisation(organisationName, newProduct.ProductBrand)
		colourCount, errCategory = database.GetColourNameCountByOrganisation(organisationName, newProduct.ProductColour)
		categoryCount, errColour = database.GetCategoryNameCountByOrganisation(organisationName, newProduct.ProductCategory)
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

	// Check if size name exists
	if len(newProduct.Sizes) >= 1 {
		for _, size := range(newProduct.Sizes) {
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
				utils.ResponseJson(w, http.StatusNotFound, "The size name does not exist. Please try again.")
				return
			}
		}
	}

	// Insert product details to `products` table
	productId, err := database.InsertNewProduct(newProduct.ProductName, newProduct.ProductDescription, newProduct.ProductSku, newProduct.ProductCost)
	if err != nil {
		utils.InternalServerError(w, "Internal server error in inserting new product into products table: ", err)
		return
	}

	if len(newProduct.Sizes) >= 1 {
		for _, size := range(newProduct.Sizes) {
			sizeName := size.SizeName
			sizeQuantity := size.SizeQuantity
			
			var insertFunc func(string, int, int) error
			if organisationName == "InvenNexus" {
				insertFunc = database.InsertIntoUserProductSizesMapping
			} else {
				insertFunc = database.InsertIntoOrganisationProductSizesMapping
			}
			err = insertFunc(sizeName, sizeQuantity, productId)
			if err != nil {
				utils.InternalServerError(w, "Internal server error in inserting new size: ", err)
				return
			}
		}
	}

	// Insert to product_user_mapping or product_organisation_mapping table
	if organisationName == "InvenNexus" {
		err = database.InsertIntoProductUserMapping(productId, userId, newProduct.ProductColour, newProduct.ProductCategory, newProduct.ProductBrand)
		if err != nil {
			utils.InternalServerError(w, "Internal server error in inserting into product_user_mapping table: ", err)
			return
		}
	} else {
		err = database.InsertIntoProductOrganisationMapping(productId, organisationName, newProduct.ProductColour, newProduct.ProductCategory, newProduct.ProductBrand)
		if err != nil {
			utils.InternalServerError(w, "Internal server error in inserting into product_organisation_mapping table: ", err)
			return
		}
	}

	utils.ResponseJson(w, http.StatusOK, "Successfully created a new product!")
}

