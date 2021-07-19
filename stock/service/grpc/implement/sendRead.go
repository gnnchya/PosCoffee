package implement

import (
	"context"
	"fmt"
	"github.com/gnnchya/PosCoffee/stock/service/grpc/protobuf"
	"github.com/gnnchya/PosCoffee/stock/service/user/userin"
	)

func(impl implementation) ReadStock(ctx context.Context, input *protobuf.RequestRead) (*protobuf.ReplyRead, error){
	fmt.Println("from product ask for report", input)
	readInput := &userin.ReadInput{ID: input.Id}
	result, err := impl.userService.Read(ctx, readInput)
	if err != nil {
		return nil, err
	}
	out := &protobuf.ReplyRead{
		ID:          result.ID,
		ItemName:    result.ItemName,
		Category:    result.Category,
		Amount:      result.Amount,
		Unit:        result.Unit,
		CostPerUnit: result.CostPerUnit,
		EXPDate:     result.EXPDate,
		ImportDate:  result.ImportDate,
		Supplier:    result.Supplier,
		TotalCost:   result.TotalCost,
		TotalAmount: result.TotalAmount,
		Status:      result.Status,
	}
	return out, nil
}
