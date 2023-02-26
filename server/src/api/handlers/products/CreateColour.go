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

type CreateColourJson struct {
	ColourName string `json:"colour_name"`
}

func CreateColour(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json");
	var newColour CreateColourJson

	// Reading the request body and UnMarshal the body to the CreateColourJson struct
	bs, _ := io.ReadAll(req.Body);
	if err := json.Unmarshal(bs, &newColour); err != nil {
		utils.InternalServerError(w, "Internal Server Error in UnMarshal JSON body in CreateColour route: ", err)
		return;
	}

	// CheckUserGroup: IMS User and Operations
	if !handlers_user_management.RetrieveIssuer(w, req) {return}
	if !utils.CheckUserGroup(w, w.Header().Get("username"), "InvenNexus User", "Operations") {return}

	// Trim White Spaces in colour name
	newColour.ColourName = strings.TrimSpace(newColour.ColourName)

	// Colour Name Form Validation
	if !ProductColourFormValidation(w, newColour.ColourName, "CREATE") {return}

	// Check User Organisation
	username := w.Header().Get("username")
	organisationName, userId, err := database.GetOrganisationNameByUsername(username)
	if err != nil {
		utils.InternalServerError(w, "Internal server error in getting company name from database: ", err)
		return
	}

	// Check category name to see if it already exists in database (cannot have duplicates within the same organisation or username)
	var count int
	if organisationName == "InvenNexus" {
		// Check the colour name for duplicates based on username
		count, err = database.GetColourNameCountByUsername(userId, newColour.ColourName)
	} else {
		// Check the colour name for duplicates based on organisation name
		count, err = database.GetColourNameCountByOrganisation(organisationName, newColour.ColourName)
	}

	if err != nil {
		utils.InternalServerError(w, "Internal server error in getting colour name count: ", err)
		return
	}
	if count >= 1 {
		utils.ResponseJson(w, http.StatusBadRequest, "Colour Name already exists. Please try again.")
		return
	}

	// Insert colour name to `organisation_colours` or `user_colours` tables
	if organisationName == "InvenNexus" {
		// insert into `user_colours` table
		err = database.InsertIntoUserColours(userId, newColour.ColourName)
	} else {
		err = database.InsertIntoOrganisationColours(organisationName, newColour.ColourName)
	}

	if err != nil {
		utils.InternalServerError(w, "Internal server error in inserting colour name into database: ", err)
		return
	}

	utils.ResponseJson(w, http.StatusOK, "Successfully created a colour!")
}