package helper

import (
	"github.com/krifik/groupchat-be/exception"

	"golang.org/x/crypto/bcrypt"
)

func ToHashedPassword(password string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	exception.PanicIfNeeded(err)
	return string(hashedPassword)
}
