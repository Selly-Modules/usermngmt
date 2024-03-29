package internal

import (
	"fmt"
	"github.com/thoas/go-funk"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strings"

	"github.com/Selly-Modules/mongodb"
	"golang.org/x/crypto/bcrypt"
)

// HashPassword ...
func HashPassword(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), passwordHashingCost)
	return string(bytes)
}

// CheckPasswordHash ...
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// GetSearchString ...
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

// GenerateCode ...
func GenerateCode(s string) string {
	var (
		underscore = "_"
		emptySpace = " "
	)
	return strings.ReplaceAll(mongodb.NonAccentVietnamese(s), emptySpace, underscore)
}

// ConvertObjectIDsToStrings ...
func ConvertObjectIDsToStrings(ids []primitive.ObjectID) []string {
	return funk.Map(ids, func(item primitive.ObjectID) string {
		return item.Hex()
	}).([]string)
}
