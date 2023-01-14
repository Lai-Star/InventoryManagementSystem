package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/LeonLow97/inventory-management-system-golang-react-postgresql/api/routes"
	"github.com/LeonLow97/inventory-management-system-golang-react-postgresql/database"
	"github.com/LeonLow97/inventory-management-system-golang-react-postgresql/keys"
	"github.com/joho/godotenv"
)

func main() {

	db := database.ConnectToPostgreSQL();

	// Close the database to prevent data leak
	defer db.Close()

	// Generating & validating the public and private keys for signed Json
	// keys.GenerateKeys()
	keys.CheckKeys()

	// get application routes
	mux := routes.Routes()

	// Loading the .env file in the config folder
	err := godotenv.Load("../config/.env");
	if err != nil {
		log.Println("Error loading .env file in main.go: ", err)
	}

	// Connecting to localhost
	port := fmt.Sprintf(":%v", os.Getenv("SERVER_PORT"))
	fmt.Println("Server is running on port", port, "!")
	err = http.ListenAndServe(port, mux)
	if err != nil {
		log.Println("Error in listening to port in main.go: ",err)
	}
}