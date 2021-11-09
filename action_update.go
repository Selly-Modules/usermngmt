package usermngmt

import (
	"context"
	"errors"

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

// ChangePasswordOptions ...
type ChangePasswordOptions struct {
	OldPassword string
	NewPassword string
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

// ChangeUserPassword ...
func (s Service) ChangeUserPassword(userID string, opt ChangePasswordOptions) error {
	var (
		ctx = context.Background()
	)

	// Validate payload
	err := opt.validate(userID)
	if err != nil {
		return err
	}

	// Find user
	id, _ := mongodb.NewIDFromString(userID)
	user, _ := s.userFindByID(ctx, id)
	if user.ID.IsZero() {
		return errors.New("user not found")
	}

	// Check old password
	if isValid := checkPasswordHash(opt.OldPassword, user.HashedPassword); !isValid {
		return errors.New("the password is incorrect")
	}

	// Update password
	err = s.userUpdateOneByCondition(ctx, bson.M{"_id": user.ID}, bson.M{
		"$set": bson.M{
			"hashedPassword": hashPassword(opt.NewPassword),
			"updatedAt":      now(),
		},
	})
	if err != nil {
		return err
	}

	return nil
}

// ChangeUserStatus ...
func (s Service) ChangeUserStatus(userID, newStatus string) error {
	var (
		ctx = context.Background()
	)

	// Validate userID
	id, isValid := mongodb.NewIDFromString(userID)
	if !isValid {
		return errors.New("invalid user id data")
	}

	// Update status
	err := s.userUpdateOneByCondition(ctx, bson.M{"_id": id}, bson.M{
		"$set": bson.M{
			"status":    newStatus,
			"updatedAt": now(),
		},
	})
	if err != nil {
		return err
	}

	return nil
}
