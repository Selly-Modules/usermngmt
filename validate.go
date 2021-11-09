package usermngmt

import (
	"context"
	"errors"

	"github.com/Selly-Modules/logger"
	"github.com/Selly-Modules/mongodb"
)

func (co CreateOptions) validate(ctx context.Context) error {
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

	//  Find roleID exists or not
	roleID, isValid := mongodb.NewIDFromString(co.RoleID)
	if !isValid {
		return errors.New("invalid role id data")
	}
	if !s.isRoleIDExisted(ctx, roleID) {
		return errors.New("role id does not exist")
	}

	// Find phone number,email exists or not
	if s.isPhoneNumberOrEmailExisted(ctx, co.Phone, co.Email) {
		return errors.New("phone number or email already existed")
	}

	return nil
}

func (co UpdateOptions) validate(ctx context.Context) error {
	// Name
	if co.Name == "" {
		logger.Error("usermngmt - Update: no name data", logger.LogData{
			"payload": co,
		})
		return errors.New("no name data")
	}

	// Phone
	if co.Phone == "" {
		logger.Error("usermngmt - Update: no phone data", logger.LogData{
			"payload": co,
		})
		return errors.New("no phone data")
	}

	// Email
	if co.Email == "" {
		logger.Error("usermngmt - Update: no email data", logger.LogData{
			"payload": co,
		})
		return errors.New("no email data")
	}

	// RoleID
	if co.RoleID == "" {
		logger.Error("usermngmt - Update: no roleID data", logger.LogData{
			"payload": co,
		})
		return errors.New("no role id data")
	}

	//  Find roleID exists or not
	roleID, isValid := mongodb.NewIDFromString(co.RoleID)
	if !isValid {
		return errors.New("invalid role id data")
	}
	if !s.isRoleIDExisted(ctx, roleID) {
		return errors.New("role id does not exist")
	}

	// Find phone number,email exists or not
	if s.isPhoneNumberOrEmailExisted(ctx, co.Phone, co.Email) {
		return errors.New("phone number or email already existed")
	}

	return nil
}

func (co ChangePasswordOptions) validate(userID string) error {
	// OldPassword, NewPassword
	if co.OldPassword == "" || co.NewPassword == "" {
		logger.Error("usermngmt - ChangePassword: old or new password cannot be empty", logger.LogData{
			"payload": co,
		})
		return errors.New("old or new password cannot be empty")
	}

	// UserID
	if _, isValid := mongodb.NewIDFromString(userID); !isValid {
		logger.Error("usermngmt - ChangePassword: invalid userID data", logger.LogData{
			"payload": co,
		})
		return errors.New("invalid user id data")
	}

	return nil
}
