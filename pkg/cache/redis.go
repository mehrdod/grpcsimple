package cache

import (
	"context"
	"strconv"

	"github.com/go-redis/redis/v8"
)

type RedisCache struct {
	client     *redis.Client
	cacheLimit int
}

func NewRedisCache(host, port string) *RedisCache {
	return &RedisCache{
		client: redis.NewClient(&redis.Options{
			Addr:     host + ":" + port,
			Password: "", // no password set
			DB:       0,  // use default DB
		}),
	}
}

func (r *RedisCache) Get(key string) (int64, error) {

	value, err := r.client.Get(context.Background(), key).Result()
	if err != nil {
		return 0, err
	}
	return strconv.ParseInt(value, 10, 64)
}

func (r *RedisCache) Set(key string, values int64) error {
	return r.client.Set(context.Background(), key, values, 0).Err()
}
