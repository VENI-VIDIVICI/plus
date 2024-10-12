package hash

import (
	"github.com/VENI-VIDIVICI/plus/pkg/logger"
	"golang.org/x/crypto/bcrypt"
)

func BcryptHash(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	logger.LogIf(err)
	return string(bytes)
}

func BcryptCheck(hash string, answer string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(answer))
	return err != nil
}

func BcryptIsHashed(str string) bool {
	return len(str) == 60
}
