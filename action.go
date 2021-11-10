package usermngmt

import "github.com/Selly-Modules/usermngmt/internal"

// Create ...
func (s Service) Create(payload internal.CreateOptions) error {
	return s.userHandle().Create(payload)
}

// Update ...
func (s Service) Update(userID string, payload internal.UpdateOptions) error {
	return s.userHandle().UpdateByUserID(userID, payload)
}

// ChangeUserPassword ...
func (s Service) ChangeUserPassword(userID string, payload internal.ChangePasswordOptions) error {
	return s.userHandle().ChangeUserPassword(userID, payload)
}

func (s Service) ChangeUserStatus(userID, newStatus string) error {
	return s.userHandle().ChangeUserStatus(userID, newStatus)
}

func (s Service) All(query internal.AllQuery) internal.UserAll {
	return s.userHandle().All(query)
}

func (s Service) RoleCreate(payload internal.RoleCreateOptions) error {
	return s.roleHandle().Create(payload)
}
