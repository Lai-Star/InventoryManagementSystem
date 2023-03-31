package types

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/LeonLow97/inventory-management-system-golang-react-postgresql/utils"
)

type AdminUserJSON struct {
	Username         string   `json:"username"`
	Password         string   `json:"password"`
	Email            string   `json:"email"`
	UserGroup        []string `json:"user_groups"`
	OrganisationName string   `json:"organisation_name"`
	IsActive         int      `json:"is_active"`
}

type AdminCreateUserGroupJson struct {
	UserGroup []string `json:"user_group"`
}

type AdminCreateOrganisationJSON struct {
	OrganisationName string `json:"organisation_name"`
}

type AdminCreateUserGroupJSON struct {
	UserGroup   string `json:"user_group"`
	Description string `json:"description"`
}

func (u *AdminUserJSON) ReadJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(u)
}

func (ug *AdminCreateUserGroupJSON) ReadJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(ug)
}

func (org *AdminCreateOrganisationJSON) ReadJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(org)
}

func (u *AdminUserJSON) UserFieldsTrimSpaces() *AdminUserJSON {
	u.Username = strings.TrimSpace(u.Username)
	u.Password = strings.TrimSpace(u.Password)
	u.Email = strings.TrimSpace(u.Email)
	u.OrganisationName = strings.TrimSpace(u.OrganisationName)

	for idx, ug := range u.UserGroup {
		u.UserGroup[idx] = strings.TrimSpace(ug)
	}

	return u
}

func (ug *AdminCreateUserGroupJSON) UGFieldsTrimSpaces() *AdminCreateUserGroupJSON {
	ug.UserGroup = strings.TrimSpace(ug.UserGroup)
	return ug
}

func (org *AdminCreateOrganisationJSON) OrgFieldsTrimSpaces() *AdminCreateOrganisationJSON {
	org.OrganisationName = strings.TrimSpace(org.OrganisationName)
	return org
}

func (u *AdminUserJSON) CreateUserFormValidation(w http.ResponseWriter) error {

	// Username Validation
	switch {
	case utils.IsBlankField(u.Username):
		return utils.ApiError{Err: "Username cannot be blank", Status: http.StatusBadRequest}
	case !utils.CheckLengthRange(u.Username, 5, 50):
		return utils.ApiError{Err: "Username must have a length of 5 - 50 characters. Please try again.", Status: http.StatusBadRequest}
	case !utils.CheckUsernameSpecialChar(u.Username):
		return utils.ApiError{Err: "Username cannot contain special characters except 'underscore'", Status: http.StatusBadRequest}
	}

	// Password Validation
	switch {
	case utils.IsBlankField(u.Password):
		return utils.ApiError{Err: "Password cannot be blank. Please try again.", Status: http.StatusBadRequest}
	case (len(u.Password) > 0) && !utils.CheckLengthRange(u.Password, 8, 20):
		return utils.ApiError{Err: "Password must have a length of 8 - 20 characters. Please try again.", Status: http.StatusBadRequest}
	case (len(u.Password) > 0) && !utils.CheckPasswordSpecialChar(u.Password):
		return utils.ApiError{Err: "Password must contain at least one lowercase, uppercase, number, and special character.", Status: http.StatusBadRequest}
	}

	// Email Validation
	switch {
	case utils.IsBlankField(u.Email):
		return utils.ApiError{Err: "Email address cannot be blank. Please try again.", Status: http.StatusBadRequest}
	case !utils.CheckLengthRange(u.Email, 1, 255):
		return utils.ApiError{Err: "Email address has a maximum length of 255 characters. Please try again.", Status: http.StatusBadRequest}
	case !utils.CheckEmailAddress(u.Email):
		return utils.ApiError{Err: "Email address is not in the correct format. Please try again.", Status: http.StatusBadRequest}
	}

	// Organisation Name Validation
	switch {
	case utils.IsBlankField(u.OrganisationName):
		return utils.ApiError{Err: "Organisation name cannot be blank.", Status: http.StatusBadRequest}
	case !utils.CheckLengthRange(u.OrganisationName, 1, 255):
		return utils.ApiError{Err: "Organisation name has a maximum length of 255 characters.", Status: http.StatusBadRequest}
	}

	// User Group Validation
	for _, ug := range u.UserGroup {
		if !utils.CheckLengthRange(ug, 1, 255) {
			return utils.ApiError{Err: "User Group cannot have a length of more than 255 characters. Please try again.", Status: http.StatusBadRequest}
		}
	}

	return nil
}

