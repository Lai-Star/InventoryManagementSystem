package repository

import (
	"github.com/jackc/pgx/v4"
)

type DatabaseRepo interface {
	Connection() *pgx.Conn
}
