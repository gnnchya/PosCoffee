package kafka

import (
	"context"
	"os"
	"time"
)

func (message Kafka) Consumer() {
	go message.Broker.Start(func(ctx context.Context, err error) {
		time.Sleep(2 * time.Second)
		os.Exit(0)
	})
	return
}