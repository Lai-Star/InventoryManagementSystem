package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/LeonLow97/inventory-management-system-golang-react-postgresql/api/routes"
	"github.com/LeonLow97/inventory-management-system-golang-react-postgresql/database/repository"
	"github.com/LeonLow97/inventory-management-system-golang-react-postgresql/database/repository/dbrepo"
	"github.com/LeonLow97/inventory-management-system-golang-react-postgresql/keys"
	"github.com/joho/godotenv"
)

type application struct {
	DB repository.DatabaseRepo
}

func main() {

	// Set up an app config
	app := application{}

	conn, err := app.connectToDB()
	if err != nil {
		log.Fatal("Failed to connect to PostgreSQL:", err)
	}

	// Close the database to prevent data leak
	defer func() {
		conn.Close()
		fmt.Println("Closed database connection.")
	}()

	// Passing the PostgreSQL connection into PostgresDBRepo struct
	app.DB = &dbrepo.PostgresDBRepo{DB: conn}

	// Generating & validating the public and private keys for signed Json
	// keys.GenerateKeys()
	keys.CheckKeys()

	// get application routes (passing in conn for dependency injection)
	mux := routes.Routes(app.DB)

	// Loading the .env file in the config folder
	err = godotenv.Load("./config/.env")
	if err != nil {
		log.Println("Error loading .env file in main.go: ", err)
	}

	// Connecting to localhost
	port := fmt.Sprintf(":%v", os.Getenv("SERVER_PORT"))
	fmt.Println("Server is running on port", port, "!")

	err = http.ListenAndServe(port, mux)
	if err != nil {
		log.Println("Error in listening to port in main.go: ", err)
	}
}
