package routes

import (
	"net/http"
	"strings"
	"testing"

	"github.com/go-chi/chi"
)

func Test_Routes(t *testing.T) {
	// mux := Routes()

	// // Type Assertion of mux to chiRoutes
	// // Able to access the routing tree
	// chiRoutes := mux.(chi.Routes)

	// var registered = []struct {
	// 	route  string
	// 	method string
	// }{
	// 	{"/login", "POST"},
	// 	{"/logout", "GET"},
	// 	{"/signup", "POST"},
	// 	{"/admin/create-user", "POST"},
	// 	{"/admin/users", "GET"},
	// 	{"/admin/update-user", "PATCH"},
	// 	{"/admin/delete-user", "DELETE"},
	// 	{"/admin/create-user-group", "POST"},
	// 	{"/admin/create-organisation", "POST"},
	// 	{"/product/create", "POST"},
	// 	{"/product/products", "GET"},
	// 	{"/product/update/{product_id}", "PATCH"},
	// 	{"/product/delete/{product_id}", "DELETE"},
	// 	{"/product/create-brand", "POST"},
	// 	{"/product/create-colour", "POST"},
	// 	{"/product/create-category", "POST"},
	// 	{"/product/create-size", "POST"},
	// }

	// // Check to see if the route exists in the list of registered routes (chi.Walk)
	// for _, route := range registered {
	// 	if !routeExists(route.route, route.method, chiRoutes) {
	// 		t.Errorf("Route %q is not registered", route.route)
	// 	}
	// }
}

func routeExists(testRoute, testMethod string, chiRoutes chi.Routes) bool {
	found := false

	// chi.Walk allows us to iterate over all the registered routes in our routing tree and examine each one individually.
	_ = chi.Walk(chiRoutes, func(method string, route string, handler http.Handler, middleware ...func(http.Handler) http.Handler) error {
		if strings.EqualFold(method, testMethod) && strings.EqualFold(route, testRoute) {
			found = true
		}
		return nil
	})

	return found
}
