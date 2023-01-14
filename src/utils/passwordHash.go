package utils

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func GenerateHash(password string) (string) {
	// Generate the hash using a default cost of 10
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Error in generating password hash: ", err)
	}

	return string(hash);
}

func CompareHash(hashedPassword string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil;
}