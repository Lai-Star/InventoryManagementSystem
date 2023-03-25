package dbrepo

import (
	"time"

	"github.com/jackc/pgx/v4"
	_ "github.com/lib/pq"
)

const dbTimeout = time.Second * 3

type PostgresDBRepo struct {
	DB *pgx.Conn
}

// For querying in other files inside database folder
var conn *pgx.Conn

func (m *PostgresDBRepo) Connection() *pgx.Conn {
	return m.DB
}
