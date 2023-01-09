package routes

import (
	"net/http"

	"github.com/LeonLow97/inventory-management-system-golang-react-postgresql/api/controller"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func Routes() http.Handler {

	mux := chi.NewRouter()

	// Register middleware
	mux.Use(middleware.Recoverer)

	// Register Routes
	mux.Post("/login", controller.Login)
	mux.Get("/logout", controller.Logout)
	mux.Post("/signup", controller.SignUp)

	return mux
}