package cache

// CachedRole ...
type CachedRole struct {
	Role        string   `json:"role"`
	IsAdmin     bool     `json:"isAdmin"`
	Permissions []string `json:"permissions"`
}
