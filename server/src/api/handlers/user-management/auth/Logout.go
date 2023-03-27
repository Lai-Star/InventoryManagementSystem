package auth

import (
	"log"
	"net/http"

	"github.com/LeonLow97/inventory-management-system-golang-react-postgresql/utils"
)

func (app application) Logout(w http.ResponseWriter, req *http.Request) error {
	RetrieveIssuer(w, req)

	cookie, err := req.Cookie("leon-jwt-token")
	if err != nil {
		log.Println("req.Cookie:", err)
		return utils.ApiError{Err: "Internal Server Error", Status: http.StatusInternalServerError}
	}
	cookie.Value = ""

	cookie = &http.Cookie{
		Name:     "leon-jwt-token",
		Value:    "",
		MaxAge:   -1,
		Path:     "",
		Domain:   "",
		Secure:   false,
		HttpOnly: true,
	}

	http.SetCookie(w, cookie)

	return utils.WriteJSON(w, http.StatusOK, "Successfully Logged Out!")
}
