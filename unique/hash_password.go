package unique

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("خطا در تولید هش %w", err) 
	}
	return string(hashPassword), nil
}
func CheckPassword(password string, hashpassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashpassword), []byte(password))
}