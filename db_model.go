package usermngmt

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// dbUser ...
type dbUser struct {
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

// dbRole ...
type dbRole struct {
	ID        primitive.ObjectID `bson:"_id" json:"_id"`
	Name      string             `bson:"name" json:"name"`
	Code      string             `bson:"code" json:"code"`
	IsAdmin   bool               `bson:"isAdmin" json:"isAdmin"`
	CreatedAt time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt time.Time          `bson:"updatedAt" json:"updatedAt"`
}
