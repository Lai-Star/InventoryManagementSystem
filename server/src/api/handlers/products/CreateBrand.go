package handlers_products

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"

	handlers_user_management "github.com/LeonLow97/inventory-management-system-golang-react-postgresql/api/handlers/user-management"
	"github.com/LeonLow97/inventory-management-system-golang-react-postgresql/database"
	"github.com/LeonLow97/inventory-management-system-golang-react-postgresql/utils"
)

type CreateBrandJson struct {
	BrandName string `json:"brand_name"`
}

func CreateBrand(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json");
	var newBrand CreateBrandJson

	// Reading the request body and UnMarshal the body to the CreateBrandJson struct
	bs, _ := io.ReadAll(req.Body);
	if err := json.Unmarshal(bs, &newBrand); err != nil {
		utils.InternalServerError(w, "Internal Server Error in UnMarshal JSON body in CreateBrand route: ", err)
		return;
	}

	// CheckUserGroup: IMS User and Operations
	if !handlers_user_management.RetrieveIssuer(w, req) {return}
	if !utils.CheckUserGroup(w, w.Header().Get("username"), "InvenNexus User", "Operations") {return}

	// Trim White Spaces in brand name
	newBrand.BrandName = strings.TrimSpace(newBrand.BrandName)

	// Brand Name Form Validation
	if !ProductBrandFormValidation(w, newBrand.BrandName) {return}

	// Check User Organisation
	username := w.Header().Get("username")
	organisationName, userId, err := database.GetOrganisationNameByUsername(username)
	if err != nil {
		utils.InternalServerError(w, "Internal server error in getting company name from database: ", err)
		return
	}

	// Check brand name to see if it already exists in database (cannot have duplicates within the same organisation or username)
	var count int
	if organisationName == "InvenNexus" {
		// Check the brand name for duplicates based on username
		count, err = database.GetBrandNameCountByUsername(userId, newBrand.BrandName)
	} else {
		// Check the brand name for duplicates based on organisation name
		count, err = database.GetBrandNameCountByOrganisation(organisationName, newBrand.BrandName)
	}

	if err != nil {
		utils.InternalServerError(w, "Internal server error in getting brand name count: ", err)
		return
	}
	if count >= 1 {
		utils.ResponseJson(w, http.StatusBadRequest, "Product Brand Name already exists. Please try again.")
		return
	}

	// Insert brand name to `organisation_brands` or `user_brands` tables
	if organisationName == "InvenNexus" {
		// insert into `user_brands` table
		err = database.InsertIntoUserBrands(userId, newBrand.BrandName)
	} else {
		err = database.InsertIntoOrganisationBrands(organisationName, newBrand.BrandName)
	}

	if err != nil {
		utils.InternalServerError(w, "Internal server error in inserting brand name into database: ", err)
		return
	}

	utils.ResponseJson(w, http.StatusOK, "Successfully created a brand!")
}