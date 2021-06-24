package main

import (
	"reflect"
	"testing"
)

func TestDefault(t *testing.T){
	thai := []CreateMoneyStruct{{"-",1000,20,"-"},{"-",500,20,"-"},{"-",100,20,"-"},{"-",50,20,"-"},
		{"-",20,20,"-"},{"-",10,20,"-"},{"-",5,20,"-"},{"-",2,20,"-"},{"-",1,20,"-"}}
	_,change,_:=Calculation(9999,0, thai)
	expectedResult := map[int64]int64{2:2,5:1,20:2,50:1,100:4,500:1,1000:9}
	eq := reflect.DeepEqual(change, expectedResult)
	if eq != true {
		t.Errorf("got %v want %v", change, expectedResult)
	}
}

func TestNotEnoughNote(t *testing.T) {
	thai := []CreateMoneyStruct{{"-",1000,2,"-"},{"-",500,20,"-"},{"-",100,20,"-"},{"-",50,20,"-"},
		{"-",20,20,"-"},{"-",10,20,"-"},{"-",5,20,"-"},{"-",2,20,"-"},{"-",1,20,"-"}}
	_,change,_:=Calculation(9999,0, thai)
	expectedResult := map[int64]int64{2:2,5:1, 20:2, 50:1, 100:4, 500:15, 1000:2}
	eq := reflect.DeepEqual(change, expectedResult)
	if eq != true {
		t.Errorf("got %v want %v", change, expectedResult)
	}
}

func TestNotEnoughChange(t *testing.T) {
	thai := []CreateMoneyStruct{{"-",1000,2,"-"},{"-",500,20,"-"},{"-",100,20,"-"},{"-",50,20,"-"},
		{"-",20,20,"-"},{"-",10,20,"-"},{"-",5,20,"-"},{"-",2,20,"-"},{"-",1,20,"-"}}
	_,_, err:=Calculation(99999,0, thai)
	if err == nil {
		t.Errorf("%v",err)
	}
}