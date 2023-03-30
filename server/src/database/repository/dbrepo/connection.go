package dbrepo

import (
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/jackc/pgx/v5"
	_ "github.com/lib/pq"
)

const dbTimeout = time.Second * 3

type PostgresDBRepo struct {
	DB *pgxpool.Pool
}

type DB struct{}

func New() *DB {
	return &DB{}
}

func (m *PostgresDBRepo) Connection() *pgx.Conn {
	return m.DB
}
