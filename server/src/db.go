package main

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/joho/godotenv"
)

func (app *application) connectToDB() (*pgxpool.Pool, error) {
	// Loading the .env file in the config folder
	if err := godotenv.Load("./config/.env"); err != nil {
		log.Println("Error loading .env file when connecting to PostgreSQL: ", err)
	}

	dsn := url.URL{
		Scheme: "postgres",
		Host:   fmt.Sprintf("%s:%s", os.Getenv("POSTGRESQL_HOST"), os.Getenv("POSTGRESQL_PORT")),
		User:   url.UserPassword(os.Getenv("POSTGRESQL_USER"), os.Getenv("POSTGRESQL_PASSWORD")),
		Path:   os.Getenv("POSTGRESQL_DB"),
	}

	q := dsn.Query()
	q.Add("sslmode", "disable")

	dsn.RawQuery = q.Encode()

	var err error
	conn, err := pgxpool.Connect(context.Background(), dsn.String())
	if err != nil {
		return nil, fmt.Errorf("pgx.Connect %w", err)
	}

	// Ping the connection to Postgres
	if err = conn.Ping(context.Background()); err != nil {
		return nil, fmt.Errorf("conn.Ping %w", err)
	}

	fmt.Println("Successfully connected to PostgreSQL!")

	return conn, err
}
