package internal

import (
	"fmt"

	"github.com/Selly-Modules/mongodb"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), passwordHashingCost)
	return string(bytes)
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GetSearchString(fieldList ...string) string {
	var (
		searchList = make([]interface{}, 0)
		format     = ""
	)

	for i, value := range fieldList {
		searchList = append(searchList, mongodb.NonAccentVietnamese(value))
		if i == 0 {
			format += "%s"
			continue
		}
		format += " %s"
	}
	return fmt.Sprintf(format, searchList...)
}
