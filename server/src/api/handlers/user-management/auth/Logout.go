package handlers_user

import (
	"net/http"

	handlers_user_management "github.com/LeonLow97/inventory-management-system-golang-react-postgresql/api/handlers/user-management"
	"github.com/LeonLow97/inventory-management-system-golang-react-postgresql/utils"
)

func Logout(w http.ResponseWriter, req *http.Request) {
	handlers_user_management.RetrieveIssuer(w, req)
	// fmt.Println("Retrieved Issuer from Logout", w.Header().Get("username"))

	w.Header().Set("Content-Type", "application/json")
	cookie, err := req.Cookie("leon-jwt-token")
	if err != nil {
		utils.InternalServerError(w, "Internal Server Error in retrieving cookie: ", err)
		return;
	}
	cookie.Value = ""

	cookie = &http.Cookie {
		Name: "leon-jwt-token",
		Value: "",
		MaxAge: -1,
		Path: "",
		Domain: "",
		Secure: false,
		HttpOnly: true,
	}

	http.SetCookie(w, cookie)

	utils.ResponseJson(w, http.StatusOK, "Successfully Logged Out!")
}