package consumer

import (
	"context"
	"github.com/gnnchya/PosCoffee/email/service/consumer/consumerin"
)

//go:generate mockery --name=Service
type Service interface {
	Create(ctx context.Context, input *consumerin.CreateConsumerInput) (ID string, err error)
}
