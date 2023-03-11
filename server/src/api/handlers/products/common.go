package handlers_products

import (
	"net/http"
	"reflect"
	"regexp"
	"strconv"
	"strings"

	"github.com/LeonLow97/inventory-management-system-golang-react-postgresql/utils"
)

type ProductJson struct {
	ProductId int `json:"product_id"`
	ProductName        string  `json:"product_name"`
	ProductDescription string  `json:"product_description"`
	ProductSku         string  `json:"product_sku"`
	ProductBrand       string  `json:"product_brand"`
	ProductColour      string  `json:"product_colour"`
	ProductCategory    string  `json:"product_category"`
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

func ProductNameFormValidation(w http.ResponseWriter, productName, action string) bool {
	if action == "CREATE" {
		// Check if product name is empty
		if utils.IsBlankField(productName) {
			utils.ResponseJson(w, http.StatusBadRequest, "Product Name cannot be empty. Please try again.")
			return false
		}
	}

	if action == "CREATE" || (action == "UPDATE" && len(productName) > 0) {
		// Check Product Name (Length 1 - 255)
		if !utils.CheckLengthRange(productName, 1, 255) {
			utils.ResponseJson(w, http.StatusBadRequest, "Product Name must have a length of 1 - 255 characters.")
			return false
		}
	}

	return true
}

func ProductSkuFormValidation(w http.ResponseWriter, productSku, action string) bool {
	if action == "CREATE" {
		// Check if product sku is empty
		if utils.IsBlankField(productSku) {
			utils.ResponseJson(w, http.StatusBadRequest, "Product Sku cannot be empty. Please try again.")
			return false
		}
	}

	if action == "CREATE" || (action == "UPDATE" && len(productSku) > 0) {
		// Check Product Sku (Length 1 - 50, must be unique)
		if !utils.CheckLengthRange(productSku, 1, 50) {
			utils.ResponseJson(w, http.StatusBadRequest, "Product Sku must have a length of 1 - 50 characters.")
			return false
		}
	}

	return true
}

func ProductBrandFormValidation(w http.ResponseWriter, productBrand, action string) bool {
	if action == "CREATE" {
		// Check if product brand is empty
		if utils.IsBlankField(productBrand) {
			utils.ResponseJson(w, http.StatusBadRequest, "Product Brand cannot be empty. Please try again.")
			return false
		}
	}

	if action == "CREATE" || (action == "UPDATE" && len(productBrand) > 0) {
		// Check Product Brand (Length 1 - 50)
		if !utils.CheckLengthRange(productBrand, 1, 50) {
			utils.ResponseJson(w, http.StatusBadRequest, "Product Brand must have a length of 1 - 50 characters.")
			return false
		}
	}

	return true
}

func ProductColourFormValidation(w http.ResponseWriter, productColour, action string) bool {
	if action == "CREATE" {
		// Check if product colour is empty
		if utils.IsBlankField(productColour) {
			utils.ResponseJson(w, http.StatusBadRequest, "Product Colour cannot be empty. Please try again.")
			return false
		}
	}

	if action == "CREATE" || (action == "UPDATE" && len(productColour) > 0) {
		// Check Product Colour (Length 1 - 7)
		if !utils.CheckLengthRange(productColour, 1, 7) {
			utils.ResponseJson(w, http.StatusBadRequest, "Product Colour must have a length of 1 - 7 characters.")
			return false
		}
	}

	return true
}

func ProductCategoryFormValidation(w http.ResponseWriter, productCategory, action string) bool {
	if action == "CREATE" {
		// Check if product category is empty
		if utils.IsBlankField(productCategory) {
			utils.ResponseJson(w, http.StatusBadRequest, "Product Category cannot be empty. Please try again.")
			return false
		}
	}

	if action == "CREATE" || (action == "UPDATE" && len(productCategory) > 0) {
		// Check Product Category (Length 1 - 20)
		if !utils.CheckLengthRange(productCategory, 1, 20) {
			utils.ResponseJson(w, http.StatusBadRequest, "Product Category must have a length of 1 - 20 characters.")
			return false
		}
	}

	return true
}

func ProductCostFormValidation(w http.ResponseWriter, productCost float32) bool {
	// Check if product cost is 0
	if productCost == 0 {
		utils.ResponseJson(w, http.StatusBadRequest, "Product Cost cannot be 0. Please try again.")
		return false
	}

	// Check Product Cost (Decimal(10, 2))
	convertedProductCost := strconv.FormatFloat(float64(productCost), 'f', -1, 32)
	isValidProductCost, _ := regexp.MatchString(`^[0-9]{1,10}(\.[0-9]{1,2})?$`, convertedProductCost)
	if !isValidProductCost {
		utils.ResponseJson(w, http.StatusBadRequest, "Product Cost must have a maximum of 8 digits and 2 decimal places.")
		return false
	}
	return true
}

func SizeNameQuantityFormValidation(w http.ResponseWriter, sizes []Size) bool {
	for idx, size := range(sizes) {
		// trim white spaces for size name
		sizes[idx].SizeName = strings.TrimSpace(size.SizeName)

		// Check if size name is empty
		if utils.IsBlankField(size.SizeName) {
			utils.ResponseJson(w, http.StatusBadRequest, "Size Name cannot be empty. Please try again.")
			return false
		}

		// Check if size name is within 1 and 5 characters
		if !utils.CheckLengthRange(size.SizeName, 1, 5) {
			utils.ResponseJson(w, http.StatusBadRequest, "Size Name must be within 1 - 5 characters. Please try again.")
			return false
		}

		// Check if the type of size quantity is an integer
		if kind := reflect.TypeOf(size.SizeQuantity).Kind(); kind != reflect.Int &&
		kind != reflect.Int8 && kind != reflect.Int16 &&
		kind != reflect.Int32 && kind != reflect.Int64 {
			utils.ResponseJson(w, http.StatusBadRequest, "Size Quantity must be an integer value. Please try again.")
			return false
		}
		
		// Check if size quantity is negative
		if size.SizeQuantity < 0 {
			utils.ResponseJson(w, http.StatusBadRequest, "Size Quantity cannot be negative. Please try again.")
			return false
		}

		// Check if size name is valid
		
	}
	return true
}

func SizeNameFormValidation(w http.ResponseWriter, sizeName string) bool {
	// Check if size name is empty
	if utils.IsBlankField(sizeName) {
		utils.ResponseJson(w, http.StatusBadRequest, "Size Name cannot be empty. Please try again.")
		return false
	}

	// Check if size name is within 1 and 5 characters
	if !utils.CheckLengthRange(sizeName, 1, 5) {
		utils.ResponseJson(w, http.StatusBadRequest, "Size Name must be within 1 - 5 characters. Please try again.")
		return false
	}
	return true
}

func ProductFormValidation(w http.ResponseWriter, product ProductJson, action string) bool {

	if !ProductNameFormValidation(w, product.ProductName, action) {return false}
	if !ProductSkuFormValidation(w, product.ProductSku, action) {return false}
	if !ProductBrandFormValidation(w, product.ProductBrand, action) {return false}
	if !ProductColourFormValidation(w, product.ProductColour, action) {return false}
	if !ProductCategoryFormValidation(w, product.ProductCategory, action) {return false}
	if !ProductCostFormValidation(w, product.ProductCost) {return false}

	if len(product.Sizes) >= 1 {
		if !SizeNameQuantityFormValidation(w, product.Sizes) {return false}
	}
	
	return true
}

func (product ProductJson) ProductFieldsTrimSpaces() ProductJson {

	product.ProductName = strings.TrimSpace(product.ProductName)
	product.ProductSku = strings.TrimSpace(product.ProductSku)
	product.ProductColour = strings.TrimSpace(product.ProductColour)
	product.ProductCategory = strings.TrimSpace(product.ProductCategory)
	product.ProductBrand = strings.TrimSpace(product.ProductBrand)

	return product

}

func ValidateAndInsertSize(w http.ResponseWriter, sizes []Size, productId int32) (bool, []Size) {

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

		// Insert Size to sizes table
		// sizeId, err := database.InsertNewSize(size.SizeName, size.SizeQuantity)
		// if err != nil {
		// 	utils.InternalServerError(w, "Error in inserting new size to sizes table: ", err)
		// 	return false, sizes
		// }

		// // Insert product_id and size_id to product_sizes table
		// err = database.InsertNewProductSizes(productId, sizeId)
		// if err != nil {
		// 	utils.InternalServerError(w, "Error in inserting new productsizes into product_sizes table: ", err)
		// 	return false, sizes
		// }
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




