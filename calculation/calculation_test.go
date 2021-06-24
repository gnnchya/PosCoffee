package main

import (
	"reflect"
	"testing"
)

func TestDefault(t *testing.T){
	thai := []CreateMoneyStruct{{"-",1000,20,"-"},{"-",500,20,"-"},{"-",100,20,"-"},{"-",50,20,"-"},
		{"-",20,20,"-"},{"-",10,20,"-"},{"-",5,20,"-"},{"-",2,20,"-"},{"-",1,20,"-"}}
	result,change,_:=Calculation(9999,0,thai)
	expectedResult := []CreateMoneyStruct{{"-",1000,11,"-"},{"-",500,19,"-"},{"-",100,16,"-"},{"-",50,19,"-"},{"-",20,18,"-"},{"-",10,20,"-"},{"-",5,19,"-"},{"-",2,18,"-"},{"-",1,20,"-"}}
	expectedChange := map[int64]int64{2:2,5:1,20:2,50:1,100:4,500:1,1000:9}
	eq := reflect.DeepEqual(change, expectedChange)
	if eq != true {
		t.Errorf("got %v want %v", change, expectedChange)
	}
	er := reflect.DeepEqual(result, expectedResult)
	if er != true {
		t.Errorf("got %v want %v", result, expectedResult)
	}
}

func TestNotEnoughNote(t *testing.T) {
	thai := []CreateMoneyStruct{{"-",1000,2,"-"},{"-",500,20,"-"},{"-",100,20,"-"},{"-",50,20,"-"},
		{"-",20,20,"-"},{"-",10,20,"-"},{"-",5,20,"-"},{"-",2,20,"-"},{"-",1,20,"-"}}
	result,change,_:=Calculation(9999,0,thai)
	expectedResult := []CreateMoneyStruct{{"-",1000,0,"-"},{"-",500,5,"-"},{"-",100,16,"-"},{"-",50,19,"-"},{"-",20,18,"-"},{"-",10,20,"-"},{"-",5,19,"-"},{"-",2,18,"-"},{"-",1,20,"-"}}
	expectedChange := map[int64]int64{2:2,5:1, 20:2, 50:1, 100:4, 500:15, 1000:2}
	eq := reflect.DeepEqual(change, expectedResult)
	if eq != true {
		t.Errorf("got %v want %v", change, expectedChange)
	}
	er := reflect.DeepEqual(result, expectedResult)
	if er != true {
		t.Errorf("got %v want %v", result, expectedResult)
	}
}

func TestNotEnoughChange(t *testing.T) {
	thai := []CreateMoneyStruct{{"-",1000,2,"-"},{"-",500,20,"-"},{"-",100,20,"-"},{"-",50,20,"-"},
		{"-",20,20,"-"},{"-",10,20,"-"},{"-",5,20,"-"},{"-",2,20,"-"},{"-",1,20,"-"}}
	_,_, err:=Calculation(99999,0,thai)
	if err == nil {
		t.Errorf("%v",err)
	}
}