package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

type apiFunc func(http.ResponseWriter, *http.Request) error

func WriteJSON(w http.ResponseWriter, status int, v interface{}) error {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	// privateKey := keys.LoadPrivateKey();
	// hashed := sha256.Sum256(bs)
	// signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hashed[:])
	// CheckError(err)
	// io.WriteString(w, string(signature));

	return json.NewEncoder(w).Encode(v)
}

func MakeHTTPHandler(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			if e, ok := err.(ApiError); ok {
				WriteJSON(w, e.Status, e)
				return
			}
			log.Println("Error in MakeHTTPHandler:", err)
			WriteJSON(w, http.StatusInternalServerError, ApiError{Err: "Internal Serval Error", Status: http.StatusInternalServerError})
		}
	}
}
