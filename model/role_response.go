package model

type RoleShort struct {
	ID      string `json:"_id"`
	Name    string `json:"name"`
	IsAdmin bool   `json:"isAdmin"`
}