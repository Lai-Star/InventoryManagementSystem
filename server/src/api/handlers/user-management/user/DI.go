package user

import "github.com/LeonLow97/inventory-management-system-golang-react-postgresql/database/repository"

type application struct {
	DB repository.DatabaseRepo
}

func New(app repository.DatabaseRepo) *application {
	return &application{}
}
