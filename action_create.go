package usermngmt

import (
	"context"
	"errors"
	"fmt"

	"github.com/Selly-Modules/logger"
	"github.com/Selly-Modules/mongodb"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CreateOptions ...
type CreateOptions struct {
	Name           string
	Phone          string
	Email          string
	HashedPassword string
	Status         string
	RoleID         primitive.ObjectID
	Other          string
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

	//  Find roleID exists or not
	if !s.isRoleIDAlreadyExisted(ctx, userData.RoleID) {
		return errors.New("roleID does not exist")
	}

	// Find phone number,email exists or not
	if s.isPhoneNumberOrEmailAlreadyExisted(ctx, userData.Phone, userData.Email) {
		return errors.New("phone number or email already existed")
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
	return User{
		ID:             mongodb.NewObjectID(),
		Name:           payload.Name,
		Phone:          payload.Phone,
		Email:          payload.Email,
		HashedPassword: payload.HashedPassword,
		Status:         payload.Status,
		RoleID:         payload.RoleID,
		Other:          payload.Other,
		CreatedAt:      timeNow,
		UpdatedAt:      timeNow,
	}, nil
}
