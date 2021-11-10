package model

import (
	"time"
)

// User ...
type User struct {
	ID        string    `json:"_id"`
	Name      string    `json:"name"`
	Phone     string    `json:"phone"`
	Email     string    `json:"email"`
	Status    string    `json:"status"`
	Role      RoleShort `json:"role"`
	Other     string    `json:"other"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type (
	// UserAll ...
	UserAll struct {
		List  []User `json:"list"`
		Total int64  `json:"total"`
	}
)
