package redis

import (
	"context"
	"encoding/json"
)

func (repoRedis *Redis) Set(ctx context.Context, key string, value interface{})(err error) {
	str, err := json.Marshal(&value)
	if err != nil {
		return err
	}

	err = repoRedis.Client.Set(ctx, key, str, repoRedis.Expires).Err()
	if err != nil {
		return err
	}

	return nil
}
