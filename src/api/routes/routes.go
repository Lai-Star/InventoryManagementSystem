package routes

import (
	"net/http"

	"github.com/LeonLow97/inventory-management-system-golang-react-postgresql/api/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func Routes() http.Handler {

	mux := chi.NewRouter()

	// Register middleware
	mux.Use(middleware.Recoverer)

	// User Management Routes
	mux.Post("/login", handlers.Login)
	mux.Get("/logout", handlers.Logout)
	mux.Post("/signup", handlers.SignUp)

	// Admin Management Routes
	mux.Post("/admin-create-user", handlers.AdminCreateUser)
	mux.Get("/admin-get-users", handlers.AdminGetUsers)
	mux.Put("/admin-update-user", handlers.AdminUpdateUser)
	mux.Delete("/admin-delete-user", handlers.AdminDeleteUser)

	return mux
}