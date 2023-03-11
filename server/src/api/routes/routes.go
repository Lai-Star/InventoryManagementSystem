package routes

import (
	"net/http"

	handlers_products "github.com/LeonLow97/inventory-management-system-golang-react-postgresql/api/handlers/products"
	handlers_admin "github.com/LeonLow97/inventory-management-system-golang-react-postgresql/api/handlers/user-management/admin"
	handlers_auth "github.com/LeonLow97/inventory-management-system-golang-react-postgresql/api/handlers/user-management/auth"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func Routes() http.Handler {

	mux := chi.NewRouter()

	// Register middleware
	mux.Use(middleware.Recoverer)

	// User Management Routes
	mux.Post("/login", handlers_auth.Login)
	mux.Get("/logout", handlers_auth.Logout)
	mux.Post("/signup", handlers_auth.SignUp)

	// Admin Management Routes
	mux.Route("/admin", func (mux chi.Router) {
		mux.Post("/create-user", handlers_admin.AdminCreateUser)
		mux.Get("/users", handlers_admin.AdminGetUsers)
		mux.Patch("/update-user", handlers_admin.AdminUpdateUser)
		mux.Delete("/delete-user", handlers_admin.AdminDeleteUser)
		mux.Post("/create-user-group", handlers_admin.AdminCreateUserGroup)
		mux.Post("/create-organisation", handlers_admin.AdminCreateOrganisation)
	})

	// Product Routes
	mux.Route("/product", func(mux chi.Router) {
		mux.Post("/create", handlers_products.CreateProduct)
		mux.Get("/products", handlers_products.GetProducts)
		mux.Patch("/update/{product_id}", handlers_products.UpdateProduct)
		mux.Delete("/delete/{product_id}", handlers_products.DeleteProduct)
		mux.Post("/create-brand", handlers_products.CreateBrand)
		mux.Post("/create-colour", handlers_products.CreateColour)
		mux.Post("/create-category", handlers_products.CreateCategory)
		mux.Post("/create-size", handlers_products.CreateSize)
	})

	return mux
}