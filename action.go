package usermngmt

import (
	"github.com/Selly-Modules/usermngmt/model"
	"github.com/Selly-Modules/usermngmt/role"
	"github.com/Selly-Modules/usermngmt/user"
)

//
// User
//

// user methods

// CreateUser ...
func (s Service) CreateUser(payload model.UserCreateOptions) error {
	return user.Create(payload)
}

// UpdateUser ...
func (s Service) UpdateUser(userID string, payload model.UserUpdateOptions) error {
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

// GetAllUser ...
func (s Service) GetAllUser(query model.UserAllQuery) model.UserAll {
	return user.All(query)
}

// ChangeAllUsersStatus ...
func (s Service) ChangeAllUsersStatus(roleID, status string) error {
	return user.ChangeAllUsersStatus(roleID, status)
}

//
// Role
//

// role methods

// CreateRole ...
func (s Service) CreateRole(payload model.RoleCreateOptions) error {
	return role.Create(payload)
}

// UpdateRole ...
func (s Service) UpdateRole(roleID string, payload model.RoleUpdateOptions) error {
	return role.Update(roleID, payload)
}

// GetAllRoles ...
func (s Service) GetAllRoles(query model.RoleAllQuery) model.RoleAll {
	return role.All(query)
}
