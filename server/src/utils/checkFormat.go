package utils

import (
	"regexp"
)

// Return true if the string length is within the range provided (inclusive)
func CheckLengthRange(str string, minLength int, maxLength int) bool {
	if (len(str) > maxLength || len(str) < minLength) {
		return false
	}
	return true
}

// Returns true if the maximum length is met (inclusive)
func CheckMaxLength(str string, maxLength int) bool {
	return len(str) <= maxLength
}

// Returns true if the minimum length is met (inclusive)
func CheckMinLength(str string, minLength int) bool {
	return len(str) >= minLength
}

func CheckWhiteSpaces(str string) bool {
	whiteSpaceRegex := regexp.MustCompile(`^\S*$`)

	return whiteSpaceRegex.MatchString(str)
}

func CheckUsernameSpecialChar(str string) bool {
	specialCharRegex := regexp.MustCompile(`^[a-zA-Z0-9_]*$`)

	return specialCharRegex.MatchString(str)
}

// Password must contain at least one lowercase, uppercase, number and special character
func CheckPasswordSpecialChar(str string) bool {
    // Check for at least one lowercase letter
    if !regexp.MustCompile(`[a-z]`).MatchString(str) {
        return false
    }
    
    // Check for at least one uppercase letter
    if !regexp.MustCompile(`[A-Z]`).MatchString(str) {
        return false
    }
    
    // Check for at least one number
    if !regexp.MustCompile(`\d`).MatchString(str) {
        return false
    }
    
    // Check for at least one special character
    if !regexp.MustCompile(`[!@#$%^&*()_+{}[\]:;'"<>,.?/\\|~-]`).MatchString(str) {
        return false
    }
    
    // Password meets all criteria
    return true
}



func CheckEmailAddress(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,4}$`)

	return emailRegex.MatchString(email)
}





