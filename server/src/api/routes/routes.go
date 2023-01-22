package routes

import (
	"net/http"

	handlers_admin "github.com/LeonLow97/inventory-management-system-golang-react-postgresql/api/handlers/admin"
	handlers_products "github.com/LeonLow97/inventory-management-system-golang-react-postgresql/api/handlers/products"
	handlers_user "github.com/LeonLow97/inventory-management-system-golang-react-postgresql/api/handlers/user"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func Routes() http.Handler {

	mux := chi.NewRouter()

	// Register middleware
	mux.Use(middleware.Recoverer)

	// User Management Routes
	mux.Post("/login", handlers_user.Login)
	mux.Get("/logout", handlers_user.Logout)
	mux.Post("/signup", handlers_user.SignUp)

	// Admin Management Routes
	mux.Post("/admin-create-user", handlers_admin.AdminCreateUser)
	mux.Get("/admin-get-users", handlers_admin.AdminGetUsers)
	mux.Patch("/admin-update-user", handlers_admin.AdminUpdateUser)
	mux.Delete("/admin-delete-user", handlers_admin.AdminDeleteUser)

	// Product Routes
	mux.Post("/create-product", handlers_products.CreateProduct)
	mux.Get("/get-products", handlers_products.GetProducts)
	mux.Patch("/update-product/{product_id}", handlers_products.UpdateProduct)
	mux.Delete("/delete-product", handlers_products.DeleteProduct)

	return mux
}