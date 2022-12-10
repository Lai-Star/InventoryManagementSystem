package utils

import "golang.org/x/crypto/bcrypt"

func CompareHash(hashedPassword string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil;
}

func GenerateHash(password string) (string) {
	// Generate the hash using a default cost of 10
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	CheckError(err)

	return string(hash);
}