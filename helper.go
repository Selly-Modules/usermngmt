package usermngmt

import "golang.org/x/crypto/bcrypt"

func hashPassword(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), passwordHashingCost)
	return string(bytes)
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
