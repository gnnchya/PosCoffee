package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

type Redis struct {
	Client *redis.Client
	Expires time.Duration
}

func New(ctx context.Context, server string, password string, exp time.Duration) (repoRedis *Redis, err error) {
	client := redis.NewClient(&redis.Options{
		Addr:     server,
		Password: password,
		DB:       0,
	})

	retry := 3 // Retry connecting 3 times, if cannot ping
	for true {
		_, err = client.Ping(ctx).Result()
		if err != nil {
			retry--
			if retry < 0 {
				return nil, err
			}

			// Wait before reconnecting
			continue
		}
		// If we can PING without error, just return
		break
	}

	repoRedis = &Redis{
		Client: client,
		Expires: exp,
	}

	return repoRedis, nil
}