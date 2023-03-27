package types

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/LeonLow97/inventory-management-system-golang-react-postgresql/utils"
)

type SignUpJSON struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func (u *SignUpJSON) ReadJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(u)
}

func (user *SignUpJSON) UserFieldsTrimSpaces() *SignUpJSON {
	user.Username = strings.TrimSpace(user.Username)
	user.Password = strings.TrimSpace(user.Password)
	user.Email = strings.TrimSpace(user.Email)

	return user
}

func (user *SignUpJSON) SignUpFormValidation(w http.ResponseWriter) error {

	// Username Validation
	switch {
	case utils.IsBlankField(user.Username):
		return utils.ApiError{Err: "Username cannot be blank", Status: http.StatusBadRequest}
	case !utils.CheckLengthRange(user.Username, 5, 50):
		return utils.ApiError{Err: "Username must have a length of 5 - 50 characters. Please try again.", Status: http.StatusBadRequest}
	case !utils.CheckUsernameSpecialChar(user.Username):
		return utils.ApiError{Err: "Username cannot contain special characters except 'underscore'", Status: http.StatusBadRequest}
	}

	// Password Validation
	switch {
	case utils.IsBlankField(user.Password):
		return utils.ApiError{Err: "Password cannot be blank. Please try again.", Status: http.StatusBadRequest}
	case (len(user.Password) > 0) && !utils.CheckLengthRange(user.Password, 8, 20):
		return utils.ApiError{Err: "Password must have a length of 8 - 20 characters. Please try again.", Status: http.StatusBadRequest}
	case (len(user.Password) > 0) && !utils.CheckPasswordSpecialChar(user.Password):
		return utils.ApiError{Err: "Password must contain at least one lowercase, uppercase, number, and special character.", Status: http.StatusBadRequest}
	}

	// Email Validation
	switch {
	case utils.IsBlankField(user.Email):
		return utils.ApiError{Err: "Email address cannot be blank. Please try again.", Status: http.StatusBadRequest}
	case !utils.CheckLengthRange(user.Email, 1, 255):
		return utils.ApiError{Err: "Email address has a maximum length of 255 characters. Please try again.", Status: http.StatusBadRequest}
	case !utils.CheckEmailAddress(user.Email):
		return utils.ApiError{Err: "Email address is not in the correct format. Please try again.", Status: http.StatusBadRequest}
	}

	return nil
}
