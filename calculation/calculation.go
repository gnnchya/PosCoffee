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
		return note,nil, fmt.Errorf("error: not enough change")
	}
	return note, change,nil
}

func main(){
	thai := []CreateMoneyStruct{{"-",1000,20,"-"},{"-",500,20,"-"},{"-",100,20,"-"},{"-",50,20,"-"},
		{"-",20,20,"-"},{"-",10,20,"-"},{"-",5,20,"-"},{"-",2,20,"-"},{"-",1,20,"-"}}
	remain,change, err := Calculation(9999,0,thai)
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println(change)
	for _,x := range remain{
		fmt.Println("asd",x)
	}
}

//Default
//Not enough money
//Not enough change