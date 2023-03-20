package products

import (
	"log"
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

	// CheckUserGroup: IMS User and Operations
	if !handlers_user_management.RetrieveIssuer(w, req) {
		return
	}
	if !utils.CheckUserGroup(w, w.Header().Get("username"), "InvenNexus User", "Operations") {
		return
	}

	// Get productid from url params
	productIdStr := chi.URLParam(req, "product_id")
	productId, _ := strconv.Atoi(productIdStr)

	// Check User Organisation
	username := w.Header().Get("username")
	organisationName, userId, err := database.GetOrganisationNameAndUserIdByUsername(username)
	if err != nil {
		utils.ResponseJson(w, http.StatusInternalServerError, "Internal Server Error")
		log.Println("Internal Server Error in getting company name from database:", err)
		return
	}

	var count int
	// Check if product exists in database
	if organisationName == "InvenNexus" {
		count, err = database.GetCountByUserIdAndProductId(userId, productId)
	} else {
		count, err = database.GetCountByOrganisationNameAndProductId(organisationName, productId)
	}

	if err != nil {
		utils.ResponseJson(w, http.StatusInternalServerError, "Internal Server Error")
		log.Println("Internal server error in getting count by organisation/user id and product id:", err)
		return
	}
	if count == 0 {
		utils.ResponseJson(w, http.StatusNotFound, "There is no such product. Please try again.")
		return
	}

	// Delete product from products table
	err = database.DeleteProductByID(productId)
	if err != nil {
		utils.ResponseJson(w, http.StatusInternalServerError, "Internal Server Error")
		log.Println("Internal server error in deleting product by product id:", err)
		return
	}

	utils.ResponseJson(w, http.StatusOK, "Successfully Deleted Product!")
}
