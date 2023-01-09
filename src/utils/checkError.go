package utils

import "fmt"

func CheckError(err error) {
	if err != nil {
		fmt.Println("Internal Server Error:", err);
	}
}

func CheckErrorDatabase(err error) {
	if err != nil {
		fmt.Println("PostgreSQL Internal Server Error:", err)
	}
}