package usermngmt

import (
	"github.com/Selly-Modules/usermngmt/role"
	"github.com/Selly-Modules/usermngmt/user"
)

// userHandle ...
func (s Service) userHandle() user.Handle {
	return user.Handle{
		Col:     s.getUserCollection(),
		RoleCol: s.getRoleCollection(),
	}
}

// roleHandle ...
func (s Service) roleHandle() role.Handle {
	return role.Handle{
		Col: s.getRoleCollection(),
	}
}
