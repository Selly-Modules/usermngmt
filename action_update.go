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

// ChangePasswordByUserID ...
func (s Service) ChangePasswordByUserID(userID, oldPassword, newPassword string) error {
	var (
		ctx   = context.Background()
		id, _ = mongodb.NewIDFromString(userID)
	)

	if oldPassword == "" || newPassword == "" {
		return errors.New("new password and old password cannot be empty")
	}

	// Find user
	user, _ := s.userFindByID(ctx, id)
	if user.ID.IsZero() {
		return errors.New("user does not exist")
	}

	// Check old password
	isValid := checkPasswordHash(oldPassword, user.HashedPassword)
	if !isValid {
		return errors.New("the password is incorrect")
	}

	// Update password
	err := s.userUpdateOneByCondition(ctx, bson.M{"_id": user.ID}, bson.M{
		"$set": bson.M{
			"hashedPassword": hashPassword(newPassword),
			"updatedAt":      now(),
		},
	})
	if err != nil {
		return err
	}

	return nil
}

// ChangeStatusByUserID ...
func (s Service) ChangeStatusByUserID(userID, newStatus string) error {
	var (
		ctx   = context.Background()
		id, _ = mongodb.NewIDFromString(userID)
	)

	// Update password
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
