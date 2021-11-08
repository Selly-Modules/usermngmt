package usermngmt

import (
	"errors"

	"github.com/Selly-Modules/logger"
)

func (co CreateOptions) validate() error {
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

	// HashedPassword
	if co.HashedPassword == "" {
		logger.Error("usermngmt - Create: no hashedPassword data", logger.LogData{
			"payload": co,
		})
		return errors.New("no hashedPassword data")
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

	return nil
}
