package calculation

import (
	"fmt"
	"github.com/gnnchya/PosCoffee/product/domain"
)


func Calculation(paid int64, price int64, note []domain.CreateMoneyStruct) ([]domain.CreateMoneyStruct, []domain.ChangeStruct,int64, error){
	value := paid - price
	remain := paid - price
	change := make([]domain.ChangeStruct, len(note))
	for x,i := range note{
		if remain >= i.Value{
			if remain/i.Value > i.Amount{
				change[x].Value = i.Value
				change[x].Amount = i.Amount
				remain = remain - (i.Amount*i.Value)
				note[x].Amount = 0
			} else {
				change[x].Value = i.Value
				change[x].Amount = remain/i.Value
				note[x].Amount = i.Amount - (remain/i.Value)
				remain = remain % i.Value
			}
		}
	}
	if remain != 0{
		return nil,nil,0, fmt.Errorf("error: not enough change")
	}
	return note, change,value,nil
}