package main

import (
	"fmt"
	"log"
)

type CreateMoneyStruct struct {
	ID 			string
	Value   	int64
	Amount 		int64
	Currency	string
}

func Calculation(value int64, note []CreateMoneyStruct) (map[int64]int64, error){
	change := make(map[int64]int64)
	for _,i := range note{
		if value >= i.Value{
			if value/i.Value > i.Amount{
				change[i.Value] = i.Amount
				value = value - (i.Amount*i.Value)
			} else {
				change[i.Value] = value/i.Value
				value = value % i.Value
			}
		}
	}
	if value != 0{
		return nil, fmt.Errorf("error: not enough change")
	}
	return change,nil
}

func main(){
	thai := []CreateMoneyStruct{{"-",1000,200,"-"},{"-",500,2,"-"},{"-",100,2,"-"},{"-",50,2,"-"},
		{"-",20,2,"-"},{"-",10,2,"-"},{"-",5,2,"-"},{"-",2,2,"-"},{"-",1,200000,"-"}}
	change, err := Calculation(23456,thai)
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println(change)
}
