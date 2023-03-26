package admin

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"

	auth_management "github.com/LeonLow97/inventory-management-system-golang-react-postgresql/api/handlers/user-management"
	"github.com/LeonLow97/inventory-management-system-golang-react-postgresql/utils"
)

type AdminCreateOrganisationJson struct {
	OrganisationName string `json:"organisation_name"`
}

func AdminCreateOrganisation(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var organisation AdminCreateOrganisationJson

	// Reading the request body and UnMarshal the body to the AdminCreateOrganisationJson struct
	bs, _ := io.ReadAll(req.Body)
	if err := json.Unmarshal(bs, &organisation); err != nil {
		utils.ResponseJson(w, http.StatusInternalServerError, "Internal Server Error")
		log.Println("Internal Server Error in UnMarshal JSON body in AdminCreateUser route:", err)
		return
	}

	// Check User Group Admin
	if !auth_management.RetrieveIssuer(w, req) {
		return
	}
	if !utils.CheckUserGroup(w, w.Header().Get("username"), "Admin") {
		return
	}

	organisationName := organisation.OrganisationName

	// trim organisation name
	organisationName = strings.TrimSpace(organisationName)

	// Organisation form validation
	isValidOrganisationName := auth_management.OrganisationFormValidation(w, organisationName, "CREATE_ORGANISATION")
	if !isValidOrganisationName {
		return
	}

	// Check if organisation name already exists
	count, err := database.GetOrganisationNameCount(organisationName)
	if err != nil {
		utils.ResponseJson(w, http.StatusInternalServerError, "Internal Server Error")
		log.Println("Internal server error in getting organisation count:", err)
		return
	}
	if count == 1 {
		utils.ResponseJson(w, http.StatusBadRequest, organisationName+" already exists. Please try again.")
		return
	}

	// Insert organisation name into organisations table
	err = database.InsertIntoOrganisations(organisationName)
	if err != nil {
		utils.ResponseJson(w, http.StatusInternalServerError, "Internal Server Error")
		log.Println("Internal server error in inserting into organisations table:", err)
		return
	}

	utils.ResponseJson(w, http.StatusOK, "Successfully created a new organisation.")
}
