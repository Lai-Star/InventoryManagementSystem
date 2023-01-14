package utils

import (
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
)

func Generate2FA() (string) {
	// Use Cryptographically Secure Pseudorandom Number Generator (CSPRNG)

	// To generate 5 random alphabets (lower/upper) using CSPRNG
	var otpAlphabets []byte
	for i := 0; i < 5; i++ {
		n, err := rand.Int(rand.Reader, big.NewInt(52))
		if err != nil {
			log.Println("Error in generating random number for otp: ", err)
		}

		if n.Int64() < 26 {
			otpAlphabets = append(otpAlphabets, byte(n.Int64()) + 65) // appends uppercase character
		} else {
			otpAlphabets = append(otpAlphabets, byte(n.Int64()) + 71) // appends lowercase character
		}
	}
	otpAlphabetsString := string(otpAlphabets)

	// To generate 6 random numerical values using CSPRNG
	max := big.NewInt(1000000)
	otpNumbers, err := rand.Int(rand.Reader, max)
	if err != nil {
		log.Println("Error in generating random 6 numbers for otp", err)
	}

	// "%06d" pads the resulting string with leading 0s. E.g., 123 becomes 000123
	otpNumbersString := fmt.Sprintf("%06d", otpNumbers)
	
	return otpAlphabetsString + "-" + otpNumbersString
}