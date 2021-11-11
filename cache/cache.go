package cache

import (
	"log"
	"time"

	"github.com/allegro/bigcache/v3"
)

var cache *bigcache.BigCache

// Init ...
func Init() {
	// The time after which entries can be evicted is 30 days
	const cacheTime = 24 * 30 * time.Hour // 30 days
	c, err := bigcache.NewBigCache(bigcache.DefaultConfig(cacheTime))
	if err != nil {
		log.Fatalf("Cannot init Cache %v", err)
	}
	cache = c

	// Cache roles
	Roles()
}

// GetInstance ...
func GetInstance() *bigcache.BigCache {
	return cache
}
