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
	"github.com/go-chi/cors"
)

func Routes(app repository.DatabaseRepo) http.Handler {

	r := chi.NewRouter()

	// Basic CORS
	// for more ideas, see: https://developer.github.com/v3/#cross-origin-resource-sharing
	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	// Register middleware
	r.Use(middleware.Recoverer)
	r.Use(appMiddleware.AddIPToContext)
	r.Use(appMiddleware.RequestMiddleware)

	// Authentication Routes
	loginH := auth.New(app)
	logoutH := auth.New(app)
	r.Post("/login", utils.MakeHTTPHandler(loginH.Login))
	r.Get("/logout", utils.MakeHTTPHandler(logoutH.Logout))

	// User Management Routes
	signUpH := user.New(app)
	r.Post("/signup", utils.MakeHTTPHandler(signUpH.SignUp))

	// Admin Management Routes
	createOrgH := admin.New(app)
	createUgH := admin.New(app)
	createUserH := admin.New(app)
	updateUserH := admin.New(app)
	getUserH := admin.New(app)
	deleteUserH := admin.New(app)
	r.Route("/admin", func(r chi.Router) {
		r.Post("/create-user", utils.MakeHTTPHandler(createUserH.CreateUser))
		r.Get("/users", utils.MakeHTTPHandler(getUserH.GetUsers))
		r.Patch("/update-user", utils.MakeHTTPHandler(updateUserH.UpdateUser))
		r.Delete("/delete-user", utils.MakeHTTPHandler(deleteUserH.DeleteUser))
		r.Post("/create-user-group", utils.MakeHTTPHandler(createUgH.CreateUserGroup))
		r.Post("/create-organisation", utils.MakeHTTPHandler(createOrgH.CreateOrganisation))
	})

	// Product Routes
	r.Route("/product", func(r chi.Router) {
		r.Post("/create", products.CreateProduct)
		r.Get("/products", products.GetProducts)
		r.Patch("/update/{product_id}", products.UpdateProduct)
		r.Delete("/delete/{product_id}", products.DeleteProduct)
		r.Post("/create-brand", products.CreateBrand)
		r.Post("/create-colour", products.CreateColour)
		r.Post("/create-category", products.CreateCategory)
		r.Post("/create-size", products.CreateSize)
	})

	return r
}
