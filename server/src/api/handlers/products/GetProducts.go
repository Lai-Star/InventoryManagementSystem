package products

import (
	"database/sql"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/LeonLow97/inventory-management-system-golang-react-postgresql/api/handlers"
	handlers_user_management "github.com/LeonLow97/inventory-management-system-golang-react-postgresql/api/handlers/user-management"
	"github.com/LeonLow97/inventory-management-system-golang-react-postgresql/database"
	"github.com/LeonLow97/inventory-management-system-golang-react-postgresql/utils"
	"github.com/jackc/pgx/v4"
)

func GetProducts(w http.ResponseWriter, req *http.Request) {
	// CheckUserGroup: IMS User and Operations
	if !handlers_user_management.RetrieveIssuer(w, req) {
		return
	}
	if !utils.CheckUserGroup(w, w.Header().Get("username"), "InvenNexus User", "Operations") {
		return
	}

	// Retrieve products from SQL database
	var data []handlers.Product
	productSizes := make(map[string][]handlers.ProductSize)

	// To handle nullable columns in a database table
	var productName, productDescription, productSku, productColour, productCategory, productBrand sql.NullString
	var productCost sql.NullFloat64
	var sizeName sql.NullString
	var productId, userId, sizeQuantity sql.NullInt32

	// Check User Organisation
	username := w.Header().Get("username")
	organisationName, userID, err := database.GetOrganisationNameAndUserIdByUsername(username)
	if err != nil {
		utils.ResponseJson(w, http.StatusInternalServerError, "Internal Server Error")
		log.Println("Internal server error in getting company name from database:", err)
		return
	}
	userId.Int32 = int32(userID)

	var rows pgx.Rows
	if organisationName == "InvenNexus" {
		rows, err = database.GetProductsByUsername(userID)
	} else {
		rows, err = database.GetProductsByOrganisation(organisationName)
	}
	if err != nil {
		utils.ResponseJson(w, http.StatusInternalServerError, "Internal Server Error")
		log.Println("Internal Server Error in GetProducts:", err)
		return
	}

	defer rows.Close()

	// To keep track of products that have already been added to the data slice
	addedProducts := make(map[string]bool)

	for rows.Next() {
		err = rows.Scan(&productId, &productName, &productDescription, &productSku, &productColour, &productCategory, &productBrand, &productCost, &sizeName, &sizeQuantity)
		if err != nil {
			utils.ResponseJson(w, http.StatusInternalServerError, "Internal Server Error")
			log.Println("Internal Server Error in Scanning GetProducts:", err)
			return
		}

		if !addedProducts[productSku.String] {
			product := handlers.Product{
				ProductId:          int(productId.Int32),
				ProductName:        productName.String,
				ProductDescription: productDescription.String,
				ProductSku:         productSku.String,
				ProductColour:      productColour.String,
				ProductCategory:    productCategory.String,
				ProductBrand:       productBrand.String,
				ProductCost:        float32(productCost.Float64),
			}

			// only append if the product sku does not exist in the slice.
			data = append(data, product)
			addedProducts[productSku.String] = true
		}

		productSize := handlers.ProductSize{
			SizeName:     sizeName.String,
			SizeQuantity: int(sizeQuantity.Int32),
		}
		productSizes[productSku.String] = append(productSizes[productSku.String], productSize)
	}

	for i, product := range data {
		data[i].Sizes = productSizes[product.ProductSku]
	}

	jsonStatus := struct {
		Code     int                `json:"code"`
		Response []handlers.Product `json:"response"`
	}{
		Response: data,
		Code:     http.StatusOK,
	}

	bs, err := json.Marshal(jsonStatus)
	if err != nil {
		utils.ResponseJson(w, http.StatusInternalServerError, "Internal Server Error")
		log.Println("Internal Server Error in Marshal JSON body in GetProducts:", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	io.WriteString(w, string(bs))
}
