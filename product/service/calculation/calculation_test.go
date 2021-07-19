package calculation

import (
	"github.com/gnnchya/PosCoffee/product/domain"
	"reflect"
	"testing"
)

func TestDefault(t *testing.T){
	thai := []domain.CreateMoneyStruct{{"-",1000,20,"-"},{"-",500,20,"-"},{"-",100,20,"-"},{"-",50,20,"-"},
		{"-",20,20,"-"},{"-",10,20,"-"},{"-",5,20,"-"},{"-",2,20,"-"},{"-",1,20,"-"}}
	result,change,value,_:= Calculation(19999,10000,thai)
	expectedResult := []domain.CreateMoneyStruct{{"-",1000,11,"-"},{"-",500,19,"-"},{"-",100,16,"-"},{"-",50,19,"-"},{"-",20,18,"-"},{"-",10,20,"-"},{"-",5,19,"-"},{"-",2,18,"-"},{"-",1,20,"-"}}
	expectedChange := []domain.ChangeStruct{{1000 ,9},{500, 1}, {100, 4}, {50 ,1}, {20, 2}, {0 ,0}, {5, 1}, {2, 2}, {0, 0}}
	eq := reflect.DeepEqual(change, expectedChange)
	if eq != true {
		t.Errorf("got %v want %v", change, expectedChange)
	}
	er := reflect.DeepEqual(result, expectedResult)
	if er != true {
		t.Errorf("got %v want %v", result, expectedResult)
	}
	if value != 9999{
		t.Errorf("got %d want %d", value, 9999)

	}
}

func TestNotEnoughNote(t *testing.T) {
	thai := []domain.CreateMoneyStruct{{"-",1000,2,"-"},{"-",500,20,"-"},{"-",100,20,"-"},{"-",50,20,"-"},
		{"-",20,20,"-"},{"-",10,20,"-"},{"-",5,20,"-"},{"-",2,20,"-"},{"-",1,20,"-"}}
	result,change,value,_:=Calculation(19999,10000,thai)
	expectedResult := []domain.CreateMoneyStruct{{"-",1000,0,"-"},{"-",500,5,"-"},{"-",100,16,"-"},{"-",50,19,"-"},{"-",20,18,"-"},{"-",10,20,"-"},{"-",5,19,"-"},{"-",2,18,"-"},{"-",1,20,"-"}}
	expectedChange := []domain.ChangeStruct{{1000, 2} ,{500, 15} ,{100 ,4} ,{50 ,1} ,{20 ,2} ,{0 ,0} ,{5 ,1} ,{2 ,2} ,{0 ,0}}
	eq := reflect.DeepEqual(change, expectedChange)
	if eq != true {
		t.Errorf("got %v want %v", change, expectedChange)
	}
	er := reflect.DeepEqual(result, expectedResult)
	if er != true {
		t.Errorf("got %v want %v", result, expectedResult)
	}
	if value != 9999{
		t.Errorf("got %d want %d", value, 9999)

	}
}

func TestNotEnoughChange(t *testing.T) {
	thai := []domain.CreateMoneyStruct{{"-",1000,2,"-"},{"-",500,20,"-"},{"-",100,20,"-"},{"-",50,20,"-"},
		{"-",20,20,"-"},{"-",10,20,"-"},{"-",5,20,"-"},{"-",2,20,"-"},{"-",1,20,"-"}}
	_,_,value, err:=Calculation(199999,100000,thai)
	if err == nil {
		t.Errorf("%v",err)
	}
	if value != 0{
		t.Errorf("got %d want %d", value, 9999)

	}
}