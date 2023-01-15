package handlers_admin

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/LeonLow97/inventory-management-system-golang-react-postgresql/utils"
)

func AdminDeleteUser(w http.ResponseWriter, req *http.Request) {
	// Set Headers
	w.Header().Set("Content-Type", "application/json");
	var adminDeleteUser AdminDeleteUserMgmt

	// Reading the request body and UnMarshal the body to the AdminUserMgmt struct
	bs, _ := io.ReadAll(req.Body);
	if err := json.Unmarshal(bs, &adminDeleteUser); err != nil {
		utils.InternalServerError(w, "Internal Server Error in Unmarshal JSON body in AdminDeleteUser: ", err)
		return;
	}
}