package util

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func CheckPassword(providedPassword, existingHash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(existingHash), []byte(providedPassword))
	if err != nil {
		return err
	}
	return nil
}
