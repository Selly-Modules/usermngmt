package cache

import (
	"log"
	"time"

	"github.com/allegro/bigcache/v3"
)

var cache *bigcache.BigCache

// Init ...
func Init() {
	// The time after which entries can be evicted is 5 years
	c, err := bigcache.NewBigCache(bigcache.DefaultConfig(43800 * time.Hour))
	if err != nil {
		log.Fatalf("Cannot init Cache %v", err)
	}
	cache = c
}

// GetInstance ...
func GetInstance() *bigcache.BigCache {
	return cache
}
