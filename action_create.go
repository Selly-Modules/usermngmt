package usermngmt

import (
	"context"

	"github.com/Selly-Modules/mongodb"
)

// CreateOptions ...
type CreateOptions struct {
	Name     string
	Phone    string
	Email    string
	Password string
	Status   string
	RoleID   string
	Other    string
}

// Create ...
func (s Service) Create(payload CreateOptions) error {
	var (
		ctx = context.Background()
	)

	// Validate payload
	if err := payload.validate(ctx); err != nil {
		return err
	}

	// New user data from payload
	doc, err := payload.newUser()
	if err != nil {
		return err
	}

	// Create user
	if err = s.userCreate(ctx, doc); err != nil {
		return err
	}

	return nil
}

func (payload CreateOptions) newUser() (result User, err error) {
	timeNow := now()
	roleID, _ := mongodb.NewIDFromString(payload.RoleID)
	return User{
		ID:             mongodb.NewObjectID(),
		Name:           payload.Name,
		Phone:          payload.Phone,
		Email:          payload.Email,
		HashedPassword: hashPassword(payload.Password),
		Status:         payload.Status,
		RoleID:         roleID,
		Other:          payload.Other,
		CreatedAt:      timeNow,
		UpdatedAt:      timeNow,
	}, nil
}
