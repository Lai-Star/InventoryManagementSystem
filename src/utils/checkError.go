package utils

import (
	"log"
)

func CheckError(err error) {
	if err != nil {
		log.Fatalln("Internal Server Error: ", err);
	}
}

func CheckErrorDatabase(err error) {
	if err != nil {
		log.Fatalln("PostgreSQL Internal Error: ", err)
	}
}