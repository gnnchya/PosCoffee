package redis

import (
	"context"
)

func (repoRedis *Redis) Del(ctx context.Context, key string) (err error) {
	_ , err = repoRedis.Client.Del(ctx, key).Result()
	return err
}
