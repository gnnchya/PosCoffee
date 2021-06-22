package main

import "fmt"

func calculation(value int, note []int) map[int]int{
	change := make(map[int]int)
	var remaining = value
	for _,i := range note{
		if remaining >= i{
			change[i] = remaining/i
			remaining = value % i
		}
	}
	return change
}

func main(){
	thai := []int{1000,500,100,50,20,10,5,2,1}
	fmt.Println(calculation(23451,thai))
}
