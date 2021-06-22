package main

import (
	"fmt"
)

func main(){
	val := 2345
	//var temp int
	m := make(map[int]int)
	t := []int{1000,500,100,50,20,10,5,2,1}
var rem = val
for _,i := range t{
		if rem >= i{
			m[i] = rem/i
			rem = val % i
		}
	}
	fmt.Println(m)
}
