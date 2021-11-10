package model

import (
	"time"
)

// RoleShort ...
type RoleShort struct {
	ID      string `json:"_id"`
	Name    string `json:"name"`
	IsAdmin bool   `json:"isAdmin"`
}

// Role ...
type Role struct {
	ID        string    `json:"_id"`
	Name      string    `json:"name"`
	Code      string    `json:"code"`
	IsAdmin   bool      `json:"isAdmin"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type (
	// RoleAll ...
	RoleAll struct {
		List  []Role `json:"list"`
		Total int64  `json:"total"`
	}
)
