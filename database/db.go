package database

import (
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
)

// Table
var (
	tableUser       = "users"
	tableRole       = "roles"
	tablePermission = "permissions"
)

var (
	db     *mongo.Database
	prefix string
)

// Set ...
func Set(instance *mongo.Database, tablePrefix string) {
	db = instance
	prefix = tablePrefix
}

// GetUserCol ...
func GetUserCol() *mongo.Collection {
	return db.Collection(fmt.Sprintf("%s-%s", prefix, tableUser))
}

// GetRoleCol ...
func GetRoleCol() *mongo.Collection {
	return db.Collection(fmt.Sprintf("%s-%s", prefix, tableRole))
}

// GetPermissionCol ...
func GetPermissionCol() *mongo.Collection {
	return db.Collection(fmt.Sprintf("%s-%s", prefix, tablePermission))
}
