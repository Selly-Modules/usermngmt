package internal

import (
	"errors"
	"time"

	"github.com/Selly-Modules/logger"
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

// AllQuery ...
type AllQuery struct {
	Page    int64
	Limit   int64
	Keyword string
	RoleID  string
	Status  string
}

// User ...
type User struct {
	ID        string    `json:"_id"`
	Name      string    `json:"name"`
	Phone     string    `json:"phone"`
	Email     string    `json:"email"`
	Status    string    `json:"status"`
	Role      RoleShort `json:"role"`
	Other     string    `json:"other"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type (
	// UserAll ...
	UserAll struct {
		List  []User `json:"list"`
		Total int64  `json:"total"`
	}
)

// NewUser ...
func (payload CreateOptions) NewUser() (result DBUser, err error) {
	timeNow := Now()
	roleID, _ := mongodb.NewIDFromString(payload.RoleID)
	return DBUser{
		ID:             mongodb.NewObjectID(),
		Name:           payload.Name,
		SearchString:   GetSearchString(payload.Name, payload.Phone, payload.Email),
		Phone:          payload.Phone,
		Email:          payload.Email,
		HashedPassword: HashPassword(payload.Password),
		Status:         payload.Status,
		RoleID:         roleID,
		Other:          payload.Other,
		CreatedAt:      timeNow,
		UpdatedAt:      timeNow,
	}, nil
}

// Validate ...
func (co CreateOptions) Validate() error {
	// Name
	if co.Name == "" {
		logger.Error("usermngmt - Create: no Name data", logger.LogData{
			"payload": co,
		})
		return errors.New("no name data")
	}

	// Phone
	if co.Phone == "" {
		logger.Error("usermngmt - Create: no phone data", logger.LogData{
			"payload": co,
		})
		return errors.New("no phone data")
	}

	// Email
	if co.Email == "" {
		logger.Error("usermngmt - Create: no email data", logger.LogData{
			"payload": co,
		})
		return errors.New("no email data")
	}

	// Password
	if co.Password == "" {
		logger.Error("usermngmt - Create: no password data", logger.LogData{
			"payload": co,
		})
		return errors.New("no password data")
	}

	// Status
	if co.Status == "" {
		logger.Error("usermngmt - Create: no status data", logger.LogData{
			"payload": co,
		})
		return errors.New("no status data")
	}

	// RoleID
	if co.RoleID == "" {
		logger.Error("usermngmt - Create: no roleID data", logger.LogData{
			"payload": co,
		})
		return errors.New("no role id data")
	}

	return nil
}

// Validate ...
func (uo UpdateOptions) Validate() error {
	// Name
	if uo.Name == "" {
		logger.Error("usermngmt - Update: no name data", logger.LogData{
			"payload": uo,
		})
		return errors.New("no name data")
	}

	// Phone
	if uo.Phone == "" {
		logger.Error("usermngmt - Update: no phone data", logger.LogData{
			"payload": uo,
		})
		return errors.New("no phone data")
	}

	// Email
	if uo.Email == "" {
		logger.Error("usermngmt - Update: no email data", logger.LogData{
			"payload": uo,
		})
		return errors.New("no email data")
	}

	// RoleID
	if uo.RoleID == "" {
		logger.Error("usermngmt - Update: no roleID data", logger.LogData{
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
		logger.Error("usermngmt - ChangePassword: old or new password cannot be empty", logger.LogData{
			"payload": co,
		})
		return errors.New("old or new password cannot be empty")
	}

	return nil
}
