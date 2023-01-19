package routes

import (
	"net/http"
	"strings"
	"testing"

	"github.com/go-chi/chi"
)

func Test_Routes(t *testing.T) {
	var registered = []struct {
		route string
		method string
	} {
		{"/login", "POST"},
		{"/logout", "GET"},
		{"/signup", "POST"},
	}

	mux := Routes()

	// Type Assertion of mux to chiRoutes
	chiRoutes := mux.(chi.Routes)

	for _, route := range registered {
		// check to see if the route exists
		if !routeExists(route.route, route.method, chiRoutes) {
			t.Errorf("Route %q is not registered", route.route)
		}
	}
}

func routeExists(testRoute, testMethod string, chiRoutes chi.Routes) bool {
	found := false

	// chi.Walk goes through all the routes that are registered
	_ = chi.Walk(chiRoutes, func(method string, route string, handler http.Handler, middleware ...func(http.Handler) http.Handler) error {
		if strings.EqualFold(method, testMethod) && strings.EqualFold(route, testRoute) {
			found = true
		}
		return nil
	})

	return found
}