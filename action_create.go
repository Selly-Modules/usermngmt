package usermngmt

import (
	"context"
	"fmt"

	"github.com/Selly-Modules/logger"
	"github.com/Selly-Modules/mongodb"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CreateOptions ...
type CreateOptions struct {
	Name     string
	Phone    string
	Email    string
	Password string
	Status   string
	RoleID   primitive.ObjectID
	Other    string
}

// Create ...
func (s Service) Create(payload CreateOptions) error {
	var (
		col = s.getUserCollection()
		ctx = context.Background()
	)

	// Validate payload
	err := payload.validate(ctx)
	if err != nil {
		return err
	}

	// New user data from payload
	userData, err := payload.newUser()
	if err != nil {
		return err
	}

	// Create device
	_, err = col.InsertOne(ctx, userData)
	if err != nil {
		logger.Error("usermngmt - Create", logger.LogData{
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
		HashedPassword: hashPassword(payload.Password),
		Status:         payload.Status,
		RoleID:         payload.RoleID,
		Other:          payload.Other,
		CreatedAt:      timeNow,
		UpdatedAt:      timeNow,
	}, nil
}
