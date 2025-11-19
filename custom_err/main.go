package main

import (
	"fmt"
	"strings"
)

// custom error type

type ValidationErr struct {
	Field   string
	Message string
}

func (v ValidationErr) Error() string {
	return fmt.Sprintf("%s: %s", v.Field, v.Message)
}
func validateEmail(email string) error {
	if !strings.Contains(email, "@") {
		return ValidationErr{Field: "email", Message: "format email tidak valid"}
	}
	return nil
}

func validatePassword(password string) error {
	if len(password) < 8 {
		return ValidationErr{Field: "password", Message: "password minimal 8 karakter"}
	}
	return nil
}

func main() {
	// contoh error handling
	if err := validateEmail("didan.com"); err != nil {
		fmt.Println(err)
	}
	if err := validatePassword("123456"); err != nil {
		fmt.Println(err)
	}

}
