package usermngmt

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User ...
type User struct {
	ID             primitive.ObjectID `bson:"_id" json:"_id"`
	Name           string             `bson:"name" json:"name"`
	Phone          string             `bson:"phone" json:"phone"` // unique
	Email          string             `bson:"email" json:"email"` // unique
	HashedPassword string             `bson:"hashedPassword" json:"-"`
	Status         string             `bson:"status" json:"status"`
	RoleID         primitive.ObjectID `bson:"roleId" json:"roleId"`
	Other          string             `bson:"other" json:"other"`
	CreatedAt      time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt      time.Time          `bson:"updatedAt" json:"updatedAt"`
}

// Role ...
type Role struct {
	ID        primitive.ObjectID `bson:"_id" json:"_id"`
	Name      string             `bson:"name" json:"name"`
	Code      string             `bson:"code" json:"code"`
	IsAdmin   bool               `bson:"isAdmin" json:"isAdmin"`
	CreatedAt time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt time.Time          `bson:"updatedAt" json:"updatedAt"`
}
