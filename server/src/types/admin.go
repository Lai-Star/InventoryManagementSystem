package types

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/LeonLow97/inventory-management-system-golang-react-postgresql/utils"
)

type AdminCreateOrganisationJSON struct {
	OrganisationName string `json:"organisation_name"`
}

func (org *AdminCreateOrganisationJSON) ReadJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(org)
}

func (org *AdminCreateOrganisationJSON) OrgFieldsTrimSpaces() *AdminCreateOrganisationJSON {
	org.OrganisationName = strings.TrimSpace(org.OrganisationName)
	return org
}

func (org *AdminCreateOrganisationJSON) OrgFormValidation(w http.ResponseWriter) error {

	// Organisation name validation
	switch {
	case utils.IsBlankField(org.OrganisationName):
		return utils.ApiError{Err: "Organisation name cannot be blank", Status: http.StatusBadRequest}
	case !utils.CheckLengthRange(org.OrganisationName, 1, 255):
		return utils.ApiError{Err: "Organisation name has a maximum length of 255 characters.", Status: http.StatusBadRequest}
	}

	return nil
}
