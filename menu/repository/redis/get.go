package redis

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/go-redis/redis/v8"
)

func (repoRedis *Redis) Get(ctx context.Context, key string, dest interface{}) (err error) {
	val, err := repoRedis.Client.Get(ctx, key).Bytes()
	if err == redis.Nil {
		return errors.New("error : Key does not exist")
	} else if err != nil {
		return  err
	}
	return json.Unmarshal(val, &dest)
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