package mongodb

import (
	"context"
	"fmt"
	"github.com/gnnchya/PosCoffee/stock/domain"
)

func (repo *Repository)ReadAllSorted(ctx context.Context, field string, order string)(result []domain.CreateStruct, err error){
	if repo.CheckSortOrder(order) == false{
		return result, fmt.Errorf("error: wrong order")
	}
	switch field{
rew
	}
	return result, err
}


