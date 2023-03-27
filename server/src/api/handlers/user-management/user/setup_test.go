package user

import (
	"os"
	"testing"

	"github.com/LeonLow97/inventory-management-system-golang-react-postgresql/database/repository/dbrepo"
)

func TestMain(m *testing.M) {
	app.DB = &dbrepo.TestDBRepo{}

	os.Exit(m.Run())
}
