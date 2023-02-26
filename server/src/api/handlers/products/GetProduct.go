package handlers_products

import (
	"database/sql"
	"encoding/json"
	"io"
	"net/http"

	"github.com/LeonLow97/inventory-management-system-golang-react-postgresql/api/handlers"
	handlers_user_management "github.com/LeonLow97/inventory-management-system-golang-react-postgresql/api/handlers/user-management"
	"github.com/LeonLow97/inventory-management-system-golang-react-postgresql/database"
	"github.com/LeonLow97/inventory-management-system-golang-react-postgresql/utils"
)

func GetProducts(w http.ResponseWriter, req *http.Request) {

	// CheckUserGroup: IMS User and Operations
	if !handlers_user_management.RetrieveIssuer(w, req) {return}
	if !utils.CheckUserGroup(w, w.Header().Get("username"), "InvenNexus User", "Operations") {return}

	// Retrieve products from SQL database
	var data []handlers.Product
	productSizes := make(map[string][]handlers.ProductSize)

	// To handle nullable columns in a database table
	var productName, productDescription, productSku, productColour, productCategory, productBrand sql.NullString
	var productCost sql.NullFloat64
	var sizeName sql.NullString
	var sizeQuantity sql.NullInt32

	// Check User Organisation
	username := w.Header().Get("username")
	organisationName, _, err := database.GetOrganisationNameByUsername(username)
	if err != nil {
		utils.InternalServerError(w, "Internal server error in getting company name from database: ", err)
		return
	}

	var rows *sql.Rows
	if organisationName == "InvenNexus" {
		rows, err = database.GetProducts("user_sizes")
	} else {
		rows, err = database.GetProducts("organisation_sizes")
	}
	if err != nil {
		utils.InternalServerError(w, "Internal Server Error in GetProducts: ", err)
		return
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&productName, &productDescription, &productSku, &productColour, &productCategory, &productBrand, &productCost, &sizeName, &sizeQuantity)
		if err != nil {
			utils.InternalServerError(w, "Internal Server Error in Scanning GetProducts: ", err)
			return
		}

		product := handlers.Product{
			ProductName: productName.String,
			ProductDescription: productDescription.String,
			ProductSku: productSku.String,
			ProductColour: productColour.String,
			ProductCategory: productCategory.String,
			ProductBrand: productBrand.String,
			ProductCost: float32(productCost.Float64),
		}

		data = append(data, product)

		productSize := handlers.ProductSize {
			SizeName: sizeName.String,
			SizeQuantity: int(sizeQuantity.Int32),
		}
		productSizes[productSku.String] = append(productSizes[productSku.String], productSize)
	}

	for i, product := range data {
		data[i].Sizes = productSizes[product.ProductSku]
	}

	jsonStatus := struct {
		Code int `json:"code"`
		Response []handlers.Product `json:"response"`
	}{
		Response: data,
		Code: http.StatusOK,
	}

	bs, err := json.Marshal(jsonStatus);
	if err != nil {
		utils.InternalServerError(w, "Internal Server Error in Marshal JSON body in GetProducts: ", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	io.WriteString(w, string(bs));
}


