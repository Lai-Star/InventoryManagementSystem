package route

import (
	"net/http"

	"github.com/LeonLow97/inventory-management-system-golang-react-postgresql/api/controller"
)

func Routes() {
	http.HandleFunc("/login", controller.Login);
	http.HandleFunc("/logout", controller.Logout);
}