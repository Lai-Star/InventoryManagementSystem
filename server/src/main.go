package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/LeonLow97/inventory-management-system-golang-react-postgresql/api/routes"
	"github.com/LeonLow97/inventory-management-system-golang-react-postgresql/database"
	"github.com/LeonLow97/inventory-management-system-golang-react-postgresql/keys"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func main() {

	db := database.ConnectToPostgreSQL();

	// Close the database to prevent data leak
	defer db.Close()

	// Generating & validating the public and private keys for signed Json
	keys.GenerateKeys()
	keys.CheckKeys()

	// get application routes
	mux := routes.Routes()

	// Loading the .env file in the config folder
	err := godotenv.Load("./config/.env");
	if err != nil {
		log.Println("Error loading .env file in main.go: ", err)
	}

	// Connecting to localhost
	port := fmt.Sprintf(":%v", os.Getenv("SERVER_PORT"))
	fmt.Println("Server is running on port", port, "!")

	// Enabling CORS for all routes
	clientOrigin := os.Getenv("CLIENT_ORIGIN")
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{clientOrigin},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
		AllowedHeaders:   []string{"Content-Type"},
		ExposedHeaders: []string{"Content-Length"},
		AllowOriginFunc: func(origin string) bool {
			return origin == clientOrigin
		},
		MaxAge: int(12 * time.Hour.Seconds()),
	})
	handler := c.Handler(mux)

	err = http.ListenAndServe(port, handler)
	if err != nil {
		log.Println("Error in listening to port in main.go: ",err)
	}
}