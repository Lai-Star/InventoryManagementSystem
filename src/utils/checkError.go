package utils

import (
	"log"
)

func CheckError(err error) {
	if err != nil {
		log.Println("Internal Server Error: ", err);
	}
}

func CheckErrorDatabase(err error) {
	if err != nil {
		log.Println("PostgreSQL Internal Error: ", err)
	}
}