package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// For querying in other files inside database folder
var db *sql.DB

func ConnectToPostgreSQL() *sql.DB {
	// Loading the .env file in the config folder
	err := godotenv.Load("../config/.env");
	if err != nil {
		log.Println("Error loading .env file when connecting to PostgreSQL: ", err)
	}

	host := os.Getenv("POSTGRESQL_HOST")
	port := os.Getenv("POSTGRESQL_PORT")
	user := os.Getenv("POSTGRESQL_USER")
	password := os.Getenv("POSTGRESQL_PASSWORD")
	dbname := os.Getenv("POSTGRESQL_DB")

	pgsqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
	host, port, user, password, dbname)

	db, err = sql.Open("postgres", pgsqlInfo)
	if err != nil {
		log.Println("Error opening PostgreSQL: ", err)
	}
	
	// Test the connection to the database
	err = db.Ping()
	if err != nil {
		log.Println("Error pinging PostgreSQL: ", err)
	}

	if err == nil {
		fmt.Println("Successfully connected to PostgreSQL database!")
	}

	return db
}