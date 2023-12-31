package admin

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/LeonLow97/inventory-management-system-golang-react-postgresql/api/handlers/user-management/auth"
	"github.com/LeonLow97/inventory-management-system-golang-react-postgresql/types"
	"github.com/LeonLow97/inventory-management-system-golang-react-postgresql/utils"
)

func (app application) CreateOrganisation(w http.ResponseWriter, req *http.Request) error {

	if req.Method != http.MethodPost {
		return utils.ApiError{Err: "Invalid Method", Status: http.StatusMethodNotAllowed}
	}

	var newOrg types.AdminCreateOrganisationJSON

	if err := newOrg.ReadJSON(req.Body); err != nil {
		log.Println("newOrg.ReadJSON:", err)
		return utils.ApiError{Err: "Internal Server Error", Status: http.StatusInternalServerError}
	}

	// Setting timeout to follow SLA
	ctx := req.Context()
	ctx, cancel := context.WithTimeout(ctx, 2*time.Minute)
	defer cancel()

	// Check User Group Admin
	if err := auth.RetrieveIssuer(w, req); err != nil {
		return err
	}
	utilsApp := utils.Application{DB: app.DB}
	err := utils.InjectUG(utilsApp, ctx, w.Header().Get("username"), "Admin")
	if err != nil {
		return err
	}

	newOrg.OrgFieldsTrimSpaces()
	if err := newOrg.OrgFormValidation(w); err != nil {
		return err
	}

	// Check if organisation name already exists
	orgNameCount, err := app.DB.GetCountByOrganisationName(ctx, newOrg.OrganisationName)
	if err != nil {
		return utils.ApiError{Err: "Internal Server Error", Status: http.StatusInternalServerError}
	}
	if orgNameCount == 1 {
		return utils.ApiError{Err: "Organisation Name " + newOrg.OrganisationName + " already exist. Please try again.", Status: http.StatusBadRequest}
	}

	// Insert organisation name into organisations table
	err = app.DB.InsertIntoOrganisations(ctx, newOrg.OrganisationName)
	if err != nil {
		return utils.ApiError{Err: "Internal Server Error", Status: http.StatusInternalServerError}
	}

	return utils.WriteJSON(w, http.StatusCreated, utils.ApiSuccess{Success: "Successfully created a new organisation '" + newOrg.OrganisationName + "' !", Status: http.StatusCreated})
}
