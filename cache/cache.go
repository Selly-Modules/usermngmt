package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/logrusorgru/aurora"
)

var (
	c *redis.Client
)

// Init ...
func Init(uri, pwd string) {
	c = redis.NewClient(&redis.Options{
		Addr:     uri,
		Password: pwd,
		DB:       0, // use default DB
	})

	// Test
	_, err := c.Ping(context.Background()).Result()

	if err != nil {
		log.Fatal("Cannot connect to redis", uri, err)
	}

	fmt.Println(aurora.Green("*** CONNECTED TO REDIS: " + uri))

	// Cache roles
	Roles()
}

// GetInstance ...
func GetInstance() *redis.Client {
	return c
}

// SetKeyValue ...
func SetKeyValue(key string, value interface{}, expiration time.Duration) error {
	ctx := context.Background()
	dataByte, _ := json.Marshal(value)
	r := c.Set(ctx, key, dataByte, expiration)
	return r.Err()
}

// GetValueByKey ...
func GetValueByKey(key string) ([]byte, error) {
	ctx := context.Background()
	return c.Get(ctx, key).Bytes()
}
