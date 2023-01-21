package utils

import "regexp"

func CheckLengthRange(str string, minLength int, maxLength int) bool {
	if (len(str) > maxLength || len(str) < minLength) {
		return false
	}
	return true
}

func CheckMaxLength(str string, maxLength int) bool {
	return len(str) <= maxLength
}

func CheckMinLength(str string, minLength int) bool {
	return len(str) >= minLength
}

func CheckWhiteSpaces(str string) bool {
	whiteSpaceRegex := regexp.MustCompile(`^\S*$`)

	return whiteSpaceRegex.MatchString(str)
}

func CheckSpecialChar(str string) bool {
	specialCharRegex := regexp.MustCompile(`^[a-zA-Z0-9_]*$`)

	return specialCharRegex.MatchString(str)
}

func CheckEmailAddress (email string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,4}$`)

	return emailRegex.MatchString(email)
}





