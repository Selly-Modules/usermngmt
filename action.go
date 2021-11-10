package usermngmt

import (
	"github.com/Selly-Modules/usermngmt/model"
)

// Create ...
func (s Service) Create(payload model.UserCreateOptions) error {
	return s.handler.User.Create(payload)
}

// Update ...
func (s Service) Update(userID string, payload model.UserUpdateOptions) error {
	return s.handler.User.UpdateByUserID(userID, payload)
}

// ChangeUserPassword ...
func (s Service) ChangeUserPassword(userID string, payload model.ChangePasswordOptions) error {
	return s.handler.User.ChangeUserPassword(userID, payload)
}

func (s Service) ChangeUserStatus(userID, newStatus string) error {
	return s.handler.User.ChangeUserStatus(userID, newStatus)
}

func (s Service) All(query model.UserAllQuery) model.UserAll {
	return s.handler.User.All(query)
}

func (s Service) RoleCreate(payload model.RoleCreateOptions) error {
	return s.handler.Role.Create(payload)
}
