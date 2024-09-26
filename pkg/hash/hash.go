package hash

import (
	"github.com/labstack/gommon/log"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string, salt string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(salt+password), 14)
	return string(bytes), err
}

func VerifyPassword(password, hash, salt string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(salt+password))
	log.Error(err)
	return err == nil
}

func GenerateSalt() string {
	return "salt"
}
