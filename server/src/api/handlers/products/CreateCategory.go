package products

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"

	handlers_user_management "github.com/LeonLow97/inventory-management-system-golang-react-postgresql/api/handlers/user-management"
	"github.com/LeonLow97/inventory-management-system-golang-react-postgresql/database"
	"github.com/LeonLow97/inventory-management-system-golang-react-postgresql/utils"
)

type CreateCategoryJson struct {
	CategoryName string `json:"category_name"`
}

func CreateCategory(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newCategory CreateCategoryJson

	// Reading the request body and UnMarshal the body to the CreateCategoryJson struct
	bs, _ := io.ReadAll(req.Body)
	if err := json.Unmarshal(bs, &newCategory); err != nil {
		utils.InternalServerError(w, "Internal Server Error in UnMarshal JSON body in CreateCategory route: ", err)
		return
	}

	// CheckUserGroup: IMS User and Operations
	if !handlers_user_management.RetrieveIssuer(w, req) {
		return
	}
	if !utils.CheckUserGroup(w, w.Header().Get("username"), "InvenNexus User", "Operations") {
		return
	}

	// Trim white spaces in category name
	newCategory.CategoryName = strings.TrimSpace(newCategory.CategoryName)

	// Category Name Form Validation
	if !ProductCategoryFormValidation(w, newCategory.CategoryName, "CREATE") {
		return
	}

	// Check User Organisation
	username := w.Header().Get("username")
	organisationName, userId, err := database.GetOrganisationNameAndUserIdByUsername(username)
	if err != nil {
		utils.InternalServerError(w, "Internal server error in getting company name from database: ", err)
		return
	}

	// Check category name to see if it already exists in database (cannot have duplicates within the same organisation or username)
	var count int
	if organisationName == "InvenNexus" {
		// Check the category name for duplicates based on username
		count, err = database.GetCategoryNameCountByUsername(userId, newCategory.CategoryName)
	} else {
		// Check the category name for duplicates based on organisation name
		count, err = database.GetCategoryNameCountByOrganisation(organisationName, newCategory.CategoryName)
	}

	if err != nil {
		utils.InternalServerError(w, "Internal server error in getting category name count: ", err)
		return
	}
	if count >= 1 {
		utils.ResponseJson(w, http.StatusBadRequest, "Category Name already exists. Please try again.")
		return
	}

	// Insert category name to `organisation_sizes` or `user_sizes` tables
	if organisationName == "InvenNexus" {
		// insert into `user_sizes` table
		err = database.InsertIntoUserCategories(userId, newCategory.CategoryName)
	} else {
		err = database.InsertIntoOrganisationCategories(organisationName, newCategory.CategoryName)
	}

	if err != nil {
		utils.InternalServerError(w, "Internal server error in inserting category name into database: ", err)
		return
	}

	utils.ResponseJson(w, http.StatusOK, "Successfully created a category!")
}
