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

// ResponseUser ...
type ResponseUser struct {
	ID        string    `json:"_id"`
	Name      string    `json:"name"`
	Phone     string    `json:"phone"`
	Email     string    `json:"email"`
	Status    string    `json:"status"`
	Role      RoleShort `json:"role"`
	Other     string    `json:"other"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
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

type RoleShort struct {
	ID      string `json:"_id"`
	Name    string `json:"name"`
	IsAdmin bool   `json:"isAdmin"`
}

type (
	// ResponseUserAll ...
	ResponseUserAll struct {
		List  []ResponseUser `json:"list"`
		Total int64          `json:"total"`
	}
)
