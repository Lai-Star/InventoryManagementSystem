package handlers_products

import (
	"net/http"
	"strconv"

	handlers_user_management "github.com/LeonLow97/inventory-management-system-golang-react-postgresql/api/handlers/user-management"
	"github.com/LeonLow97/inventory-management-system-golang-react-postgresql/database"
	"github.com/LeonLow97/inventory-management-system-golang-react-postgresql/utils"
	"github.com/go-chi/chi"
)

func DeleteProduct(w http.ResponseWriter, req *http.Request) {
	
	// Set Headers
	w.Header().Set("Content-Type", "application/json")
	var deleteProduct DeleteProductJson

	// CheckUserGroup: IMS User and Operations
	if !handlers_user_management.RetrieveIssuer(w, req) {return}
	if !utils.CheckUserGroup(w, w.Header().Get("username"), "InvenNexus User", "Operations") {return}

	// Get productid from url params
	productIdStr := chi.URLParam(req, "product_id")
	deleteProduct.ProductId, _ = strconv.Atoi(productIdStr)

	// Check if product exists in database
	// if !database.ProductIdExists(deleteProduct.ProductId) {
	// 	utils.ResponseJson(w, http.StatusNotFound, "Product does not exist in database. Please try again.")
	// 	return
	// }

	err := database.DeleteProductFromProducts(deleteProduct.ProductId)
	if err != nil {
		utils.InternalServerError(w, "Internal Server Error in deleting product from products table: ", err)
		return
	}
	
	utils.ResponseJson(w, http.StatusOK, "Successfully Deleted Product!")

}



