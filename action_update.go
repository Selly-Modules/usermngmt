package usermngmt

import (
	"context"

	"github.com/Selly-Modules/mongodb"
	"go.mongodb.org/mongo-driver/bson"
)

// UpdateOptions ...
type UpdateOptions struct {
	Name   string
	Phone  string
	Email  string
	RoleID string
	Other  string
}

// UpdateByUserID ...
func (s Service) UpdateByUserID(userID string, payload UpdateOptions) error {
	var (
		ctx = context.Background()
	)

	// Validate payload
	err := payload.validate(ctx)
	if err != nil {
		return err
	}

	// Setup condition
	id, _ := mongodb.NewIDFromString(userID)
	cond := bson.M{
		"_id": id,
	}

	// Setup update data
	roleID, _ := mongodb.NewIDFromString(payload.RoleID)
	updateData := bson.M{
		"$set": bson.M{
			"name":      payload.Name,
			"phone":     payload.Phone,
			"email":     payload.Email,
			"roleId":    roleID,
			"other":     payload.Other,
			"updatedAt": now(),
		},
	}

	// Update
	err = s.userUpdateOneByCondition(ctx, cond, updateData)
	if err != nil {
		return err
	}

	return nil
}
