package handlers_user

import (
	"net/http"
	"os"

	"github.com/LeonLow97/inventory-management-system-golang-react-postgresql/utils"
	"github.com/golang-jwt/jwt/v4"
)

func RetrieveIssuer(w http.ResponseWriter, req *http.Request) bool {
	cookie, err := req.Cookie("leon-jwt-token")
	if err != nil {
		utils.InternalServerError(w, "Internal Server Error in retrieving cookie: ", err);
		return false;
	}

	// Parses the cookie jwt claims using the secret key to verify
	token, err := jwt.ParseWithClaims(cookie.Value, &jwt.RegisteredClaims{}, func(*jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	if err != nil {
		utils.InternalServerError(w, "Internal Server Error in parsing cookie: ", err)
		return false
	}

	// To access the issuer, we need the token claims to be type RegisteredClaims
	claims := token.Claims.(*jwt.RegisteredClaims)

	// fmt.Println("Retrieved Issuer using claims.Issuer: ", claims.Issuer)
	w.Header().Set("username", claims.Issuer)
	// fmt.Println("Retrieved Issuer using w.Header().Get():" , w.Header().Get("username"))

	return true
}