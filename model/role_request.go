package model

import (
	"errors"

	"github.com/Selly-Modules/logger"
	"github.com/Selly-Modules/usermngmt/internal"
)

// RoleCreateOptions ...
type RoleCreateOptions struct {
	Name    string
	Level   int
	IsAdmin bool
}

// RoleUpdateOptions ...
type RoleUpdateOptions struct {
	Name    string
	Level   int
	IsAdmin bool
}

// RoleAllQuery ...
type RoleAllQuery struct {
	Page  int64
	Limit int64
	Sort  interface{}
}

// Validate ...
func (co RoleCreateOptions) Validate() error {
	// Name
	if co.Name == "" {
		logger.Error("usermngmt - Role - Create: no name data", logger.LogData{
			Source:  "usermngmt.Validate",
			Message: "usermngmt - Role - Create: no name data",
			Data:    co,
		})
		return errors.New(internal.ErrorInvalidName)
	}

	return nil
}

// Validate ...
func (co RoleUpdateOptions) Validate() error {
	// Name
	if co.Name == "" {
		logger.Error("usermngmt - Role - Update: no name data", logger.LogData{
			Source:  "usermngmt.Validate",
			Message: "usermngmt - Role - Update: no name data",
			Data:    co,
		})
		return errors.New(internal.ErrorInvalidName)
	}

	return nil
}
