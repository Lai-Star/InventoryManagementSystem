package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/LeonLow97/inventory-management-system-golang-react-postgresql/api/routes"
	"github.com/LeonLow97/inventory-management-system-golang-react-postgresql/database"
	"github.com/LeonLow97/inventory-management-system-golang-react-postgresql/keys"
	"github.com/LeonLow97/inventory-management-system-golang-react-postgresql/utils"
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
	routes.Routes()

	// Loading the .env file in the config folder
	err := godotenv.Load("../config/.env");
	utils.CheckError(err);

	// Connecting to localhost
	port := fmt.Sprintf(":%v", os.Getenv("SERVER_PORT"))
	fmt.Println("Server is running on port", port, "!")
	err = http.ListenAndServe(port, nil)
	utils.CheckError(err);

}