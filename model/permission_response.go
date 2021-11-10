package model

import (
	"time"
)

// Permission ...
type Permission struct {
	ID        string    `json:"_id"`
	Name      string    `json:"name"`
	Code      string    `json:"code"`
	RoleID    string    `json:"roleId"`
	Desc      string    `json:"desc"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type (
	// PermissionAll ...
	PermissionAll struct {
		List  []Permission `json:"list"`
		Total int64        `json:"total"`
	}
)
