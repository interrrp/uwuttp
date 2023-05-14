// Contains the response caching system.

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/interrrp/uwuttp/uwu"
	"github.com/redis/go-redis/v9"
)

var (
	// cacheDB is the Redis client used for caching responses.
	cacheDB *redis.Client

	// ctx is the context used for operations.
	ctx context.Context
)

// init initializes the context.
func init() {
	ctx = context.Background()
}

// init initializes the cache and connects to the Redis server.
func init() {
	opts := redis.Options{
		Addr:     envOr("REDIS_ADDR", "localhost:6379"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	}

	cacheDB = redis.NewClient(&opts)

	if ping := cacheDB.Ping(ctx); ping.Err() != nil {
		panic(ping.Err())
	}
}

// cacheGet gets a key in the cache.
func cacheGet(key string) (string, error) {
	return cacheDB.Get(ctx, key).Result()
}

// cacheSet sets a key in the cache.
func cacheSet(key, value string) error {
	return cacheDB.Set(ctx, key, value, 0).Err()
}

// encodeUwU encodes UwU text into a key.
func encodeUwUKey(text string, cfg uwu.Config) string {
	return text + fmt.Sprintf("%v", cfg)
}
