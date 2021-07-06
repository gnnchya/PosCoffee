package consumer

import (
	"context"
	"github.com/gnnchya/PosCoffee/oAuth/service/consumer/consumerin"
)

//go:generate mockery --name=Service
type Service interface {
	Create(ctx context.Context, input *consumerin.CreateInput) (ID string, err error)
}
