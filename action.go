package usermngmt

import (
	"github.com/Selly-Modules/usermngmt/cache"
	"github.com/Selly-Modules/usermngmt/model"
	"github.com/Selly-Modules/usermngmt/permission"
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

// GetAllUsers ...
func (s Service) GetAllUsers(query model.UserAllQuery) model.UserAll {
	return user.All(query)
}

// ChangeAllUsersStatus ...
func (s Service) ChangeAllUsersStatus(roleID, status string) error {
	return user.ChangeAllUsersStatus(roleID, status)
}

// LoginWithEmailAndPassword ...
func (s Service) LoginWithEmailAndPassword(email, password string) (model.User, error) {
	return user.LoginWithEmailAndPassword(email, password)
}

// HasPermission ...
func (s Service) HasPermission(userID, permission string) bool {
	return user.HasPermission(userID, permission)
}

//
// Role
//

// role methods

// CreateRole ...
func (s Service) CreateRole(payload model.RoleCreateOptions) error {
	if err := role.Create(payload); err != nil {
		return err
	}
	cache.Roles()
	return nil
}

// UpdateRole ...
func (s Service) UpdateRole(roleID string, payload model.RoleUpdateOptions) error {
	return role.Update(roleID, payload)
}

// GetAllRoles ...
func (s Service) GetAllRoles(query model.RoleAllQuery) model.RoleAll {
	return role.All(query)
}

//
// Permission
//

// permission methods

// CreatePermission ...
func (s Service) CreatePermission(payload model.PermissionCreateOptions) error {
	if err := permission.Create(payload); err != nil {
		return err
	}
	cache.Roles()
	return nil
}

// UpdatePermission ...
func (s Service) UpdatePermission(permissionID string, payload model.PermissionUpdateOptions) error {
	if err := permission.Update(permissionID, payload); err != nil {
		return err
	}
	cache.Roles()
	return nil
}

// GetAllPermissions ...
func (s Service) GetAllPermissions(query model.PermissionAllQuery) model.PermissionAll {
	return permission.All(query)
}
