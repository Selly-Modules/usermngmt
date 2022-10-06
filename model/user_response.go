package model

import (
	"github.com/Selly-Modules/usermngmt/internal"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// User ...
type User struct {
	ID                      string      `json:"_id"`
	Name                    string      `json:"name"`
	Phone                   string      `json:"phone"`
	Email                   string      `json:"email"`
	Status                  string      `json:"status"`
	Role                    RoleShort   `json:"role"`
	RequireToChangePassword bool        `json:"requireToChangePassword"`
	Other                   interface{} `json:"other"`
	Avatar                  interface{} `json:"avatar"`
	CreatedAt               time.Time   `json:"createdAt"`
	UpdatedAt               time.Time   `json:"updatedAt"`
}

// UserOtherBson ...
type UserOtherBson struct {
	Supplier    primitive.ObjectID   `bson:"supplier"`
	Inventories []primitive.ObjectID `bson:"inventories"`
	IsPresident bool                 `bson:"isPresident"`
}

type UserOther struct {
	Supplier    string   `json:"supplier"`
	Inventories []string `json:"inventories"`
	IsPresident bool     `json:"isPresident"`
}

func (m User) GetUserOther() UserOther {
	var (
		userOtherBson UserOtherBson
	)
	bsonBytes, _ := bson.Marshal(m.Other)
	bson.Unmarshal(bsonBytes, &userOtherBson)
	return UserOther{
		Supplier:    userOtherBson.Supplier.Hex(),
		Inventories: internal.ConvertObjectIDsToStrings(userOtherBson.Inventories),
		IsPresident: userOtherBson.IsPresident,
	}
}

type (
	// UserAll ...
	UserAll struct {
		List  []User `json:"list"`
		Total int64  `json:"total"`
		Limit int64  `json:"limit"`
	}
)
