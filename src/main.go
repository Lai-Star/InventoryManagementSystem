package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/LeonLow97/inventory-management-system-golang-react-postgresql/libs"
	"github.com/joho/godotenv"
)

func main() {

	// Loading the .env file in the config folder
	err := godotenv.Load("../config/.env");
	libs.CheckError(err);

	// Connecting to localhost
	port := fmt.Sprintf(":%v", os.Getenv("SERVER_PORT"))
	fmt.Println("Server is running on port", port)
	err = http.ListenAndServe(port, nil)
	libs.CheckError(err);	
}