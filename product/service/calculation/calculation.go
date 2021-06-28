package calculation

import (
	"fmt"
	"github.com/gnnchya/PosCoffee/product/domain"
)


func Calculation(paid int64, price int64, note []domain.CreateMoneyStruct) ([]domain.CreateMoneyStruct, []domain.ChangeStruct, error){
	value := paid - price
	var change []domain.ChangeMoney
	for x,i := range note{
		if value >= i.Value{
			if value/i.Value > i.Amount{
				change[x].Value = i.Value
				change[x].Amount = i.Amount
				value = value - (i.Amount*i.Value)
				note[x].Amount = 0
			} else {
				change[x].Value = i.Value
				change[x].Amount = value/i.Value
				note[x].Amount = i.Amount - (value/i.Value)
				value = value % i.Value
			}
		}
	}
	if value != 0{
		return nil,nil, fmt.Errorf("error: not enough change")
	}
	return note, change,nil
}