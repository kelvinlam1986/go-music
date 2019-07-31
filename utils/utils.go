package utils

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(s *string) error {
	if s == nil {
		return errors.New("Reference string provided hash password is empty")
	}

	// Convert password string to byte slice so that we can use it with the bcrypt package
	sBytes :=  []byte(*s)

	// Obtain hashed password via the GenerateFromPassword() method
	hashedBytes, err := bcrypt.GenerateFromPassword(sBytes, bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// Update password string with the hashed version
	*s = string(hashedBytes[:])
	return nil
}