package dbrepo

import (
	"time"

	"github.com/jackc/pgx/v5"
	_ "github.com/lib/pq"
)

const dbTimeout = time.Second * 3

type PostgresDBRepo struct {
	DB *pgx.Conn
}

type DB struct{}

func New() *DB {
	return &DB{}
}

func (m *PostgresDBRepo) Connection() *pgx.Conn {
	return m.DB
}
