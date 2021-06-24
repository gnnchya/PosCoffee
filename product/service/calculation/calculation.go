package calculation

import (
	"fmt"
)

type CreateMoneyStruct struct {
	ID 			string
	Value   	int64
	Amount 		int64
	Currency	string
}

func Calculation(paid int64, price int64, note []CreateMoneyStruct) ([]CreateMoneyStruct, map[int64]int64, error){
	value := paid - price
	change := make(map[int64]int64)
	for x,i := range note{
		if value >= i.Value{
			if value/i.Value > i.Amount{
				change[i.Value] = i.Amount
				value = value - (i.Amount*i.Value)
				note[x].Amount = 0
			} else {
				change[i.Value] = value/i.Value
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