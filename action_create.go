package usermngmt

import (
	"context"
	"errors"
	"fmt"

	"github.com/Selly-Modules/logger"
	"github.com/Selly-Modules/mongodb"
)

// CreateOptions ...
type CreateOptions struct {
	Name         string
	Phone        string
	Email        string
	HashPassword string
	Status       string
	RoleID       string
	Other        string
}

// Create ...
func (s Service) Create(payload CreateOptions) error {
	var (
		col = s.getUserCollection()
		ctx = context.Background()
	)

	// Validate payload
	err := payload.validate()
	if err != nil {
		return err
	}

	// New user data from payload
	userData, err := payload.newUser()
	if err != nil {
		return err
	}

	// Find phone,email exists or not
	if s.haveNameOrPhoneExisted(ctx, userData.Phone, userData.Email) {
		return errors.New("have name or phone existed")
	}

	// Create device
	_, err = col.InsertOne(ctx, userData)
	if err != nil {
		logger.Error("usermngmt - Create ", logger.LogData{
			"doc": userData,
			"err": err.Error(),
		})
		return fmt.Errorf("error when create user: %s", err.Error())
	}

	return nil
}

func (payload CreateOptions) newUser() (result User, err error) {
	timeNow := now()

	// New RoleID from string
	roleID, isValid := mongodb.NewIDFromString(payload.RoleID)
	if !isValid {
		err = errors.New("invalid roleID")
		return
	}

	return User{
		ID:           mongodb.NewObjectID(),
		Name:         payload.Name,
		Phone:        payload.Phone,
		Email:        payload.Email,
		HashPassword: payload.HashPassword,
		Status:       payload.Status,
		RoleID:       roleID,
		Other:        payload.Other,
		CreatedAt:    timeNow,
		UpdatedAt:    timeNow,
	}, nil
}
