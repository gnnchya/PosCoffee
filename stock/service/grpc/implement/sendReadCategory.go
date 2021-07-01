package implement

import (
	"context"
	"fmt"
	"github.com/gnnchya/PosCoffee/stock/service/grpc/protobuf"
	"github.com/gnnchya/PosCoffee/stock/service/user/userin"
)

func(impl implementation) ReadCategoryStock(ctx context.Context, input *protobuf.RequestCategory) (*protobuf.ReplyArrRead, error){
	fmt.Println("from product ask for report", input)
	readInput := &userin.ReadCategoryAllInput{
		Category: input.Category,
		PerPage: int(input.PerPage),
		Page:     int(input.Page),
	}
	result, err := impl.userService.ReadCategoryAll(ctx, readInput)
	if err != nil {
		return nil, err
	}
	var output []*protobuf.ReplyRead
	for _, k := range result{
		out := &protobuf.ReplyRead{
			ID:          k.ID,
			ItemName:    k.ItemName,
			Category:    k.Category,
			Amount:      k.Amount,
			Unit:        k.Unit,
			CostPerUnit: k.CostPerUnit,
			EXPDate:     k.EXPDate,
			ImportDate:  k.ImportDate,
			Supplier:    k.Supplier,
			TotalCost:   k.TotalCost,
			TotalAmount: k.TotalAmount,
			Status:      k.Status,
		}
		output = append(output, out)
	}
	out := protobuf.ReplyArrRead{Reply: output}
	return &out, nil
}

