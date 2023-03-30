package utils

import (
	"context"
	"net/http"
	"strings"

	"github.com/LeonLow97/inventory-management-system-golang-react-postgresql/database/repository"
)

type Application struct {
	DB repository.DatabaseRepo
}

// Dependency Injection of Repository (Database)
func InjectUG(app Application, ctx context.Context, username string, userGroups ...string) error {
	return app.CheckUserGroup(ctx, username, userGroups...)
}

// Determines if a user has been assigned that usergroup
func (app Application) CheckUserGroup(ctx context.Context, username string, userGroups ...string) error {

	isAuthorizedUser, err := app.DB.GetUserGroupsByUsername(ctx, username, userGroups...)
	if err != nil {
		return ApiError{Err: "Internal Server Error", Status: http.StatusInternalServerError}
	}
	if isAuthorizedUser {
		return nil
	}

	return ApiError{Err: "Access Denied: User does not have permission to access this resource", Status: http.StatusForbidden}
}

func Contains(s []string, e string) bool {
	for _, a := range s {
		if strings.EqualFold(a, e) {
			return true
		}
	}

	return false
}
