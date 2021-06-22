package main

import "fmt"

type Money struct {
	//ID 			string
	Value   	int64
	Amount 		int64
	//Currency	string
}


func Calculation(value int64, note []Money) map[int64]int64{
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
	return change
}

func main(){
	thai := []Money{{1000,20},{500,2},{100,2},{50,2},
		{20,2},{10,2},{5,2},{2,2},{1,200000}}
	fmt.Println(Calculation(23451,thai))
}
