package helper

import "golang.org/x/crypto/bcrypt"

func EncryptValue(value string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(value), bcrypt.DefaultCost)
	return string(bytes), err
}

func CompareEncryptedValue(value, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(value))
	return err == nil
}
