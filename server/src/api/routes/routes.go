package routes

import (
	"net/http"

	products "github.com/LeonLow97/inventory-management-system-golang-react-postgresql/api/handlers/products"
	handlers_admin "github.com/LeonLow97/inventory-management-system-golang-react-postgresql/api/handlers/user-management/admin"
	handlers_auth "github.com/LeonLow97/inventory-management-system-golang-react-postgresql/api/handlers/user-management/auth"
	app_middleware "github.com/LeonLow97/inventory-management-system-golang-react-postgresql/appMiddleware"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func Routes() http.Handler {

	mux := chi.NewRouter()

	// Register middleware
	mux.Use(middleware.Recoverer)
	mux.Use(app_middleware.AddIPToContext)
	mux.Use(app_middleware.RequestMiddleware)

	// User Management Routes
	mux.Post("/login", handlers_auth.Login)
	mux.Get("/logout", handlers_auth.Logout)
	mux.Post("/signup", handlers_auth.SignUp)

	// Admin Management Routes
	mux.Route("/admin", func(mux chi.Router) {
		mux.Post("/create-user", handlers_admin.AdminCreateUser)
		mux.Get("/users", handlers_admin.AdminGetUsers)
		mux.Patch("/update-user", handlers_admin.AdminUpdateUser)
		mux.Delete("/delete-user", handlers_admin.AdminDeleteUser)
		mux.Post("/create-user-group", handlers_admin.AdminCreateUserGroup)
		mux.Post("/create-organisation", handlers_admin.AdminCreateOrganisation)
	})

	// Product Routes
	mux.Route("/product", func(mux chi.Router) {
		mux.Post("/create", products.CreateProduct)
		mux.Get("/products", products.GetProducts)
		mux.Patch("/update/{product_id}", products.UpdateProduct)
		mux.Delete("/delete/{product_id}", products.DeleteProduct)
		mux.Post("/create-brand", products.CreateBrand)
		mux.Post("/create-colour", products.CreateColour)
		mux.Post("/create-category", products.CreateCategory)
		mux.Post("/create-size", products.CreateSize)
	})

	return mux
}
