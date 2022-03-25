package model

import (
	"errors"

	"github.com/Selly-Modules/logger"
	"github.com/Selly-Modules/usermngmt/internal"
	"go.mongodb.org/mongo-driver/bson"
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
	Avatar                  interface{} // if not, pass default file object
}

// UserUpdateOptions ...
type UserUpdateOptions struct {
	Name   string
	Phone  string
	Email  string
	RoleID string
	Other  map[string]interface{}
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
	Other   map[string]interface{} // query fields in other object
	Cond    bson.M
}

// UserByPermissionQuery ...
type UserByPermissionQuery struct {
	Page       int64
	Limit      int64
	Keyword    string
	Status     string
	Permission string // permission code
	Sort       interface{}
	Other      map[string]interface{} // query fields in other object
	Cond       bson.M
}

// UserCountQuery ...
type UserCountQuery struct {
	RoleID string
	Other  map[string]interface{} // query fields in other object
}

// Validate ...
func (co UserCreateOptions) Validate() error {
	// Name
	if co.Name == "" {
		logger.Error("usermngmt - User - Create: no name data", logger.LogData{
			"payload": co,
		})
		return errors.New(internal.ErrorInvalidName)
	}

	// Phone
	if co.Phone == "" {
		logger.Error("usermngmt - User - Create: no phone data", logger.LogData{
			"payload": co,
		})
		return errors.New(internal.ErrorInvalidPhoneNumber)
	}

	// Email
	if co.Email == "" {
		logger.Error("usermngmt - User - Create: no email data", logger.LogData{
			"payload": co,
		})
		return errors.New(internal.ErrorInvalidEmail)
	}

	// Password
	if co.Password == "" {
		logger.Error("usermngmt - User - Create: no password data", logger.LogData{
			"payload": co,
		})
		return errors.New(internal.ErrorInvalidPassword)
	}

	// Status
	if co.Status == "" {
		logger.Error("usermngmt - User - Create: no status data", logger.LogData{
			"payload": co,
		})
		return errors.New(internal.ErrorInvalidStatus)
	}

	// RoleID
	if co.RoleID == "" {
		logger.Error("usermngmt - User - Create: no roleID data", logger.LogData{
			"payload": co,
		})
		return errors.New(internal.ErrorInvalidRole)
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
		return errors.New(internal.ErrorInvalidName)
	}

	// Phone
	if uo.Phone == "" {
		logger.Error("usermngmt - User - Update: no phone data", logger.LogData{
			"payload": uo,
		})
		return errors.New(internal.ErrorInvalidPhoneNumber)
	}

	// Email
	if uo.Email == "" {
		logger.Error("usermngmt - User - Update: no email data", logger.LogData{
			"payload": uo,
		})
		return errors.New(internal.ErrorInvalidEmail)
	}

	// RoleID
	if uo.RoleID == "" {
		logger.Error("usermngmt - User - Update: no roleID data", logger.LogData{
			"payload": uo,
		})
		return errors.New(internal.ErrorInvalidRole)
	}

	return nil
}

// Validate ...
func (co ChangePasswordOptions) Validate() error {
	// OldPassword, NewPassword
	if co.OldPassword == "" {
		logger.Error("usermngmt - User - ChangePassword: no old password data", logger.LogData{
			"payload": co,
		})
		return errors.New(internal.ErrorInvalidOldPassword)
	}

	if co.NewPassword == "" {
		logger.Error("usermngmt - User - ChangePassword: no new password data", logger.LogData{
			"payload": co,
		})
		return errors.New(internal.ErrorInvalidNewPassword)
	}

	return nil
}

// Validate ...
func (q UserByPermissionQuery) Validate() error {
	// OldPassword, NewPassword
	if q.Permission == "" {
		logger.Error("usermngmt - User - GetUsersByPermission : invalid permission", logger.LogData{
			"payload": q,
		})
		return errors.New(internal.ErrorInvalidPermission)
	}

	return nil
}
