package handlers

import (
	"time"

	"github.com/jackc/pgx/v4"
)

const dbTimeout = time.Second * 3

type PostgresDBRepo struct {
	DB *pgx.Conn
}
