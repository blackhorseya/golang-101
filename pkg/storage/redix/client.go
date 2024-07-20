package redix

import (
	"github.com/redis/go-redis/v9"
)

// Options is a struct to hold the configuration options for the Redis client.
type Options struct {
	Addr string
}

// NewClient creates a new Redis client.
func NewClient(options Options) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr: options.Addr,
	})

	return client, nil
}
