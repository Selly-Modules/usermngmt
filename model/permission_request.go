package model

import (
	"errors"

	"github.com/Selly-Modules/logger"
	"github.com/Selly-Modules/mongodb"
)

// PermissionCreateOptions ...
type PermissionCreateOptions struct {
	Name   string
	Code   string
	RoleID string
	Desc   string
}

// PermissionUpdateOptions ...
type PermissionUpdateOptions struct {
	Name   string
	Code   string
	RoleID string
	Desc   string
}

// PermissionAllQuery ...
type PermissionAllQuery struct {
	Page   int64
	Limit  int64
	Sort   interface{}
	RoleID string
}

// Validate ...
func (co PermissionCreateOptions) Validate() error {
	// Name
	if co.Name == "" {
		logger.Error("usermngmt - Permission - Create: no name data", logger.LogData{
			"payload": co,
		})
		return errors.New("no name data")
	}

	// Code
	if co.Code == "" {
		logger.Error("usermngmt - Permission - Create: no code data", logger.LogData{
			"payload": co,
		})
		return errors.New("no code data")
	}

	// RoleID
	if co.RoleID == "" {
		logger.Error("usermngmt - Permission - Create: no roleID data", logger.LogData{
			"payload": co,
		})
		return errors.New("no role id data")
	}
	if _, isValid := mongodb.NewIDFromString(co.RoleID); !isValid {
		return errors.New("invalid role id data")
	}

	// Desc
	if co.Desc == "" {
		logger.Error("usermngmt - Permission - Create: no desc data", logger.LogData{
			"payload": co,
		})
		return errors.New("no desc data")
	}
	return nil
}

// Validate ...
func (co PermissionUpdateOptions) Validate() error {
	// Name
	if co.Name == "" {
		logger.Error("usermngmt - Permission - Update: no name data", logger.LogData{
			"payload": co,
		})
		return errors.New("no name data")
	}

	// Code
	if co.Code == "" {
		logger.Error("usermngmt - Permission - Update: no code data", logger.LogData{
			"payload": co,
		})
		return errors.New("no code data")
	}

	// RoleID
	if co.RoleID == "" {
		logger.Error("usermngmt - Permission - Update: no roleID data", logger.LogData{
			"payload": co,
		})
		return errors.New("no role id data")
	}
	if _, isValid := mongodb.NewIDFromString(co.RoleID); !isValid {
		return errors.New("invalid role id data")
	}

	// Desc
	if co.Desc == "" {
		logger.Error("usermngmt - Permission - Update: no desc data", logger.LogData{
			"payload": co,
		})
		return errors.New("no desc data")
	}
	return nil
}
