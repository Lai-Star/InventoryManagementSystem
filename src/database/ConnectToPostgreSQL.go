package database

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	"github.com/LeonLow97/inventory-management-system-golang-react-postgresql/utils"
)

// For querying in other files inside database folder
var db *sql.DB

func ConnectToPostgreSQL() *sql.DB {
	// Loading the .env file in the config folder
	err := godotenv.Load("../config/.env");
	utils.CheckError(err);

	host := os.Getenv("POSTGRESQL_HOST")
	port := os.Getenv("POSTGRESQL_PORT")
	user := os.Getenv("POSTGRESQL_USER")
	password := os.Getenv("POSTGRESQL_PASSWORD")
	dbname := os.Getenv("POSTGRESQL_DB")

	pgsqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
	host, port, user, password, dbname)

	db, err = sql.Open("postgres", pgsqlInfo)
	utils.CheckError(err)
	
	// Test the connection to the database
	err = db.Ping()
	utils.CheckError(err)

	fmt.Println("Successfully connected to PostgreSQL database!")

	return db
}

var (
	InsertIntoUser = "INSERT INTO users (username, password, email, isActive, added_date, updated_date) VALUES ($1, $2, $3, $4, now(), now());"
)

func CreateNewUser(username string, password string, email string, isActive int) error {
	_, err := db.Exec(InsertIntoUser, username, password, email, isActive)
	utils.CheckError(err)
	return err
}