func (u *AdminUserJSON) UpdateUserFormValidation(w http.ResponseWriter) error {

	// User Validation
	switch {
	case utils.IsBlankField(u.Username):
		return utils.ApiError{Err: "Username cannot be blank", Status: http.StatusBadRequest}
	case !utils.CheckLengthRange(u.Username, 5, 50):
		return utils.ApiError{Err: "Username must have a length of 5 - 50 characters. Please try again.", Status: http.StatusBadRequest}
	case !utils.CheckUsernameSpecialChar(u.Username):
		return utils.ApiError{Err: "Username cannot contain special characters except 'underscore'", Status: http.StatusBadRequest}
	}

	// Password Validation
	if len(u.Password) > 0 {
		switch {
		case (len(u.Password) > 20):
			return utils.ApiError{Err: "Password must have a length of 8 - 20 characters. Please try again.", Status: http.StatusBadRequest}
		case (len(u.Password) > 0) && !utils.CheckPasswordSpecialChar(u.Password):
			return utils.ApiError{Err: "Password must contain at least one lowercase, uppercase, number, and special character.", Status: http.StatusBadRequest}
		}
	}

	// Email Validation
	if len(u.Email) > 0 {
		switch {
		case !utils.CheckLengthRange(u.Email, 1, 255):
			return utils.ApiError{Err: "Email address has a maximum length of 255 characters. Please try again.", Status: http.StatusBadRequest}
		case !utils.CheckEmailAddress(u.Email):
			return utils.ApiError{Err: "Email address is not in the correct format. Please try again.", Status: http.StatusBadRequest}
		}
	}

	// Organisation Name Validation
	if len(u.OrganisationName) > 0 {
		if !utils.CheckLengthRange(u.OrganisationName, 1, 255) {
			return utils.ApiError{Err: "Organisation name has a maximum length of 255 characters.", Status: http.StatusBadRequest}
		}
	}

	// IsActive Validation
	if u.IsActive != 0 && u.IsActive != 1 {
		return utils.ApiError{Err: "Invalid IsActive Status provided. Please try again.", Status: http.StatusBadRequest}
	}

	// User Group Validation
	if len(u.UserGroup) > 0 {
		for _, ug := range u.UserGroup {
			if !utils.CheckLengthRange(ug, 1, 255) {
				return utils.ApiError{Err: "User Group cannot have a length of more than 255 characters. Please try again.", Status: http.StatusBadRequest}
			}
		}
	}

	return nil

}

func (ug *AdminCreateUserGroupJSON) UGFormValidation(w http.ResponseWriter) error {

	// User group validation
	switch {
	case utils.IsBlankField(ug.UserGroup):
		return utils.ApiError{Err: "User Group cannot be blank.", Status: http.StatusBadRequest}
	case !utils.CheckLengthRange(ug.UserGroup, 1, 255):
		return utils.ApiError{Err: "User Group has a maximum length of 255 characters.", Status: http.StatusBadRequest}
	}

	return nil
}

func (org *AdminCreateOrganisationJSON) OrgFormValidation(w http.ResponseWriter) error {

	// Organisation name validation
	switch {
	case utils.IsBlankField(org.OrganisationName):
		return utils.ApiError{Err: "Organisation name cannot be blank.", Status: http.StatusBadRequest}
	case !utils.CheckLengthRange(org.OrganisationName, 1, 255):
		return utils.ApiError{Err: "Organisation name has a maximum length of 255 characters.", Status: http.StatusBadRequest}
	}

	return nil
}
