package routes

import (
	"net/http"

	products "github.com/LeonLow97/inventory-management-system-golang-react-postgresql/api/handlers/products"
	admin "github.com/LeonLow97/inventory-management-system-golang-react-postgresql/api/handlers/user-management/admin"
	"github.com/LeonLow97/inventory-management-system-golang-react-postgresql/api/handlers/user-management/auth"
	"github.com/LeonLow97/inventory-management-system-golang-react-postgresql/api/handlers/user-management/user"
	"github.com/LeonLow97/inventory-management-system-golang-react-postgresql/database/repository"
	appMiddleware "github.com/LeonLow97/inventory-management-system-golang-react-postgresql/middleware"
	"github.com/LeonLow97/inventory-management-system-golang-react-postgresql/utils"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func Routes(app repository.DatabaseRepo) http.Handler {

	mux := chi.NewRouter()

	// Register middleware
	mux.Use(middleware.Recoverer)
	mux.Use(appMiddleware.AddIPToContext)
	mux.Use(appMiddleware.RequestMiddleware)

	// Authentication Routes
	loginH := auth.New(app)
	logoutH := auth.New(app)
	mux.Post("/login", utils.MakeHTTPHandler(loginH.Login))
	mux.Get("/logout", utils.MakeHTTPHandler(logoutH.Logout))

	// User Management Routes
	signUpH := user.New(app)
	mux.Post("/signup", utils.MakeHTTPHandler(signUpH.SignUp))

	// Admin Management Routes
	mux.Route("/admin", func(mux chi.Router) {
		mux.Post("/create-user", admin.AdminCreateUser)
		mux.Get("/users", admin.AdminGetUsers)
		mux.Patch("/update-user", admin.AdminUpdateUser)
		mux.Delete("/delete-user", admin.AdminDeleteUser)
		mux.Post("/create-user-group", admin.AdminCreateUserGroup)
		mux.Post("/create-organisation", admin.AdminCreateOrganisation)
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
