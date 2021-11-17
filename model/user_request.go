package model

import (
	"errors"

	"github.com/Selly-Modules/logger"
)

// UserCreateOptions ...
type UserCreateOptions struct {
	Name                    string
	Phone                   string
	Email                   string
	Password                string
	Status                  string
	RoleID                  string
	RequireToChangePassword bool
	Other                   interface{}
}

// UserUpdateOptions ...
type UserUpdateOptions struct {
	Name   string
	Phone  string
	Email  string
	RoleID string
	Other  interface{}
}

// ChangePasswordOptions ...
type ChangePasswordOptions struct {
	OldPassword string
	NewPassword string
}

// UserAllQuery ...
type UserAllQuery struct {
	Page    int64
	Limit   int64
	Keyword string
	RoleID  string
	Status  string
	Sort    interface{}
}

// Validate ...
func (co UserCreateOptions) Validate() error {
	// Name
	if co.Name == "" {
		logger.Error("usermngmt - User - Create: no name data", logger.LogData{
			"payload": co,
		})
		return errors.New("no name data")
	}

	// Phone
	if co.Phone == "" {
		logger.Error("usermngmt - User - Create: no phone data", logger.LogData{
			"payload": co,
		})
		return errors.New("no phone data")
	}

	// Email
	if co.Email == "" {
		logger.Error("usermngmt - User - Create: no email data", logger.LogData{
			"payload": co,
		})
		return errors.New("no email data")
	}

	// Password
	if co.Password == "" {
		logger.Error("usermngmt - User - Create: no password data", logger.LogData{
			"payload": co,
		})
		return errors.New("no password data")
	}

	// Status
	if co.Status == "" {
		logger.Error("usermngmt - User - Create: no status data", logger.LogData{
			"payload": co,
		})
		return errors.New("no status data")
	}

	// RoleID
	if co.RoleID == "" {
		logger.Error("usermngmt - User - Create: no roleID data", logger.LogData{
			"payload": co,
		})
		return errors.New("no role id data")
	}

	return nil
}

// Validate ...
func (uo UserUpdateOptions) Validate() error {
	// Name
	if uo.Name == "" {
		logger.Error("usermngmt - User - Update: no name data", logger.LogData{
			"payload": uo,
		})
		return errors.New("no name data")
	}

	// Phone
	if uo.Phone == "" {
		logger.Error("usermngmt - User - Update: no phone data", logger.LogData{
			"payload": uo,
		})
		return errors.New("no phone data")
	}

	// Email
	if uo.Email == "" {
		logger.Error("usermngmt - User - Update: no email data", logger.LogData{
			"payload": uo,
		})
		return errors.New("no email data")
	}

	// RoleID
	if uo.RoleID == "" {
		logger.Error("usermngmt - User - Update: no roleID data", logger.LogData{
			"payload": uo,
		})
		return errors.New("no role id data")
	}

	return nil
}

// Validate ...
func (co ChangePasswordOptions) Validate() error {
	// OldPassword, NewPassword
	if co.OldPassword == "" || co.NewPassword == "" {
		logger.Error("usermngmt - User - ChangePassword: old or new password cannot be empty", logger.LogData{
			"payload": co,
		})
		return errors.New("old or new password cannot be empty")
	}

	return nil
}
