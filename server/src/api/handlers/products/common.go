package handlers_products

import (
	"net/http"
	"reflect"
	"regexp"
	"strconv"
	"strings"

	handlers_user "github.com/LeonLow97/inventory-management-system-golang-react-postgresql/api/handlers/user"
	"github.com/LeonLow97/inventory-management-system-golang-react-postgresql/utils"
)

type ProductJson struct {
	ProductId int
	ProductName        string  `json:"product_name"`
	ProductDescription string  `json:"product_description"`
	ProductSku         string  `json:"product_sku"`
	ProductColour      string  `json:"product_colour"`
	ProductCategory    string  `json:"product_category"`
	ProductBrand       string  `json:"product_brand"`
	ProductCost        float32 `json:"product_cost"`
	Sizes              []Size  `json:"sizes"`
}

type Size struct {
	SizeName     string `json:"size_name"`
	SizeQuantity int    `json:"size_quantity"`
}

type DeleteProductJson struct {
	ProductId int
}

func CheckProductsUserGroup(w http.ResponseWriter, req *http.Request) bool {
	// CheckUserGroup: IMS User and Operations
	if !handlers_user.RetrieveIssuer(w, req) {return false}
	checkUserGroupIMSUser := utils.CheckUserGroup(w.Header().Get("username"), "IMS User")
	checkUserGroupOperations := utils.CheckUserGroup(w.Header().Get("username"), "Operations")
	if !checkUserGroupIMSUser || !checkUserGroupOperations {
		utils.ResponseJson(w, http.StatusForbidden, "Access Denied: You do not have permission to access this resource.")
		return false
	}
	return true
}

func ProductFormValidation(w http.ResponseWriter, product ProductJson) bool {

	// Check Product Name (Length 0 - 255)
	isValidProductName := utils.CheckLengthRange(product.ProductName, 0, 255)
	if !isValidProductName {
		utils.ResponseJson(w, http.StatusBadRequest, "Product Name must have a length of 0 - 255 characters.")
		return false
	}

	// Check Product Sku (Length 0 - 50, must be unique)
	isValidProductSku := utils.CheckLengthRange(product.ProductSku, 0, 50)
	if !isValidProductSku {
		utils.ResponseJson(w, http.StatusBadRequest, "Product Sku must have a length of 0 - 50 characters.")
		return false
	}

	// Check Product Colour (Length 0 - 7)
	isValidProductColour := utils.CheckLengthRange(product.ProductColour, 0, 7)
	if !isValidProductColour {
		utils.ResponseJson(w, http.StatusBadRequest, "Product Colour must have a length of 0 - 7 characters.")
		return false
	}

	// Check Product Category (Length 0 - 20)
	isValidProductCategory := utils.CheckLengthRange(product.ProductCategory, 0, 20)
	if !isValidProductCategory {
		utils.ResponseJson(w, http.StatusBadRequest, "Product Category must have a length of 0 - 20 characters.")
		return false
	}
	
	// Check Product Brand (Length 0 - 50)
	isValidProductBrand := utils.CheckLengthRange(product.ProductBrand, 0, 50)
	if !isValidProductBrand {
		utils.ResponseJson(w, http.StatusBadRequest, "Product Brand must have a length of 0 - 50 characters.")
		return false
	}

	// Check Product Cost (Decimal(10, 2))
	productCost := strconv.FormatFloat(float64(product.ProductCost), 'f', -1, 32)
	isValidProductCost, _ := regexp.MatchString(`^[0-9]{1,10}(\.[0-9]{1,2})?$`, productCost)
	if !isValidProductCost {
		utils.ResponseJson(w, http.StatusBadRequest, "Product Cost must have a maximum of 8 digits and 2 decimal places.")
		return false
	}
	
	return true
}

func (product ProductJson) ProductFieldsTrimSpaces() ProductJson {

	product.ProductName = strings.TrimSpace(product.ProductName)
	product.ProductSku = strings.TrimSpace(product.ProductSku)
	product.ProductColour = strings.TrimSpace(product.ProductColour)
	product.ProductCategory = strings.TrimSpace(product.ProductCategory)
	product.ProductBrand = strings.TrimSpace(product.ProductCategory)

	return product

}

func SizesFormValidation(w http.ResponseWriter, sizes []Size) (bool, []Size) {

	for _, size := range sizes {
		// Check length of SizeName (length of 0 - 5)
		isValidSizeName := utils.CheckLengthRange(size.SizeName, 0, 5)
		if !isValidSizeName {
			utils.ResponseJson(w, http.StatusBadRequest, "Size name must have a length of 0 - 5 characters. Please try again!")
			return false, sizes
		}

		// Trim white space for each size
		size.SizeName = strings.TrimSpace(size.SizeName)

		// SizeName must be in uppercase
		size.SizeName = strings.ToUpper(size.SizeName)

		// Check that this size name is valid (XXS, XS, S, M, L, XL, XXL)
		isValidSize := IsAllowedProductSize(size.SizeName)
		if !isValidSize {
			utils.ResponseJson(w, http.StatusBadRequest, size.SizeName + " is not a valid size name. Please try again.")
			return false, sizes
		}

		// Check that size quantity is of type int
		// reflect: inspect and manipulate values of different types dynamically at runtime
		// reflect.TypeOf gets the value stored in SizeQuantity and .Kind() checks the data type of the value
		if reflect.TypeOf(size.SizeQuantity).Kind() != reflect.Int {
			utils.ResponseJson(w, http.StatusBadRequest, strconv.Itoa(size.SizeQuantity) + " is not in the correct Integer format. Please try again!")
			return false, sizes
		}
	}

	return true, sizes

}


func IsAllowedProductSize(size string) bool {
	allowedSizes := []string {"XXS", "XS", "S", "M", "L", "XL", "XXL"}
	for _, allowedSize := range allowedSizes {
		if size == allowedSize {
			return true
		}
	}
	return false
}




