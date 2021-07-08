package role

import (
	"context"

	"github.com/gnnchya/PosCoffee/authorize/repository/mongodb"
)

type Repository struct {
	*mongodb.Repository
}

func New(ctx context.Context, config *mongodb.Config) (repo *Repository, err error) {
	mongoDB, err := mongodb.New(ctx, config)
	if err != nil {
		return nil, err
	}

	return &Repository{mongoDB}, nil
}
