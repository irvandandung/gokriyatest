package helpers

import(
	"log"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 12)
    return string(bytes), err
}

func CheckPasswordHash(password string, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(password), []byte(hash))
    log.Println(err)
    return err == nil
}