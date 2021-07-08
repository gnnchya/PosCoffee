package migrations

import (
	"context"
)

type Lang struct {
	Th string `bson:"th"`
	En string `bson:"en"`
}

var (
	ctx = context.Background()
)
