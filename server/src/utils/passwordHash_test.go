package utils

import (
	"testing"
)

func Test_CompareHash(t *testing.T) {
	compareHashTests := []struct {
		testName       string
		hashedPassword string
		password       string
		expected       bool
	}{
		{"Correct Password Hash", "$2a$10$L4K07kb8rru64Q9f9UwXiO443LmlDkGw83N2KmNs4UscNEOGLBZxm", "Correct0!@", true},
		{"Wrong Password Hash", "$2a$10$8WUviqVs02eSxmbn3JPR3.Zbs9zLb4NYpWPLKeTbiqiaGWTxWmBhx", "WrongPW!@", false},
		{"Empty Password", "DidNotHash", "DidNotHash", false},
		{"Generate Hash", GenerateHash("HelloWorld1!"), "HelloWorld1!", true},
		{"Incorrect generated hash", GenerateHash("TestPass12@"), "TestPass13!", false},
		{"Short Password", GenerateHash("Abc123"), "Abc123", true},
		{"Long Password", GenerateHash("P@ssw0rdLongEnough123456"), "P@ssw0rdLongEnough123456", true},
		{"Weak Password", GenerateHash("password"), "password", false},
		{"Non-ASCII Password", GenerateHash("Привет123!"), "Привет123!", true},
		{"Password with Special Characters", GenerateHash("TestPass12!@#$%^&*()_+"), "TestPass12!@#$%^&*()_+", true},
	}

	for _, e := range compareHashTests {
		result := CompareHash(e.hashedPassword, e.password)
		if e.expected && !result {
			t.Errorf("%s: expected %v but got %v", e.testName, e.expected, result)
		}
	}
}