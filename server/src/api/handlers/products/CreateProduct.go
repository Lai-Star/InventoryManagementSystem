package handlers_products

import (
	"encoding/json"
	"io"
	"net/http"

	handlers_user_management "github.com/LeonLow97/inventory-management-system-golang-react-postgresql/api/handlers/user-management"
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
	if !ProductFormValidation(w, newProduct) {return}

	// // Check User Organisation
	// username := w.Header().Get("username")
	// userOrganisation, err := database.GetOrganisationNameByUsername(username)
	// if err != nil {
	// 	utils.InternalServerError(w, "Internal Server Error in getting company name from database: ", err)
	// 	return
	// }

	// // Check Product Sku to see if it exists in database (cannot have duplicates within the same organisation)
	// // if user belongs to "IMSer" means it is a regular user, check for existing product sku based on username
	// if userOrganisation == "IMSer" {
	// 	// isExistProductSku := database.ProductSkuExists(newProduct.ProductSku)
	// 	// if isExistProductSku {
	// 	// 	utils.ResponseJson(w, http.StatusBadRequest, "Product Sku already exists. Please try again.")
	// 	// 	return
	// 	// }
	// } else {
	// 	// user does not belong to "IMSer", user belongs to specific organisation
	// }


	// // Insert new product into products table
	// productId, err := database.InsertNewProduct(createProduct.ProductName, createProduct.ProductDescription, createProduct.ProductSku, createProduct.ProductColour, createProduct.ProductCategory, createProduct.ProductBrand, createProduct.ProductCost)
	// if err != nil {
	// 	utils.InternalServerError(w, "Internal Server Error in InsertNewProduct: ", err)
	// 	return
	// }

	// var isValidFormValidation bool
	// // Check if user provided a size
	// if len(createProduct.Sizes) > 0 {
	// 	// Trim White Spaces in each size, make all size name uppercase and check format provided.
	// 	isValidFormValidation, createProduct.Sizes = ValidateAndInsertSize(w, createProduct.Sizes, productId)
	// 	if !isValidFormValidation {return}
	// }

	utils.ResponseJson(w, http.StatusOK, "Successfully created a new product!")
}

