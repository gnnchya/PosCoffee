package redis

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

func (repoRedis *Redis) Get(ctx context.Context, key string) (result string, err error) {
	val, err := repoRedis.Client.Get(ctx, key).Result()
	if err == redis.Nil {
		// Key does not exists
		return "", nil
	} else if err != nil {
		return "", err
	}

	return val, nil
}

func (repoRedis *Redis) GetExpire(ctx context.Context, key string) (result time.Duration, err error) {
	val, err := repoRedis.Client.TTL(ctx, key).Result()
	if err == redis.Nil {
		// Key does not exists
		return 0, nil
	} else if err != nil {
		return 0, err
	}

	return val, nil
}