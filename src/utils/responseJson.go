package utils

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func ResponseJson(w http.ResponseWriter, Code int, Message string) {
	jsonStatus := struct {
		Code int `json:"code"`
		Message string `json:"message"`
	}{
		Message: Message,
		Code: Code,
	}

	bs, err := json.Marshal(jsonStatus);
	if err != nil {
		log.Fatalln("Error in ResponseJson utils folder: ", err)
	}

	// privateKey := keys.LoadPrivateKey();
	// hashed := sha256.Sum256(bs)
	// signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hashed[:])
	// CheckError(err)
	// io.WriteString(w, string(signature));

	io.WriteString(w, string(bs));
}

func InternalServerError(w http.ResponseWriter, message string, err error) {
	log.Fatalln(message, err)
	ResponseJson(w, http.StatusInternalServerError, "Internal Server Error");
}