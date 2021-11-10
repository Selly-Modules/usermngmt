package database

import (
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
)

// Table
var (
	tableUser = "users"
	tableRole = "roles"
)

var (
	db     *mongo.Database
	prefix string
)

func Set(instance *mongo.Database, tablePrefix string) {
	db = instance
	prefix = tablePrefix
}

func GetUserCol() *mongo.Collection {
	return db.Collection(fmt.Sprintf("%s-%s", prefix, tableUser))
}

func GetRoleCol() *mongo.Collection {
	return db.Collection(fmt.Sprintf("%s-%s", prefix, tableRole))
}
