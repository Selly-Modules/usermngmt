package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// DBRole ...
type DBRole struct {
	ID        primitive.ObjectID `bson:"_id"`
	Name      string             `bson:"name"`
	Code      string             `bson:"code"`
	IsAdmin   bool               `bson:"isAdmin"`
	CreatedAt time.Time          `bson:"createdAt"`
	UpdatedAt time.Time          `bson:"updatedAt"`
}

// DBUser ...
type DBUser struct {
	ID             primitive.ObjectID `bson:"_id"`
	Name           string             `bson:"name"`
	SearchString   string             `bson:"searchString"`
	Phone          string             `bson:"phone"` // unique
	Email          string             `bson:"email"` // unique
	HashedPassword string             `bson:"hashedPassword"`
	Status         string             `bson:"status"`
	RoleID         primitive.ObjectID `bson:"roleId"`
	Other          string             `bson:"other"`
	CreatedAt      time.Time          `bson:"createdAt"`
	UpdatedAt      time.Time          `bson:"updatedAt"`
}
