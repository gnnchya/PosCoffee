package redis

import (
	"context"
	"encoding/json"
	"time"
)

func (repoRedis *Redis) Set(ctx context.Context, key string, value interface{}, expire time.Duration) (err error) {
	str, err := json.Marshal(value)
	if err != nil {
		return err
	}

	err = repoRedis.Client.Set(ctx, key, str, expire).Err()
	if err != nil {
		return err
	}

	return nil
}
