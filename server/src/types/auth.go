package types

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/LeonLow97/inventory-management-system-golang-react-postgresql/utils"
)

type LoginJSON struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (auth *LoginJSON) ReadJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(auth)
}

func (auth *LoginJSON) AuthFieldsTrimSpaces() *LoginJSON {
	auth.Username = strings.TrimSpace(auth.Username)
	auth.Password = strings.TrimSpace(auth.Password)

	return auth
}

func (auth *LoginJSON) AuthFormValidation(w http.ResponseWriter) error {

	if len(auth.Username) == 0 || len(auth.Password) == 0 {
		return utils.ApiError{Err: "Please fill in all fields - Username and Password.", Status: http.StatusBadRequest}
	}

	// Sanitise Username
	if !utils.CheckLengthRange(auth.Username, 5, 50) || !utils.CheckUsernameSpecialChar(auth.Username) {
		return utils.ApiError{Err: "You have entered an incorrect username/password. Please try again.", Status: http.StatusUnauthorized}
	}

	// Sanitise Password
	if !utils.CheckLengthRange(auth.Password, 8, 20) || !utils.CheckPasswordSpecialChar(auth.Password) {
		return utils.ApiError{Err: "You have entered an incorrect username/password. Please try again.", Status: http.StatusUnauthorized}
	}

	return nil
}
