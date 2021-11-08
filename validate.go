package usermngmt

import (
	"context"
	"errors"

	"github.com/Selly-Modules/logger"
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
	if co.RoleID.IsZero() {
		logger.Error("usermngmt - Create: invalid roleID data", logger.LogData{
			"payload": co,
		})
		return errors.New("invalid roleID data")
	}

	//  Find roleID exists or not
	if !s.isRoleIDExisted(ctx, co.RoleID) {
		return errors.New("role id does not exist")
	}

	// Find phone number,email exists or not
	if s.isPhoneNumberOrEmailExisted(ctx, co.Phone, co.Email) {
		return errors.New("phone number or email already existed")
	}

	return nil
}
