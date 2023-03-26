package handlers

import (
	"time"

	"github.com/jackc/pgx/v5"
)

const dbTimeout = time.Second * 3

type PostgresDBRepo struct {
	DB *pgx.Conn
}
