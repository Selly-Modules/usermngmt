package usermngmt

import (
	"github.com/Selly-Modules/usermngmt/model"
	"github.com/Selly-Modules/usermngmt/role"
	"github.com/Selly-Modules/usermngmt/user"
)

// Create ...
func (s Service) Create(payload model.UserCreateOptions) error {
	return user.Create(payload)
}

// Update ...
func (s Service) Update(userID string, payload model.UserUpdateOptions) error {
	return user.UpdateByUserID(userID, payload)
}

// ChangeUserPassword ...
func (s Service) ChangeUserPassword(userID string, payload model.ChangePasswordOptions) error {
	return user.ChangeUserPassword(userID, payload)
}

// ChangeUserStatus ...
func (s Service) ChangeUserStatus(userID, newStatus string) error {
	return user.ChangeUserStatus(userID, newStatus)
}

// All ...
func (s Service) All(query model.UserAllQuery) model.UserAll {
	return user.All(query)
}

// RoleCreate ...
func (s Service) RoleCreate(payload model.RoleCreateOptions) error {
	return role.Create(payload)
}
