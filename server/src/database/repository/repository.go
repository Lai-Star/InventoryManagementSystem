package repository

import "github.com/jackc/pgx/v5"

type DatabaseRepo interface {
	Connection() *pgx.Conn
	GetCountByUsername(username string) (int, error)
}
