package store

import "github.com/go-redis/redis"

// RedisStorer implements a persistent key-value store using Redis
type RedisStorer struct {
	client *redis.Client
}

func (r *RedisStorer) Get(key string) (string, error) {
	value, err := r.client.Get(key).Result()
	return value, err
}

func (r *RedisStorer) Set(key string, value string) error {
	return r.client.Set(key, value, 0).Err()
}

func (r *RedisStorer) Increment(key string) error {
	return r.client.Incr(key).Err()
}

// NewRedisStorer creates a new RedisStorer
func NewRedisStorer(client *redis.Client) *RedisStorer {
	return &RedisStorer{client}
}
